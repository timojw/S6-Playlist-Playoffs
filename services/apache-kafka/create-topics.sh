#!/bin/bash

# Create the Kafka topic
/opt/bitnami/kafka/bin/kafka-topics.sh --create --topic GetVotes --partitions 1 --replication-factor 1 --bootstrap-server kafka:9092

echo "Topics are created"
