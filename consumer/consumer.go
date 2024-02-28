package main

import "github.com/confluentinc/confluent-kafka-go/v2/kafka"

const (
	BootstrapServers  = "localhost"
	group             = "group1"
	ipv4              = "v4"
	connectionTimeOut = 6000
)

func getConsumer() (*kafka.Consumer, error) {
	return kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":     BootstrapServers,
		"broker.address.family": ipv4,
		"group.id":              group,
		"session.timeout.ms":    connectionTimeOut,
		// Start reading from the first message of each assigned
		// partition if there are no previously committed offsets
		// for this group.
		"auto.offset.reset":        "earliest",
		"enable.auto.offset.store": false,
	})
}
