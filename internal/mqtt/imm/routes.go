package imm

import (
	"log"

	paho "github.com/eclipse/paho.mqtt.golang"
)

// ============================================================
// REGISTER MQTT ROUTES
// ============================================================

func (m *Module) RegisterMQTTIMMRoutes(
	client paho.Client,
	subscribe func(
		paho.Client,
		string,
		paho.MessageHandler,
	),
) {

	log.Println("------------------------------------------------")
	log.Println("Registering IMM MQTT Routes...")
	log.Println("------------------------------------------------")

	// ========================================================
	// PRODUCTION
	// ========================================================

	subscribe(
		client,
		TopicProduction,
		m.Handlers.Production.ProductionHandler(),
	)

	log.Printf(
		"✅ IMM Production Route Registered | Topic=%s",
		TopicProduction,
	)

	// ========================================================
	// HEARTBEAT
	// ========================================================

	subscribe(
		client,
		TopicHeartbeat,
		m.Handlers.Heartbeat.HeartbeatHandler(),
	)

	log.Printf(
		"✅ IMM Heartbeat Route Registered | Topic=%s",
		TopicHeartbeat,
	)

	log.Println("------------------------------------------------")
	log.Println("IMM MQTT Routes Registered")
	log.Println("------------------------------------------------")
}

// package imm

// import (
// 	"log"

// 	paho "github.com/eclipse/paho.mqtt.golang"
// )

// // ============================================================
// // REGISTER MQTT ROUTES
// // ============================================================

// func (m *Module) RegisterMQTTIMMRoutes(
// 	client paho.Client,
// 	subscribe func(
// 		paho.Client,
// 		string,
// 		paho.MessageHandler,
// 	),
// ) {

// 	log.Println("------------------------------------------------")
// 	log.Println("Registering IMM MQTT Routes...")
// 	log.Println("------------------------------------------------")

// 	// ========================================================
// 	// PRODUCTION
// 	// ========================================================

// 	subscribe(
// 		client,
// 		TopicProduction,
// 		m.Handlers.Production.ProductionHandler(),
// 	)

// 	log.Printf(
// 		"✅ IMM Production Route Registered | Topic=%s",
// 		TopicProduction,
// 	)

// 	// ========================================================
// 	// HEARTBEAT
// 	// ========================================================

// 	subscribe(
// 		client,
// 		TopicHeartbeat,
// 		m.Handlers.Heartbeat.HeartbeatHandler(),
// 	)

// 	log.Printf(
// 		"✅ IMM Heartbeat Route Registered | Topic=%s",
// 		TopicHeartbeat,
// 	)

// 	log.Println("------------------------------------------------")
// 	log.Println("IMM MQTT Routes Registered")
// 	log.Println("------------------------------------------------")
// }
