package dto

import "time"

type ImmDataFetchRequest struct {
	DeviceID  string `json:"device_id"`
	MachineID string `json:"machine_id"`
	Limit     int    `json:"limit"`
}

type ImmDataFetechResponse struct {
	ID                     int64     `json:"id"`
	EventID                string    `json:"event_id"`
	TenantID               int       `json:"tenant_id"`
	DeviceID               string    `json:"device_id"`
	MachineID              string    `json:"machine_id"`
	MoldNo                 string    `json:"mold_no"`
	PartName               string    `json:"part_name"`
	ProductionCount        int64     `json:"production_count"`
	CycleTimeSec           float64   `json:"cycle_time_sec"`
	ProductionTime         time.Time `json:"production_time"`
	Reset                  bool      `json:"reset"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
	ShiftName              string    `json:"shift_name"`
	CurrentProductionCount int64     `json:"current_production_count"`
	DayProductionCount     int64     `json:"day_production_count"`
	OEE                    OEE       `json:"oee"`
	DayOEE                 OEE       `json:"day_oee"`
}

type ImmDataResponse struct {
	ID              int64     `json:"id"`
	EventID         string    `json:"event_id"`
	TenantID        int       `json:"tenant_id"`
	DeviceID        string    `json:"device_id"`
	MachineID       string    `json:"machine_id"`
	MoldNo          string    `json:"mold_no"`
	ProductionCount int64     `json:"production_count"`
	CycleTimeSec    float64   `json:"cycle_time_sec"`
	ProductionTime  time.Time `json:"production_time"`
	Reset           bool      `json:"reset"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type ProductionLogResponse struct {
	Data            []ImmDataResponse `json:"data"`
	MinCycleTimeSec float64           `json:"min_cycle_time_sec"`
	MaxCycleTimeSec float64           `json:"max_cycle_time_sec"`
	AvgCycleTimeSec float64           `json:"avg_cycle_time_sec"`
}
