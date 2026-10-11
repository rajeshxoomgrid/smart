package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	"strconv"
	"time"

	dto "github.com/rajeshbond/smart/internal/mqtt/imm/dto"
)

// START WORKER

func (h *ProductionHandler) startWorker() {
	workerID := int(h.workerCount.Add(1))

	h.workersWG.Add(1)

	go h.worker(workerID)

	log.Printf(
		"IMM Production Worker Started | Worker=%d | Workers=%d",
		workerID,
		h.workerCount.Load(),
	)
}

// WORKER

func (h *ProductionHandler) worker(workerID int) {
	defer h.workersWG.Done()

	defer func() {
		h.workerCount.Add(-1)

		log.Printf(
			"IMM Production Worker Stopped | Worker=%d | Workers=%d",
			workerID,
			h.workerCount.Load(),
		)
	}()

	idleTimer := time.NewTimer(WorkerIdleTimeout)
	defer idleTimer.Stop()

	for {
		if h.ctx.Err() != nil {
			return
		}

		select {
		case job := <-h.queue:
			if !idleTimer.Stop() {
				select {
				case <-idleTimer.C:
				default:
				}
			}

			idleTimer.Reset(WorkerIdleTimeout)

			h.processProduction(workerID, job)

		case <-idleTimer.C:
			if h.ctx.Err() != nil {
				return
			}

			if h.workerCount.Load() > int32(MinProductionWorkers) {
				return
			}

			idleTimer.Reset(WorkerIdleTimeout)

		case <-h.ctx.Done():
			return
		}
	}
}

// NORMALIZE TENANT ID
//
// Supports both:
//     "tenant_id": "2"  (old SD-card events)
//     "tenant_id": 2    (new ESP32 events)
//
// Converts the value to a JSON integer before DTO unmarshalling.

func normalizeTenantID(payload []byte) ([]byte, error) {
	var fields map[string]json.RawMessage

	if err := json.Unmarshal(payload, &fields); err != nil {
		return nil, fmt.Errorf("decode production JSON: %w", err)
	}

	rawTenantID, exists := fields["tenant_id"]
	if !exists {
		return nil, fmt.Errorf("tenant_id is missing")
	}

	rawTenantID = bytes.TrimSpace(rawTenantID)

	// Already an integer in JSON.
	var numericID int
	if err := json.Unmarshal(rawTenantID, &numericID); err == nil {
		if numericID <= 0 {
			return nil, fmt.Errorf("tenant_id must be greater than zero")
		}

		return payload, nil
	}

	// Legacy SD event: tenant_id is a JSON string.
	var tenantIDString string
	if err := json.Unmarshal(rawTenantID, &tenantIDString); err != nil {
		return nil, fmt.Errorf(
			"tenant_id must be an integer or numeric string",
		)
	}

	numericID, err := strconv.Atoi(tenantIDString)
	if err != nil || numericID <= 0 {
		return nil, fmt.Errorf(
			"invalid tenant_id string %q",
			tenantIDString,
		)
	}

	// Replace the JSON string with a JSON integer.
	fields["tenant_id"], err = json.Marshal(numericID)
	if err != nil {
		return nil, fmt.Errorf("encode tenant_id: %w", err)
	}

	normalizedPayload, err := json.Marshal(fields)
	if err != nil {
		return nil, fmt.Errorf("encode normalized production JSON: %w", err)
	}

	return normalizedPayload, nil
}

// PROCESS PRODUCTION

func (h *ProductionHandler) processProduction(
	workerID int,
	job productionJob,
) {
	defer func() {
		if r := recover(); r != nil {
			h.failedCount.Add(1)

			log.Printf(
				"IMM Production Worker Panic | Worker=%d | Panic=%v",
				workerID,
				r,
			)

			log.Println(stackTrace())
		}
	}()

	log.Printf(
		"IMM Production Raw Payload | Worker=%d | Payload=%s",
		workerID,
		string(job.payload),
	)

	// Normalize tenant_id before unmarshalling into the DTO.
	normalizedPayload, err := normalizeTenantID(job.payload)
	if err != nil {
		h.failedCount.Add(1)

		log.Printf(
			"IMM Production Tenant ID Error | Worker=%d | Error=%v",
			workerID,
			err,
		)
		return
	}

	var req dto.ProductionDTO

	if err := json.Unmarshal(normalizedPayload, &req); err != nil {
		h.failedCount.Add(1)

		log.Printf(
			"IMM Production JSON Error | Worker=%d | Error=%v",
			workerID,
			err,
		)
		return
	}

	if req.EventID == "" ||
		req.TenantID <= 0 ||
		req.DeviceID == "" ||
		req.MachineID == "" ||
		req.Timestamp.IsZero() ||
		req.Count < 0 ||
		req.CycleTimeSec < 0 {

		h.failedCount.Add(1)

		log.Printf(
			"IMM Production Validation Error | Worker=%d | EventID=%q | TenantID=%d | DeviceID=%q | MachineID=%q | Timestamp=%v | Count=%d | CycleTimeSec=%f",
			workerID,
			req.EventID,
			req.TenantID,
			req.DeviceID,
			req.MachineID,
			req.Timestamp,
			req.Count,
			req.CycleTimeSec,
		)

		return
	}

	start := time.Now()

	if err := h.saveWithRetry(&req); err != nil {
		h.failedCount.Add(1)

		log.Printf(
			"IMM Production Save FAILED | Worker=%d | EventID=%s | Error=%v",
			workerID,
			req.EventID,
			err,
		)
		return
	}

	h.processedCount.Add(1)

	log.Printf(
		"IMM Production Saved | Worker=%d | EventID=%s | Device=%s | Duration=%v",
		workerID,
		req.EventID,
		req.DeviceID,
		time.Since(start),
	)

	// Publish ACK only after a successful save.
	if err := h.publishSavedAck(
		job.client,
		req.DeviceID,
		req.EventID,
	); err != nil {
		log.Printf(
			"IMM Production ACK Failed | Worker=%d | EventID=%s | Error=%v",
			workerID,
			req.EventID,
			err,
		)
	}
}

// STACK TRACE

func stackTrace() string {
	buf := make([]byte, 64*1024)
	n := runtime.Stack(buf, false)

	return fmt.Sprintf("%s", buf[:n])
}
