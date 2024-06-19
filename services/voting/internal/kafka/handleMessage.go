package kafka

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/timojw/S6-Playlist-Playoffs/services/voting/internal/models"
)

func handleMessage(kafkaMessage sarama.ConsumerMessage) {
	var message models.ConsumeMessage

	// Convert message to JSON model
	err := json.Unmarshal(kafkaMessage.Value, &message)
	if err != nil {
		fmt.Printf("Failed to unmarshal Kafka message: %v\n", err)
		return
	}

	log.Printf("Received message: %+v\n", message)
	// Handle the parameter (uncomment and implement this line according to your business logic)
	// services.ValidateParameter(message)
}
