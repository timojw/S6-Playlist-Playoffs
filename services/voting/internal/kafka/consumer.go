package kafka

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/sarama"
)

func RunConsumer() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.AutoCommit.Enable = false

	var brokers = []string{os.Getenv("KAFKA_BOOTSTRAP_SERVER")}
	var consumerGroupId = "VotesServiceGroup"

	consumerGroup, err := setupConsumerGroup(brokers, consumerGroupId, config)
	if err != nil {
		log.Panicf("Error setting up consumer: %v", err)
	}

	consumer := Consumer{}
	ctx := context.Background()
	var topicName = "Votes_Added"

	println("Listening to topic:", topicName)

	for {
		err := consumerGroup.Consume(ctx, []string{topicName}, &consumer)
		if err != nil {
			log.Printf("Error from consumer: %v", err)
			log.Println("Retrying in 3 seconds...")
			time.Sleep(3 * time.Second)
			continue
		}
	}
}

type Consumer struct{}

func (consumer Consumer) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (consumer Consumer) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (consumer Consumer) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d value:%s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Value))

		switch msg.Topic {
		case "Votes_Added":
			handleMessage(*msg)
		default:
			fmt.Println("Unknown topic:", msg.Topic)
		}

		sess.MarkMessage(msg, "")
		sess.Commit()
	}
	return nil
}

func setupConsumerGroup(brokers []string, consumerGroupId string, config *sarama.Config) (sarama.ConsumerGroup, error) {
	var consumerGroup sarama.ConsumerGroup
	var err error

	for {
		consumerGroup, err = sarama.NewConsumerGroup(brokers, consumerGroupId, config)
		if err != nil {
			log.Printf("Error creating consumer group: %v", err)
			log.Println("Retrying in 3 seconds...")
			time.Sleep(3 * time.Second)
		} else {
			break
		}
	}

	return consumerGroup, err
}
