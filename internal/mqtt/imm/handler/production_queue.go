package handler

import (
	"log"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"
)

// var (
// 	ErrProductionHandlerClosed = errors.New("production handler is closed")
// 	ErrProductionQueueFull     = errors.New("production queue is full")
// )

func (h *ProductionHandler) ProductionHandler() paho.MessageHandler {
	return func(client paho.Client, msg paho.Message) {
		if msg == nil {
			log.Println("IMM Production message is nil")
			return
		}

		h.receivedCount.Add(1)

		job := productionJob{
			payload: append([]byte(nil), msg.Payload()...),
			client:  client,
		}

		if err := h.enqueue(job); err != nil {
			h.droppedCount.Add(1)
			log.Printf("IMM Production Enqueue Failed | Error=%v", err)
			return
		}

		h.queuedCount.Add(1)
	}
}

func (h *ProductionHandler) enqueue(job productionJob) error {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if !h.accepting {
		return ErrProductionHandlerClosed
	}

	timer := time.NewTimer(ProductionEnqueueTimeout)
	defer timer.Stop()

	select {
	case h.queue <- job:
		return nil

	case <-timer.C:
		return ErrProductionQueueFull

	case <-h.ctx.Done():
		return ErrProductionHandlerClosed
	}
}
