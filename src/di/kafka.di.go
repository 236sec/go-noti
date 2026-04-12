package di

import (
	"strings"
	"sync"

	"goboilerplate.com/config"
	"goboilerplate.com/src/pkg/kafka"
)

var GetKafkaProducer = sync.OnceValue(func() kafka.IProducer {
	cfg := config.GetConfig()

	if !cfg.YMLConfig.Broker.Enabled {
		return kafka.NewNoOpProducer()
	}

	urls := strings.Split(cfg.EnvConfig.Kafka.URL, ",")
	return kafka.NewSegmentioKafkaProducer(urls).
		WithAllowedAutoCreateTopic(cfg.YMLConfig.Broker.AllowAutoTopicCreation).
		WithMaxAttempts(cfg.YMLConfig.Broker.MaxAttempts)
})

var GetKafkaConsumer = sync.OnceValue(func() kafka.IConsumer {
	cfg := config.GetConfig()

	if !cfg.YMLConfig.Broker.Enabled {
		return kafka.NewNoOpConsumer()
	}

	urls := strings.Split(cfg.EnvConfig.Kafka.URL, ",")
	return kafka.NewSegmentioKafkaConsumer(urls, "go-noti-group")
})
