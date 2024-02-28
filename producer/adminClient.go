package main

import (
	"context"
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/pkg/errors"
)

const (
	BootstrapServers      = "broker1"
	adminOperationTimeout = time.Second * 60
)

var topicsSpec = []kafka.TopicSpecification{
	{Topic: "1", NumPartitions: 10, ReplicationFactor: 3},
	// {Topic: "2", NumPartitions: 5, ReplicationFactor: 1},
	// {Topic: "3", NumPartitions: 3, ReplicationFactor: 2},
	// {Topic: "4", NumPartitions: 4, ReplicationFactor: 2},
	// {Topic: "5", NumPartitions: 5, ReplicationFactor: 3},
	// {Topic: "6", NumPartitions: 3, ReplicationFactor: 3},
	// {Topic: "7", NumPartitions: 9, ReplicationFactor: 2},
}

func configureAdminClient() ([]string, error) {
	fmt.Println("creating admin client")
	adminClient, err := kafka.NewAdminClient(&kafka.ConfigMap{
		"bootstrap.servers": BootstrapServers,
	})
	if err != nil {
		return nil, errors.Wrap(err, "error creating new admin")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fmt.Println("creating topics")
	results, err := adminClient.CreateTopics(
		ctx,
		topicsSpec,
		kafka.SetAdminOperationTimeout(adminOperationTimeout))
	if err != nil {
		return nil, errors.Wrap(err, "couldnt create topics")
	}

	// Print results
	for _, result := range results {
		fmt.Printf("%s\n", result)
	}

	adminClient.Close()

	return fromTopicSpecsToTopics(topicsSpec), nil
}

func fromTopicSpecsToTopics(topicsSpecs []kafka.TopicSpecification) []string {
	var topics []string
	for _, topicSpec := range topicsSpecs {
		topics = append(topics, topicSpec.Topic)
	}
	return topics
}
