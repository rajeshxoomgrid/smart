package immhttpservice

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"strings"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"

	immconstant "github.com/rajeshbond/smart/internal/http/imm/imm_constant"
	immdto "github.com/rajeshbond/smart/internal/http/imm/imm_dto"
	immerror "github.com/rajeshbond/smart/internal/http/imm/imm_error"
)

// ============================================================
// IMM MQTT CONSTANTS
// ============================================================

// const (
// 	// MQTT command name.
// 	CommandSetMoldNo = "SET_MOLD_NO"

// 	// MQTT command topic prefix.
// 	// Final topic:
// 	// factory/imm/command/{device_id}
// 	TopicIMMCommand = "factory/imm/command/"

// 	// Command states.
// 	StatusPending = "PENDING"
// 	StatusSuccess = "SUCCESS"
// 	StatusFailed  = "FAILED"
// )

// ============================================================
// SET MOLD NO
// ============================================================
//
// Flow:
//
// 1. Validate device_id and mold_no.
// 2. Start short DB transaction.
// 3. Lock/read device for tenant.
// 4. Validate device type = IMM.
// 5. Validate requested mold is different.
// 6. Check whether another PENDING command exists.
// 7. Create PENDING command.
// 8. COMMIT transaction.
//
// 9. Publish MQTT command AFTER DB COMMIT.
//
// 10. If MQTT is disconnected:
//       PENDING -> FAILED
//
// 11. If MQTT publish timeout:
//       PENDING -> FAILED
//
// 12. If MQTT publish error:
//       PENDING -> FAILED
//
// IMPORTANT:
//
// device_master.mold_no is NOT changed here.
//
// It is changed only after ESP32 sends:
//
// SUCCESS ACK
//
// ============================================================

