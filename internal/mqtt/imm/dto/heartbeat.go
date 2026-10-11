package dto

import (
	"encoding/json"
	"fmt"
	"time"
)

// ============================================================
// LOCAL DATE TIME
// ============================================================
//
// ESP32 sends:
//
//     2026-08-31 16:41:43
//
// Go time.Time expects RFC3339 by default:
//
//     2026-08-31T16:41:43+05:30
//
// This custom type allows the DTO to directly accept the
// ESP32 timestamp format.
// ============================================================

type LocalDateTime struct {
	time.Time
}

// ============================================================
// JSON UNMARSHAL
// ============================================================

func (t *LocalDateTime) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return fmt.Errorf("timestamp must be a string: %w", err)
	}

	layouts := []string{
		time.RFC3339Nano,
		time.RFC3339,
		"2006-01-02 15:04:05",
	}

	for _, layout := range layouts {
		parsed, err := time.Parse(layout, value)
		if err == nil {
			t.Time = parsed
			return nil
		}
	}

	return fmt.Errorf("invalid timestamp format: %q", value)
}

// func (t *LocalDateTime) UnmarshalJSON(data []byte) error {

// 	var value string

// 	if err := json.Unmarshal(data, &value); err != nil {
// 		return err
// 	}

// 	parsed, err := time.Parse(
// 		"2006-01-02 15:04:05",
// 		value,
// 	)

// 	if err != nil {
// 		return err
// 	}

// 	t.Time = parsed

// 	return nil
// }

// ============================================================
// JSON MARSHAL
// ============================================================

// func (t LocalDateTime) MarshalJSON() ([]byte, error) {

// 	return json.Marshal(
// 		t.Format("2006-01-02 15:04:05"),
// 	)
// }

// ============================================================
// IMM HEARTBEAT DTO
// ============================================================

type IMMHeartbeatDTO struct {
	EventID         string        `json:"event_id"`
	TenantID        int           `json:"tenant_id"`
	DeviceID        string        `json:"device_id"`
	MachineID       string        `json:"machine_id"`
	DeviceType      string        `json:"device_type"`
	FirmwareVersion string        `json:"firmware_version"`
	Status          string        `json:"status"`
	MACID           string        `json:"mac_id"`
	IPAddress       string        `json:"ip_address"`
	RSSI            int           `json:"rssi"`
	MQTTClientID    string        `json:"mqtt_client_id"`
	Timestamp       LocalDateTime `json:"timestamp"`
}

// type IMMHeartbeatDTO struct {
// 	EventID string `json:"event_id"`

// 	// TenantID int `json:"tenant_id"`
// 	TenantID string `json:"tenant_id"`

// 	DeviceID string `json:"device_id"`

// 	MachineID string `json:"machine_id"`

// 	MoldNo string `json:"mold_no"`

// 	Station string `json:"station"`

// 	DeviceType string `json:"device_type"`

// 	FirmwareVersion string `json:"firmware_version"`

// 	Status string `json:"status"`

// 	MACID string `json:"mac_id"`

// 	IPAddress string `json:"ip_address"`

// 	RSSI int `json:"rssi"`

// 	ProductionCount int64 `json:"production_count"`

// 	LastCycleTimeSec float64 `json:"last_cycle_time_sec"`

// 	MQTTClientID string `json:"mqtt_client_id"`

// 	Timestamp LocalDateTime `json:"timestamp"`
// }
