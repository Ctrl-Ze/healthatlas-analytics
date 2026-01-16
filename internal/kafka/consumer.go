package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	Reader *kafka.Reader
}

func NewConsumer(brokers []string, topic, groupID string) *KafkaConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic: topic,
		GroupID: groupID,
		MinBytes: 10e3,
		MaxBytes: 10e6,
		StartOffset: kafka.FirstOffset,
	})

	return &KafkaConsumer{Reader: reader}
}

func (c *KafkaConsumer) Start(ctx context.Context) {
	go func() {
		log.Println("Kafka consumer started...")
		for {
			m, err := c.Reader.ReadMessage(ctx)
			if err != nil {
				log.Println("Error reading Kafka message:", err)
				continue
			}

			log.Printf("Message received: %s\n", string(m.Value))
		}
	}()
}