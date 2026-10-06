package handler

import (
	"encoding/json"
	"log"

	paho "github.com/eclipse/paho.mqtt.golang"

	dto "github.com/rajeshbond/smart/internal/mqtt/imm/dto"
)

// ============================================================
// HEARTBEAT HANDLER
// ============================================================

type HeartbeatHandler struct{}

// ============================================================
// CONSTRUCTOR
// ============================================================

func NewHeartbeatHandler() *HeartbeatHandler {
	return &HeartbeatHandler{}
}

// ============================================================
// MQTT HEARTBEAT HANDLER
// ============================================================

func (h *HeartbeatHandler) HeartbeatHandler() paho.MessageHandler {

	return func(
		client paho.Client,
		msg paho.Message,
	) {

		// --------------------------------------------------------
		// NIL MESSAGE CHECK
		// --------------------------------------------------------

		if msg == nil {
			log.Println("❌ IMM Heartbeat message is nil")
			return
		}

		// --------------------------------------------------------
		// DTO
		// --------------------------------------------------------

		var req dto.IMMHeartbeatDTO

		// --------------------------------------------------------
		// JSON UNMARSHAL
		// --------------------------------------------------------

		if err := json.Unmarshal(
			msg.Payload(),
			&req,
		); err != nil {

			log.Printf(
				"❌ IMM Heartbeat JSON Error | Error=%v | Payload=%s",
				err,
				string(msg.Payload()),
			)

			return
		}

		// --------------------------------------------------------
		// LOG HEARTBEAT
		// --------------------------------------------------------

		log.Printf(
			"💓 IMM Heartbeat |Event = %s | Device=%s | Machine=%s | Tenant=%d | Status=%s | RSSI=%d | Mold_no=%s",
			req.EventID,
			req.DeviceID,
			req.MachineID,
			req.TenantID,
			req.Status,
			req.RSSI,
			req.MoldNo,
		)

		// --------------------------------------------------------
		// DETAILED LOG
		// --------------------------------------------------------

		log.Printf(
			"   Mold=%s | Production=%d | CycleTime=%.3f sec | IP=%s | MAC=%s | Timestamp=%s",
			req.MoldNo,
			req.ProductionCount,
			req.LastCycleTimeSec,
			req.IPAddress,
			req.MACID,
			req.Timestamp.Format("2006-01-02 15:04:05"),
		)
	}
}
