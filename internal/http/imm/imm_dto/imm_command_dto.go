package immdto

// ============================================================
// HTTP REQUEST
// ============================================================

type SetMoldNoRequest struct {
	DeviceID string `json:"device_id" validate:"required,max=100"`
	MoldNo   string `json:"mold_no" validate:"required,max=100"`
}

// ============================================================
// HTTP RESPONSE
// ============================================================

type SetMoldNoResponse struct {
	Message  string `json:"message"`
	DeviceID string `json:"device_id"`
	MoldNo   string `json:"mold_no"`
	Status   string `json:"status"`
}

// ============================================================
// STATUS RESPONSE
// ============================================================

type MoldStatusResponse struct {
	DeviceID        string `json:"device_id"`
	CurrentMoldNo   string `json:"current_mold_no"`
	RequestedMoldNo string `json:"requested_mold_no,omitempty"`
	Status          string `json:"status"`
	FailureReason   string `json:"failure_reason,omitempty"`
	CreatedAt       string `json:"created_at,omitempty"`
}

// ============================================================
// MQTT COMMAND
// ============================================================

type SetMoldNoMQTTRequest struct {
	DeviceID string `json:"device_id"`
	Command  string `json:"command"`
	MoldNo   string `json:"mold_no"`
}

// ============================================================
// MQTT ACK
// ============================================================

type SetMoldNoMQTTAck struct {
	DeviceID string `json:"device_id"`
	Command  string `json:"command"`
	Status   string `json:"status"`
	MoldNo   string `json:"mold_no,omitempty"`
	Error    string `json:"error,omitempty"`
}
