package kafka

import (
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/timojw/S6-Playlist-Playoffs/services/game/internal/models"
	// services "gitlab.com/validation/internal/services"
)

func handleMessage(kafkaMessage sarama.ConsumerMessage) {
	var message = models.ConsumeMessage{}

	// Convert message to json model
	err := json.Unmarshal(kafkaMessage.Value, &message)
	if err != nil {
		fmt.Printf("Failed to unmarshal Kafka message: %v\n", err)
		return
	}

	// Handle the parameter
	// services.ValidateParameter(message)
}
