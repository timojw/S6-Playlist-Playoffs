package kafka

import (
	"encoding/json"
	"log"
	"os"

	"github.com/IBM/sarama"
)

func RunProducer(gameID int, getVotes bool) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	producer, err := sarama.NewSyncProducer([]string{os.Getenv("KAFKA_BOOTSTRAP_SERVER")}, config)
	if err != nil {
		log.Panicf("Failed to start Sarama producer: %v", err)
	}
	defer producer.Close()

	message := map[string]interface{}{
		"GameID":   gameID,
		"GetVotes": getVotes,
	}
	messageBytes, err := json.Marshal(message)
	if err != nil {
		log.Panicf("Failed to marshal message: %v", err)
	}

	topic := "GET_VOTES"
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(messageBytes),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Panicf("Failed to send message to Kafka: %v", err)
	} else {
		log.Printf("Message sent to topic %s, partition %d, offset %d", topic, partition, offset)
	}
}