func (s *ImmHttpService) SetMoldNo(
	ctx context.Context,
	tenantID int64,
	userID int64,
	deviceID string,
	moldNo string,
) error {

	// --------------------------------------------------------
	// STEP 0
	// Normalize input.
	// --------------------------------------------------------

	deviceID = strings.TrimSpace(deviceID)
	moldNo = strings.TrimSpace(moldNo)

	if deviceID == "" {
		return immerror.ErrDeviceNotFound
	}

	if moldNo == "" {
		return errors.New("mold number is required")
	}

	var commandID int64

	// --------------------------------------------------------
	// STEP 1
	// CREATE PENDING COMMAND
	//
	// Do NOT keep the DB transaction open while publishing
	// MQTT.
	// --------------------------------------------------------

	err := s.ImmHttpService.WithTransaction(
		ctx,
		func(tx *sql.Tx) error {

			// ------------------------------------------------
			// Get device for tenant.
			// ------------------------------------------------

			device, err := s.ImmHttpService.GetDeviceForTenantTx(
				ctx,
				tx,
				tenantID,
				deviceID,
			)
			if err != nil {
				return err
			}

			// ------------------------------------------------
			// Device must be IMM.
			// ------------------------------------------------

			if !strings.EqualFold(
				strings.TrimSpace(device.DeviceType),
				"imm",
			) {
				return immerror.ErrInvalidDevice
			}

			// ------------------------------------------------
			// Requested mold must be different from current
			// mold.
			// ------------------------------------------------

			if strings.EqualFold(
				strings.TrimSpace(device.MoldNo),
				moldNo,
			) {
				return immerror.ErrSameMoldNo
			}

			// ------------------------------------------------
			// Check whether this device already has a
			// PENDING mold-change command.
			// ------------------------------------------------

			pending, err := s.ImmHttpService.GetPendingForDeviceTx(
				ctx,
				tx,
				tenantID,
				deviceID,
			)
			if err != nil {
				return err
			}

			if pending != nil {
				return immerror.ErrCommandPending
			}

			// ------------------------------------------------
			// Create PENDING command.
			// ------------------------------------------------

			commandID, err = s.ImmHttpService.CreatePendingTx(
				ctx,
				tx,
				tenantID,
				userID,
				device.DeviceID,
				device.MachineID,
				device.MoldNo,
				moldNo,
			)
			if err != nil {
				return err
			}

			return nil
		},
	)

	if err != nil {
		return err
	}

	// --------------------------------------------------------
	// STEP 2
	// Prepare MQTT payload.
	// --------------------------------------------------------

	payload := immdto.SetMoldNoMQTTRequest{
		DeviceID: deviceID,
		Command:  immconstant.CommandSetMoldNo,
		MoldNo:   moldNo,
	}

	data, err := json.Marshal(payload)
	if err != nil {

		reason := "failed to create MQTT payload: " + err.Error()

		// PENDING -> FAILED.
		_ = s.markCommandFailed(
			context.Background(),
			commandID,
			reason,
		)

		return err
	}

	// --------------------------------------------------------
	// STEP 3
	// MQTT topic.
	//
	// Example:
	//
	// factory/imm/command/05@xoom
	// --------------------------------------------------------

	topic := immconstant.TopicIMMCommand + deviceID

	// --------------------------------------------------------
	// STEP 4
	// Check MQTT connection.
	// --------------------------------------------------------

	if s.MQTT == nil {

		reason := "MQTT client is nil"

		_ = s.markCommandFailed(
			context.Background(),
			commandID,
			reason,
		)

		return errors.New(reason)
	}

	if !s.MQTT.IsConnected() {

		reason := "MQTT client is not connected"

		_ = s.markCommandFailed(
			context.Background(),
			commandID,
			reason,
		)

		return errors.New(reason)
	}

	// --------------------------------------------------------
	// STEP 5
	// Publish MQTT command.
	//
	// QoS = 1
	// Retained = false
	// --------------------------------------------------------

	token := s.MQTT.Publish(
		topic,
		1,
		false,
		data,
	)

	// --------------------------------------------------------
	// STEP 6
	// Wait maximum 5 seconds for MQTT library publish
	// operation.
	//
	// IMPORTANT:
	//
	// This is NOT the ESP32 ACK timeout.
	//
	// ESP32 ACK timeout is handled separately by the
	// background watchdog.
	//
	// MQTT publish timeout only means the broker publish
	// itself did not complete.
	// --------------------------------------------------------

	if !token.WaitTimeout(5 * time.Second) {

		reason := "MQTT publish timeout"

		_ = s.markCommandFailed(
			context.Background(),
			commandID,
			reason,
		)

		return errors.New(reason)
	}

	// --------------------------------------------------------
	// STEP 7
	// Check MQTT publish error.
	// --------------------------------------------------------

	if err := token.Error(); err != nil {

		reason := "MQTT publish failed: " + err.Error()

		_ = s.markCommandFailed(
			context.Background(),
			commandID,
			reason,
		)

		return err
	}

	// --------------------------------------------------------
	// STEP 8
	// MQTT command successfully handed to MQTT client/broker.
	//
	// DO NOT update device_master.mold_no here.
	//
	// device_master.mold_no will be updated only after:
	//
	// ESP32 -> SUCCESS ACK
	//
	// --------------------------------------------------------

	log.Printf(
		"IMM mold change command published: command_id=%d device=%s mold=%s topic=%s",
		commandID,
		deviceID,
		moldNo,
		topic,
	)

	return nil
}

// ============================================================
// MARK COMMAND FAILED
// ============================================================
//
// This helper changes:
//
// PENDING -> FAILED
//
// It intentionally uses a NEW short transaction.
//
// It does NOT change SUCCESS or an already FAILED command.
//
// This protects against a race between:
//   - MQTT ACK handler
//   - MQTT publish failure
//   - 30-second timeout worker
//
// ============================================================

func (s *ImmHttpService) markCommandFailed(
	ctx context.Context,
	commandID int64,
	reason string,
) error {

	reason = strings.TrimSpace(reason)

	if reason == "" {
		reason = "unknown failure"
	}

	return s.ImmHttpService.WithTransaction(
		ctx,
		func(tx *sql.Tx) error {

			return s.ImmHttpService.MarkFailedTx(
				ctx,
				tx,
				commandID,
				reason,
			)
		},
	)
}

// ============================================================
// KEEP PAHO IMPORT USED
// ============================================================
//
// This compile-time reference ensures the MQTT dependency
// remains explicit in this service file if your ImmHttpService
// MQTT field is:
//
//     MQTT paho.Client
//
// If your struct already imports paho.Client in its own file,
// this declaration can be removed.
//
// ============================================================

var _ paho.Client
