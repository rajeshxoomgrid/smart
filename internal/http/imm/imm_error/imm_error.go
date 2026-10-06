package immerror

import "errors"

var (
	ErrDeviceNotFound          = errors.New("device not found")
	ErrNotIMMDevice            = errors.New("device is not an IMM device")
	ErrPendingCommandNotFound  = errors.New("pending mold change command not found")
	ErrInvalidDevice           = errors.New("invalid device")
	ErrSameMoldNo              = errors.New("requested mold number is already configured")
	ErrCommandPending          = errors.New("a mold change command is already pending")
	ErrInvalidACK              = errors.New("invalid mold change acknowledgement")
	ErrMoldMismatch            = errors.New("acknowledged mold number does not match requested mold number")
	ErrCommandAlreadyCompleted = errors.New("Command Already completed")
)
