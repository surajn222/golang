package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/IBM/sarama"
)

func main() {
	// Kafka broker addresses
	brokers := []string{"localhost:9092"}
	topic := "test-topic-2"
	group := "your-consumer-group"

	// Create a new Sarama configuration
	config := sarama.NewConfig()
	config.Version = sarama.V2_6_0_0 // Specify the Kafka version you're using

	// Enable auto-commit of offsets
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second // Interval for auto-committing offsets

	// Optional: Configure the rebalance strategy
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin

	// Create a new consumer group
	consumerGroup, err := sarama.NewConsumerGroup(brokers, group, config)
	if err != nil {
		log.Fatalf("Failed to create consumer group: %v", err)
	}
	defer consumerGroup.Close()

	// Create a context to manage shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle signals for graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signals
		cancel()
	}()

	// Create a consumer group handler
	handler := ConsumerGroupHandler{}

	// Consume messages in a loop
	for {
		err := consumerGroup.Consume(ctx, []string{topic}, handler)
		if err != nil {
			log.Fatalf("Error during consumption: %v", err)
		}

		// Check if context was canceled
		if ctx.Err() != nil {
			break
		}
	}
}

// ConsumerGroupHandler implements sarama.ConsumerGroupHandler
type ConsumerGroupHandler struct{}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (h ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (h ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim starts a consumer loop of ConsumerGroupClaim's Messages().
func (h ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		// Process the message
		fmt.Printf("Partition: %d, Offset: %d, Message: %s\n", message.Partition, message.Offset, string(message.Value))
	}
	return nil
}
