package handler

import "time"

const (
	ProductionQueueSize      = 1000
	MinProductionWorkers     = 5
	MaxProductionWorkers     = 20
	WorkerIdleTimeout        = 60 * time.Second
	WorkerScaleInterval      = 5 * time.Second
	ProductionSaveRetries    = 3
	ProductionDBTimeout      = 10 * time.Second
	ProductionRetryDelay     = 500 * time.Millisecond
	ProductionEnqueueTimeout = 250 * time.Millisecond
	ProductionStatsInterval  = 30 * time.Second
)
