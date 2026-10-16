package mqtt

import (
	"log"

	paho "github.com/eclipse/paho.mqtt.golang"

	"github.com/rajeshbond/smart/database"
	"github.com/rajeshbond/smart/internal/mqtt/assembly"
	"github.com/rajeshbond/smart/internal/mqtt/imm"
)

// TokenAuth := service.GetTokenAuth()
// ============================================================
// REGISTER MQTT ROUTES
// ============================================================

func RegisterRoutes(client paho.Client, db *database.DB) {

	log.Println("========================================")
	log.Println("Registering MQTT Modules...")
	log.Println("========================================")

	// Create Modules once

	// --------------------------------------------------------
	// IMM MODULE
	// --------------------------------------------------------

	immModule := imm.NewModule(db.SQLDB)

	// --------------------------------------------------------
	// ASSEMBLY MODULE
	// --------------------------------------------------------

	assemblyModule := assembly.NewModule(db.SQLDB)

	// Register subscription

	// --------------------------------------------------------
	// IMM MQTT ROUTES
	// --------------------------------------------------------

	immModule.RegisterMQTTIMMRoutes(
		client,
		subscribe,
	)

	// --------------------------------------------------------
	// ASSEMBLY MQTT ROUTES
	// --------------------------------------------------------

	assemblyModule.RegisterMQTTAssemblyRoute(
		client,
		subscribe,
	)

	log.Println("========================================")
	log.Println("MQTT Route Registration Complete")
	log.Println("========================================")
}

// ============================================================
// REGISTER BACKGROUND WORKERS
// ============================================================

func RegisterWorkers(db *database.DB) {

	log.Println(
		"🐕 Initializing background factory hardware watchdogs...",
	)

	// Boot up IMM Watchdog (30 Second Timeout)
	// immStore := imm.NewStore(db.PGX)
	// imm.StartWatchdog(immStore, 30*time.Second)

	// Boot up Assembly Watchdog if needed (1 Minute Timeout)
	// assemblyStore := assembly.NewStore(db)
	// assembly.StartWatchdog(assemblyStore, 60*time.Second)

	log.Println(
		"✅ All background factory watchdogs are active",
	)
}

// package mqtt

// import (
// 	"log"

// 	paho "github.com/eclipse/paho.mqtt.golang"

// 	"github.com/rajeshbond/smart/database"
// 	"github.com/rajeshbond/smart/internal/mqtt/assembly"
// 	"github.com/rajeshbond/smart/internal/mqtt/imm"
// )

// func RegisterRoutes(client paho.Client, db *database.DB) {
// 	log.Println("========================================")
// 	log.Println("Registering MQTT Modules...")
// 	log.Println("========================================")

// 	// Create Modules once

// 	assemblyModule := assembly.NewModule(db.SQLDB)
// 	immModule := imm.NewModule(db.SQLDB)

// 	// Register subscription

// 	// immModule.RegisterMQTTIMMRoute(client, subscribe)
// 	assemblyModule.RegisterMQTTAssemblyRoute(client, subscribe)
// 	immModule.RegisterMQTTIMMRoutes(client, subscribe)

// 	log.Println("========================================")
// 	log.Println("MQTT Route Registration Complete")
// 	log.Println("========================================")

// }

// func RegisterWorkers(db *database.DB) {
// 	log.Println("🐕 Initializing background factory hardware watchdogs...")

// 	// Boot up IMM Watchdog (30 Second Timeout)
// 	// immStore := imm.NewStore(db.PGX)
// 	// imm.StartWatchdog(immStore, 30*time.Second)

// 	// Boot up Assembly Watchdog if needed (1 Minute Timeout)
// 	// assemblyStore := assembly.NewStore(db)
// 	// assembly.StartWatchdog(assemblyStore, 60*time.Second)

// 	log.Println("✅ All background factory watchdogs are active")
// }
