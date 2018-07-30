package common

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
	"strings"
)

//NewConsumer returns a new kafka Consumer
func NewConsumer() *kafka.Consumer {

	broker := GetEnvironmentVar("KAFKA_BROKERS", "localhost:9092")
	consumerGroup := fmt.Sprintf("%v_%v", GetEnvironmentVar("ENVIRONMENT", "local"), GetEnvironmentVar("KAFKA_CONSUMER_GROUP", "es_index"))
	topics := strings.Split(GetEnvironmentVar("KAFKA_TOPICS", "turnpike-events-dev"), ",")

	c := createKafkaConsumer(broker, consumerGroup)
	err := c.SubscribeTopics(topics, nil)
	if err != nil {
		panic(err)
	}

	return c
}

func createKafkaConsumer(broker, consumerGroup string) *kafka.Consumer {
	conf := &kafka.ConfigMap{
		"bootstrap.servers":               broker,
		"group.id":                        consumerGroup,
		"session.timeout.ms":              6000,
		"go.events.channel.enable":        true,
		"go.application.rebalance.enable": true,
		"default.topic.config":            kafka.ConfigMap{"auto.offset.reset": "earliest"},
	}

	consumer, err := kafka.NewConsumer(conf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	return consumer
}
