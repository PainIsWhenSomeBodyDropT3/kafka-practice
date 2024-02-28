package main

import "github.com/confluentinc/confluent-kafka-go/v2/kafka"

const (
	asks     = "all"
	clientID = "producer"
)

func configureProducer() (*kafka.Producer, error) {
	return kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": BootstrapServers,
		"acks":              asks,
		"client.id":         clientID,
	})
}
