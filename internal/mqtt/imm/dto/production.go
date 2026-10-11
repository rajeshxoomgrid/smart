package dto

import "time"

// ProductionDTO contains the production payload received from ESP32.
type ProductionDTO struct {
	EventID      string    `json:"event_id"`
	TenantID     int       `json:"tenant_id"`
	DeviceID     string    `json:"device_id"`
	MachineID    string    `json:"machine_id"`
	Count        int64     `json:"production_count"`
	CycleTimeSec float64   `json:"cycle_time_sec"`
	Timestamp    time.Time `json:"production_time"`
}

// ProductionResponseDTO contains production data enriched with database values.
type ProductionResponseDTO struct {
	EventID      string    `json:"event_id"`
	TenantID     int       `json:"tenant_id"`
	DeviceID     string    `json:"device_id"`
	MachineID    string    `json:"machine_id"`
	MoldNo       string    `json:"mold_no"`
	Count        int64     `json:"production_count"`
	CycleTimeSec float64   `json:"cycle_time_sec"`
	Timestamp    time.Time `json:"production_time"`
	Reset        bool      `json:"reset"`
}

// type ProductionDTO struct {
// 	EventID      string    `json:"event_id"`
// 	TenantID     string    `json:"tenant_id"`
// 	CustomerID   string    `json:"customer_id"`
// 	DeviceID     string    `json:"device_id"`
// 	MachineID    string    `json:"machine_id"`
// 	Station      string    `json:"station"`
// 	Count        int64     `json:"production_count"`
// 	CycleTimeSec float64   `json:"cycle_time_sec"`
// 	Timestamp    time.Time `json:"production_time"`
// }
