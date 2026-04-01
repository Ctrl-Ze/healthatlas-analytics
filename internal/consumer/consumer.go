package consumer

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Ctrl-Ze/healthatlas-analytics/internal/db"
	"github.com/segmentio/kafka-go"
)

type store interface {
	InsertMetric(ctx context.Context, m db.Metric) error
	UpsertTrend(ctx context.Context, t db.Trend) error
}

type Consumer struct {
	reader *kafka.Reader
	store  store
}

func New(brokers []string, topic, groupID string, s store) *Consumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     brokers,
		Topic:       topic,
		GroupID:     groupID,
		MinBytes:    10e3,
		MaxBytes:    10e6,
		StartOffset: kafka.FirstOffset,
	})

	return &Consumer{reader: reader, store: s}
}

func (c *Consumer) Start(ctx context.Context) {
	go func() {
		defer c.reader.Close()
		log.Println("Kafka consumer started...")
		for {
			m, err := c.reader.ReadMessage(ctx)
			if err != nil {
				if ctx.Err() != nil {
					return
				}
				log.Println("Error reading Kafka message:", err)
				continue
			}

			var event Event
			if err := json.Unmarshal(m.Value, &event); err != nil {
				log.Println("Failed to parse event:", err)
				continue
			}

			for _, r := range event.Results {
				if err := c.store.InsertMetric(ctx, db.Metric{
					UserID:      event.UserID,
					AnalyteCode: r.AnalyteCode,
					Value:       r.Value,
					Time:        event.Timestamp,
				}); err != nil {
					log.Println("InsertMetrics error:", err)
				}
				if err := c.store.UpsertTrend(ctx, db.Trend{
					UserID:      event.UserID,
					AnalyteCode: r.AnalyteCode,
					Value:       r.Value,
				}); err != nil {
					log.Println("UpsertTrend error:", err)
				}
			}

			log.Printf("Stored blood test %s: %d analytes\n", event.BloodTestID, len(event.Results))
		}
	}()
}
