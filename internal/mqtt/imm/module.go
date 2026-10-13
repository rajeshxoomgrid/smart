package imm

import (
	"database/sql"

	"github.com/rajeshbond/smart/internal/mqtt/imm/handler"
	"github.com/rajeshbond/smart/internal/mqtt/imm/service"
	"github.com/rajeshbond/smart/internal/mqtt/imm/store"
)

// ============================================================
// MQTT HANDLERS
// ============================================================

type MQTTHandlers struct {
	Production *handler.ProductionHandler
	Heartbeat  *handler.HeartbeatHandler
}

// ============================================================
// IMM MODULE
// ============================================================

type Module struct {
	Store    *store.ImmProductionStore
	Service  service.ImmService
	Handlers MQTTHandlers
}

// ============================================================
// NEW IMM MODULE
// ============================================================

func NewModule(db *sql.DB) *Module {

	// --------------------------------------------------------
	// Validate DB
	// --------------------------------------------------------

	if db == nil {
		panic("IMM database connection cannot be nil")
	}

	// --------------------------------------------------------
	// Production Store
	// --------------------------------------------------------

	productionStore := store.NewImmProductionStore(db)

	// --------------------------------------------------------
	// IMM Service
	// --------------------------------------------------------

	immService := service.NewImmService(
		productionStore,
	)

	// --------------------------------------------------------
	// Production Handler
	// --------------------------------------------------------

	productionHandler := handler.NewProductionHandler(
		&immService,
	)

	// --------------------------------------------------------
	// Heartbeat Handler
	// --------------------------------------------------------

	heartbeatHandler := handler.NewHeartbeatHandler()

	// --------------------------------------------------------
	// Create IMM Module
	// --------------------------------------------------------

	return &Module{

		Store: productionStore,

		Service: immService,

		Handlers: MQTTHandlers{

			Production: productionHandler,

			Heartbeat: heartbeatHandler,
		},
	}
}
