package consumers

import (
	"context"
	"encoding/json"
	"log/slog"

	"goboilerplate.com/src/pkg/kafka"
	"goboilerplate.com/src/usecases/otp"
)

type OTPConsumer struct {
	consumer            kafka.IConsumer
	sendOTPEmailUseCase otp.ISendOTPEmailUseCase
}

func NewOTPConsumer(consumer kafka.IConsumer, sendOTPEmailUseCase otp.ISendOTPEmailUseCase) *OTPConsumer {
	return &OTPConsumer{
		consumer:            consumer,
		sendOTPEmailUseCase: sendOTPEmailUseCase,
	}
}

func (oc *OTPConsumer) Consume(ctx context.Context) {
	slog.Info("Starting OTP consumer on topic OTP_REQUESTED")

	err := oc.consumer.Consume(ctx, "OTP_REQUESTED", func(ctx context.Context, message []byte) error {
		var payload otp.SendOTPPayload
		if err := json.Unmarshal(message, &payload); err != nil {
			slog.ErrorContext(ctx, "Failed to unmarshal OTP_REQUESTED message", "error", err, "message", string(message))
			return nil
		}

		err := oc.sendOTPEmailUseCase.Apply(ctx, payload)
		if err != nil {
			slog.ErrorContext(ctx, "Failed to process OTP email send", "error", err)
			return err
		}

		slog.InfoContext(ctx, "Successfully processed OTP_REQUESTED message", "UserID", payload.UserID)
		return nil
	})

	if err != nil {
		slog.Error("OTP consumer stopped with error", "error", err)
	}
}
