package immconstant

import "time"

const (
	StatusPending = "PENDING"
	StatusSuccess = "SUCCESS"
	StatusFailed  = "FAILED"

	CommandSetMoldNo = "SET_MOLD_NO"

	MoldChangeTimeout = 30 * time.Second

	TopicIMMCommand = "factory/imm/command/"
)
