package handler

import (
	"log"
	"time"
)

func (h *ProductionHandler) workerSupervisor() {
	defer h.supervisorWG.Done()

	ticker := time.NewTicker(WorkerScaleInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			h.scaleWorkers()

		case <-h.ctx.Done():
			return
		}
	}
}

func (h *ProductionHandler) scaleWorkers() {
	queueLength := len(h.queue)
	queueCapacity := cap(h.queue)

	if queueCapacity == 0 || queueLength == 0 {
		return
	}

	h.mu.RLock()
	accepting := h.accepting
	h.mu.RUnlock()

	if !accepting || h.ctx.Err() != nil {
		return
	}

	workers := int(h.workerCount.Load())
	if workers >= MaxProductionWorkers {
		return
	}

	utilization := float64(queueLength) / float64(queueCapacity)

	addWorkers := 0

	switch {
	case utilization >= 0.75:
		addWorkers = 4
	case utilization >= 0.50:
		addWorkers = 3
	case utilization >= 0.25:
		addWorkers = 2
	case utilization >= 0.10:
		addWorkers = 1
	}

	targetWorkers := workers + addWorkers
	if targetWorkers > MaxProductionWorkers {
		targetWorkers = MaxProductionWorkers
	}

	for workers < targetWorkers {
		h.mu.RLock()
		accepting = h.accepting
		h.mu.RUnlock()

		if !accepting || h.ctx.Err() != nil {
			return
		}

		// startWorker increments workerCount itself.
		h.startWorker()
		workers++

		log.Printf(
			"IMM Production Worker Scaling | Workers=%d | Queue=%d/%d",
			workers,
			queueLength,
			queueCapacity,
		)
	}
}
