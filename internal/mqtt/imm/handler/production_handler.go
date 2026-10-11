package handler

import (
	"context"
	"log"
	"sync"
	"sync/atomic"

	paho "github.com/eclipse/paho.mqtt.golang"

	service "github.com/rajeshbond/smart/internal/mqtt/imm/service"
)

type productionJob struct {
	payload []byte
	client  paho.Client
}

type ProductionHandler struct {
	service service.ImmService

	queue chan productionJob

	ctx    context.Context
	cancel context.CancelFunc

	mu        sync.RWMutex
	accepting bool

	workersWG sync.WaitGroup

	workerCount atomic.Int32

	receivedCount  atomic.Uint64
	queuedCount    atomic.Uint64
	processedCount atomic.Uint64
	failedCount    atomic.Uint64
	droppedCount   atomic.Uint64
	retryCount     atomic.Uint64

	supervisorWG sync.WaitGroup
}

func NewProductionHandler(
	productionService *service.ImmService,
) *ProductionHandler {
	if productionService == nil {
		panic("productionService cannot be nil")
	}

	ctx, cancel := context.WithCancel(context.Background())

	h := &ProductionHandler{
		service:   *productionService,
		queue:     make(chan productionJob, ProductionQueueSize),
		ctx:       ctx,
		cancel:    cancel,
		accepting: true,
	}

	for i := 0; i < MinProductionWorkers; i++ {
		h.startWorker()
	}

	h.supervisorWG.Add(1)
	go h.workerSupervisor()

	h.supervisorWG.Add(1)
	go h.statisticsLogger()

	log.Printf(
		"IMM Production Handler Started | Queue=%d | MinWorkers=%d | MaxWorkers=%d",
		ProductionQueueSize,
		MinProductionWorkers,
		MaxProductionWorkers,
	)

	return h
}

// Shutdown stops accepting new messages and stops background workers.
// This retains the original cancellation-based shutdown behavior.
func (h *ProductionHandler) Shutdown() {
	h.mu.Lock()
	h.accepting = false
	h.mu.Unlock()

	h.cancel()

	h.supervisorWG.Wait()
	h.workersWG.Wait()

	log.Println("IMM Production Handler Shutdown Complete")
}
