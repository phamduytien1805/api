package pubsub

import (
	"context"
	"log/slog"

	"github.com/phamduytien1805/pkgmodule/config"
	"github.com/twmb/franz-go/pkg/kgo"
)

type Publisher struct {
	config *config.Config
	c      *kgo.Client
	logger *slog.Logger
}

func NewPublisher(config *config.Config, logger *slog.Logger) (*Publisher, error) {
	c, err := kgo.NewClient(kgo.SeedBrokers(config.Kafka.Brokers...))
	logger.Info("Kafka client connected", "brokers", config.Kafka.Brokers)
	if err != nil {
		return nil, err
	}
	return &Publisher{config: config, c: c, logger: logger}, nil
}

func (p *Publisher) PublishMessage(data []byte) error {
	p.logger.Info("Publishing message", "data", string(data))
	err := p.c.ProduceSync(context.Background(), &kgo.Record{
		Topic: "test",
		Value: data,
	})
	return err.FirstErr()
}
