package kafka

import (
	"log"
	"os"

	"github.com/IBM/sarama"
)

func RunProducer(message string) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	producer, err := sarama.NewSyncProducer([]string{os.Getenv("KAFKA_BOOTSTRAP_SERVER")}, config)
	if err != nil {
		log.Panicf("Failed to start Sarama producer: %v", err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Panicf("Failed to close Sarama producer: %v", err)
		}
	}()

	topic := "Votes_Added"
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Panicf("Failed to send message to Kafka: %v", err)
	} else {
		log.Printf("Message sent to topic %s, partition %d, offset %d\n", topic, partition, offset)
	}
}
