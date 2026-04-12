package kafka

import (
	"context"
	"errors"
	"log"

	"github.com/segmentio/kafka-go"
)

type SegmentioKafkaProducer struct {
	writer     *kafka.Writer
	brokerURLs []string
}

func NewSegmentioKafkaProducer(brokerURLs []string) *SegmentioKafkaProducer {
	w := &kafka.Writer{
		Addr:     kafka.TCP(brokerURLs...),
		Balancer: &kafka.LeastBytes{},
	}
	return &SegmentioKafkaProducer{
		writer:     w,
		brokerURLs: brokerURLs,
	}
}

func (p *SegmentioKafkaProducer) WithMaxAttempts(attempts int) *SegmentioKafkaProducer {
	p.writer.MaxAttempts = attempts
	return p
}

func (p *SegmentioKafkaProducer) WithAllowedAutoCreateTopic(allow bool) *SegmentioKafkaProducer {
	p.writer.AllowAutoTopicCreation = allow
	return p
}

func (p *SegmentioKafkaProducer) Publish(ctx context.Context, topic string, key []byte, value []byte) error {
	msg := kafka.Message{
		Topic: topic,
		Key:   key,
		Value: value,
	}

	err := p.writer.WriteMessages(ctx, msg)
	if err != nil {
		log.Printf("Failed to write to kafka topic %s: %v", topic, err)
		return err
	}
	return nil
}

func (p *SegmentioKafkaProducer) Ping(ctx context.Context) error {
	if len(p.brokerURLs) == 0 {
		return errors.New("kafka broker URLs missing")
	}
	conn, err := kafka.DialContext(ctx, "tcp", p.brokerURLs[0])
	if err != nil {
		return err
	}
	return conn.Close()
}

func (p *SegmentioKafkaProducer) Close() error {
	if p.writer != nil {
		return p.writer.Close()
	}
	return nil
}

type SegmentioKafkaConsumer struct {
	brokerURLs []string
	groupID    string
	readers    []*kafka.Reader
}

func NewSegmentioKafkaConsumer(brokerURLs []string, groupID string) *SegmentioKafkaConsumer {
	return &SegmentioKafkaConsumer{
		brokerURLs: brokerURLs,
		groupID:    groupID,
		readers:    make([]*kafka.Reader, 0),
	}
}

func (c *SegmentioKafkaConsumer) Consume(ctx context.Context, topic string, handler func(ctx context.Context, message []byte) error) error {
	if len(c.brokerURLs) == 0 {
		return errors.New("kafka broker URLs missing")
	}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: c.brokerURLs,
		GroupID: c.groupID,
		Topic:   topic,
	})

	c.readers = append(c.readers, reader)

	go func() {
		defer func() {
			if err := reader.Close(); err != nil {
				log.Printf("Failed to close reader for topic %s: %v", topic, err)
			}
		}()
		for {
			m, err := reader.FetchMessage(ctx)
			if err != nil {
				if errors.Is(err, context.Canceled) {
					return
				}
				log.Printf("Failed to fetch message from topic %s: %v", topic, err)
				continue
			}

			if err := handler(ctx, m.Value); err != nil {
				log.Printf("Failed to handle message from topic %s: %v", topic, err)
				// Retry or dead letter queue logic could be added here
			} else {
				if err := reader.CommitMessages(ctx, m); err != nil {
					log.Printf("Failed to commit message: %v", err)
				}
			}
		}
	}()

	return nil
}

func (c *SegmentioKafkaConsumer) Ping(ctx context.Context) error {
	if len(c.brokerURLs) == 0 {
		return errors.New("kafka broker URLs missing")
	}
	conn, err := kafka.DialContext(ctx, "tcp", c.brokerURLs[0])
	if err != nil {
		return err
	}
	return conn.Close()
}

func (c *SegmentioKafkaConsumer) Close() error {
	var errs []error
	for _, reader := range c.readers {
		if err := reader.Close(); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return errors.New("failed to close one or more readers")
	}
	return nil
}
