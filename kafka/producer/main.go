package main

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

func main() {
	// Kafka broker addresses
	brokers := []string{"localhost:9092"}

	// Create a new Sarama configuration
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	// Create a new synchronous producer
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("Failed to start Sarama producer: %v", err)
	}
	defer producer.Close()

	// Message to be sent
	msg := &sarama.ProducerMessage{
		Topic: "test-topic-2",
		Value: sarama.StringEncoder("Hello Kafka!"),
	}

	// Send the message
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)
}
