package main

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var topics = []string{"1"}

const MIN_COMMIT_COUNT = 5

func main() {
	c, err := getConsumer()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created Consumer %v\n", c)
	err = c.SubscribeTopics(topics, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to subscribe to topics: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Subscribed to topics %v\n", topics)
	consumeMessages(c)
}

func consumeMessages(consumer *kafka.Consumer) {
	msg_count := 0
	for {
		ev := consumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			msg_count += 1
			if msg_count%MIN_COMMIT_COUNT == 0 {
				_, err := consumer.Commit()
				if err == nil {
					fmt.Printf("%% Committing offset\n")
				}
			}
			fmt.Printf("%%  Message on %s: %s\n",
				e.TopicPartition, string(e.Value))
		case kafka.PartitionEOF:
			fmt.Printf("%% Reached %v\n", e)
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
		}
	}
}
