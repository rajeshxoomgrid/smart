package handler

import (
	"log"
	"time"
)

type ProductionStats struct {
	QueueLength int
	QueueSize   int
	Workers     int
	Received    uint64
	Queued      uint64
	Processed   uint64
	Failed      uint64
	Dropped     uint64
	Retries     uint64
}

func (h *ProductionHandler) Stats() ProductionStats {
	return ProductionStats{
		QueueLength: len(h.queue),
		QueueSize:   cap(h.queue),
		Workers:     int(h.workerCount.Load()),
		Received:    h.receivedCount.Load(),
		Queued:      h.queuedCount.Load(),
		Processed:   h.processedCount.Load(),
		Failed:      h.failedCount.Load(),
		Dropped:     h.droppedCount.Load(),
		Retries:     h.retryCount.Load(),
	}
}

func (h *ProductionHandler) statisticsLogger() {
	defer h.supervisorWG.Done()

	ticker := time.NewTicker(ProductionStatsInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			h.logStatistics()

		case <-h.ctx.Done():
			return
		}
	}
}

func (h *ProductionHandler) logStatistics() {
	stats := h.Stats()

	log.Printf(
		"IMM Production Stats | Queue=%d/%d | Workers=%d | Received=%d | Queued=%d | Processed=%d | Failed=%d | Dropped=%d | Retries=%d",
		stats.QueueLength,
		stats.QueueSize,
		stats.Workers,
		stats.Received,
		stats.Queued,
		stats.Processed,
		stats.Failed,
		stats.Dropped,
		stats.Retries,
	)
}
