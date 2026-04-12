package kafka

import (
	"context"
	"log"
)

type NoOpProducer struct{}

func NewNoOpProducer() IProducer {
	return &NoOpProducer{}
}

func (p *NoOpProducer) Publish(ctx context.Context, topic string, key []byte, value []byte) error {
	log.Printf("[NoOpBroker] Ignored publishing to topic: %s", topic)
	return nil
}

func (p *NoOpProducer) Ping(ctx context.Context) error {
	return nil
}

func (p *NoOpProducer) Close() error {
	return nil
}

type NoOpConsumer struct{}

func NewNoOpConsumer() IConsumer {
	return &NoOpConsumer{}
}

func (c *NoOpConsumer) Consume(ctx context.Context, topic string, handler func(ctx context.Context, message []byte) error) error {
	log.Printf("[NoOpBroker] Ignored consumption setup for topic: %s", topic)
	return nil
}

func (c *NoOpConsumer) Ping(ctx context.Context) error {
	return nil
}

func (c *NoOpConsumer) Close() error {
	return nil
}
