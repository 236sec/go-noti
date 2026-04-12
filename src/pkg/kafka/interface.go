package kafka

import "context"

type IProducer interface {
	Publish(ctx context.Context, topic string, key []byte, value []byte) error
	Ping(ctx context.Context) error
	Close() error
}

type IConsumer interface {
	Consume(ctx context.Context, topic string, handler func(ctx context.Context, message []byte) error) error
	Ping(ctx context.Context) error
	Close() error
}
