package dto

type HrSaveProductionRequest struct {
	DeviceID string `json:"device_id"`
	Station  string `json:"station,omitempty"`
	Variant  string `json:"variant,omitempty"`
}
