package consumers

import (
	"context"

	"goboilerplate.com/src/di"
)

func RegisterConsumers(ctx context.Context) {
	kafkaConsumer := di.GetKafkaConsumer()
	otpConsumer := NewOTPConsumer(kafkaConsumer, di.GetSendOTPEmailUseCase())

	go otpConsumer.Consume(ctx)
}
