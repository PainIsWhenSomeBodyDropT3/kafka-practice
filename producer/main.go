package main

import (
	"fmt"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	topics, err := configureAdminClient()
	if err != nil {
		fmt.Println("configure admin error", err)
		os.Exit(1)
	}

	fmt.Println("configure producer")
	producer, err := configureProducer()
	if err != nil {
		fmt.Println("configure producer error", err)
		os.Exit(1)
	}

	fmt.Println("logic")
	done := logic(producer, topics[0])
	<-done
	fmt.Println("done")
}

func logic(p *kafka.Producer, topic string) <-chan struct{} {
	done := make(chan struct{})
	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		timeout := time.After(600 * time.Second)
		defer ticker.Stop()
		defer close(done)
		for {
			select {
			case <-timeout:
				done <- struct{}{}
				return
			case <-ticker.C:
				fmt.Println("producing")
				err := p.Produce(&kafka.Message{
					TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
					Value:          []byte("FOO"),
				}, nil)
				if err != nil {
					fmt.Println("produce error", err)
					done <- struct{}{}
					return
				}
			}
		}
	}()

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Successfully produced record to topic %s partition [%d] @ offset %v\n",
						*ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset)
				}
			}
		}
	}()

	return done
}
