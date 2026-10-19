package dto

type GetProductionYesterdayRequest struct {
	DeviceID string `json:"device_id"`
	Station  string `json:"station,omitempty"`
	Variant  string `json:"variant,omitempty"`
	Date     string `json:"date"`
}

type ProductionShiftDTO struct {
	Shift           string `json:"shift"`
	ProductionCount int64  `json:"production_count"`
}

type GetYesterdaysProductionResponse struct {
	TenantID string               `json:"tenant_id"`
	DeviceID string               `json:"device_id"`
	Station  string               `json:"station,omitempty"`
	Variant  string               `json:"variant,omitempty"`
	Date     string               `json:"date"`
	Shifts   []ProductionShiftDTO `json:"shifts"`
	Total    int64                `json:"total_production"`
}

// type GetProductionYesterdayRequest struct {
// 	DeviceID string `json:"device_id"`
// 	Station  string `json:"station,omitempty"`
// 	Variant  string `json:"variant,omitempty"`
// 	Date     string `json:"date"`
// }

// type ProductionShiftDTO struct {
// 	Shift           string `json:"shift"`
// 	StartTime       string `json:"start_time"`
// 	EndTime         string `json:"end_time"`
// 	ProductionCount int64  `json:"production_count"`
// }

// type GetYesterdaysProductionResponse struct {
// 	TenantID string               `json:"tenant_id"`
// 	DeviceID string               `json:"device_id"`
// 	Station  string               `json:"station,omitempty"`
// 	Variant  string               `json:"variant,omitempty"`
// 	Date     string               `json:"date"`
// 	Shifts   []ProductionShiftDTO `json:"shifts"`
// 	Total    int64                `json:"total_production"`
// }
