package dto

import "time"

type GetProductionRequest struct {
	TenantID string `json:"tenant_id"`
	DeviceID string `json:"device_id"`
	Station  string `json:"station,omitempty"`
	Variant  string `json:"variant,omitempty"`
	Limit    int    `json:"limit"`
}

type ProductionResponse1 struct {
	EventID string `json:"event_id"`

	TenantCode string `json:"tenant_id"`

	CustomerID string `json:"customer_id"`

	DeviceID string `json:"device_id"`

	MachineID string `json:"machine_id"`

	Station string `json:"station"`

	ProductionCount int64 `json:"production_count"`

	CycleTimeSec float64 `json:"cycle_time_sec"`

	ProductionTime time.Time `json:"production_time"`

	CreatedAt time.Time `json:"created_at"`

	//--------------------------------

	ShiftName string `json:"shift_name,omitempty"`

	ShiftStart time.Time `json:"shift_start,omitempty"`

	ShiftEnd time.Time `json:"shift_end,omitempty"`

	ShiftProduction int64 `json:"shift_production,omitempty"`
}

type ProductionResponse struct {
	EventID    string `json:"event_id"`
	TenantID   string `json:"tenant_id"`
	ShiftName  string `json:"shift_name"`
	CustomerID string `json:"customer_id"`
	DeviceID   string `json:"device_id"`
	MachineID  string `json:"machine_id"`
	Station    string `json:"station"`
	Variant    string `json:"variant"`

	ProductionCount int64 `json:"production_count"`

	DayProduction int64 `json:"day_production"`

	OEE    OEE `json:"oee"`
	DayOEE OEE `json:"day_oee"`

	CycleTimeSec   float64   `json:"cycle_time_sec"`
	ProductionTime time.Time `json:"production_time"`
	CreatedAt      time.Time `json:"created_at"`
}

// type ProductionResponse struct {
// 	EventID         string `json:"event_id"`
// 	TenantID        string `json:"tenant_id"`
// 	ShiftName       string `json:"shift_name"`
// 	CustomerID      string `json:"customer_id"`
// 	DeviceID        string `json:"device_id"`
// 	MachineID       string `json:"machine_id"`
// 	Station         string `json:"station"`
// 	ProductionCount int64  `json:"production_count"`
// 	OEE             OEE    `json:"oee"`

// 	CycleTimeSec   float64   `json:"cycle_time_sec"`
// 	ProductionTime time.Time `json:"production_time"`
// 	CreatedAt      time.Time `json:"created_at"`
// }

type ShiftInfo struct {
	ShiftID   int64
	ShiftName string

	ShiftStart time.Time
	ShiftEnd   time.Time
}

// type ShiftInfo1 struct {
// 	ShiftID    int64
// 	ShiftStart time.Time
// 	ShiftEnd   time.Time
// }
