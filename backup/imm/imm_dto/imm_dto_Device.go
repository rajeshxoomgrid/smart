package immdto

import "time"

// ============================================================
// DEVICE
// ============================================================

type Device struct {
	DeviceID   string
	MachineID  string
	TenantID   int64
	DeviceType string
	Status     string
	MoldNo     string
}

// ============================================================
// MOLD CHANGE COMMAND
// ============================================================

type MoldChangeCommand struct {
	ID              int64
	TenantID        int64
	UserID          int64
	DeviceID        string
	MachineID       string
	OldMoldNo       string
	RequestedMoldNo string
	Status          string
	FailureReason   string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	ConfirmedAt     time.Time
}
