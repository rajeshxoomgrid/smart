package handler

import (
	"encoding/json"
	"fmt"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"
)

const TopicProductionAck = "factory/imm/production/ack/"

type productionAck struct {
	EventID string `json:"event_id"`
	Status  string `json:"status"`
}

func (h *ProductionHandler) publishSavedAck(
	client paho.Client,
	deviceID string,
	eventID string,
) error {
	if client == nil {
		return fmt.Errorf("MQTT client is nil")
	}
	if deviceID == "" || eventID == "" {
		return fmt.Errorf("device_id and event_id are required")
	}

	payload, err := json.Marshal(productionAck{
		EventID: eventID,
		Status:  "saved",
	})
	if err != nil {
		return fmt.Errorf("marshal production ACK: %w", err)
	}

	topic := TopicProductionAck + deviceID
	token := client.Publish(topic, 1, false, payload)

	if !token.WaitTimeout(5 * time.Second) {
		return fmt.Errorf("production ACK publish timed out")
	}
	if err := token.Error(); err != nil {
		return fmt.Errorf("publish production ACK: %w", err)
	}

	return nil
}
