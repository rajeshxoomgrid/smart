package handler

import "errors"

// ============================================================
// PRODUCTION HANDLER ERRORS
// ============================================================

var (
	ErrProductionHandlerClosed = errors.New(
		"production handler is closed",
	)

	ErrProductionQueueFull = errors.New(
		"production queue is full",
	)
)
