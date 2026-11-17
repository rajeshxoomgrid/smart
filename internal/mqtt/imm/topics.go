package imm

const (
	// ESP32 → Go
	TopicProduction    = "factory/imm/production/+"
	TopicHeartbeat     = "factory/imm/heartbeat/+"
	TopicStatus        = "factory/imm/status/+"
	TopicAlarm         = "factory/imm/alarm/+"
	TopicProductionAck = "factory/imm/ack/+"

	// Go subscribes to device command responses/topics if needed
	TopicCommand = "factory/imm/command/+"
)

// const (
// 	TopicProduction = "factory/imm/production/+"

// 	TopicHeartbeat = "factory/imm/heartbeat/+"

// 	TopicStatus = "factory/imm/status/+"

// 	TopicAlarm = "factory/imm/alarm/+"

// 	TopicCommand = "factory/imm/command/"

// 	TopicProductionAck = "factory/imm/ack/"
// )
