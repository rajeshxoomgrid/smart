package handler

import "time"

// ============================================================
// IMM PRODUCTION HANDLER CONFIGURATION
// ============================================================

const (
	// --------------------------------------------------------
	// Queue
	// --------------------------------------------------------

	ProductionQueueSize = 1000

	// --------------------------------------------------------
	// Workers
	// --------------------------------------------------------

	MinProductionWorkers = 2
	MaxProductionWorkers = 20

	WorkerScaleInterval = 500 * time.Millisecond

	WorkerIdleTimeout = 30 * time.Second

	// --------------------------------------------------------
	// Database
	// --------------------------------------------------------

	ProductionDBTimeout = 5 * time.Second

	ProductionSaveRetries = 3

	ProductionRetryDelay = 500 * time.Millisecond

	// --------------------------------------------------------
	// MQTT enqueue
	// --------------------------------------------------------

	ProductionEnqueueTimeout = 2 * time.Second

	// --------------------------------------------------------
	// Statistics
	// --------------------------------------------------------

	ProductionStatsInterval = 30 * time.Second
)
