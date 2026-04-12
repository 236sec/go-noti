package otp

import (
	"context"
	"fmt"
	"log/slog"

	"goboilerplate.com/src/pkg/email"
	"goboilerplate.com/src/usecases"
)

type ISendOTPEmailUseCase interface {
	Apply(ctx context.Context, req SendOTPPayload) error
}

type SendOTPEmailUseCase struct {
	emailService email.IEmail
}

func NewSendOTPEmailUseCase(emailService email.IEmail) *SendOTPEmailUseCase {
	return &SendOTPEmailUseCase{
		emailService: emailService,
	}
}

func (u *SendOTPEmailUseCase) Apply(ctx context.Context, req SendOTPPayload) error {
	slog.InfoContext(ctx, "Sending OTP email", "email", req.Email, "ref", req.RefCode)

	subject := fmt.Sprintf("Your OTP Code (Ref: %s)", req.RefCode)
	body := fmt.Sprintf("Your OTP is %s. Do not share it with anyone.", req.OTPCode)

	err := u.emailService.Send(ctx, req.Email, subject, body)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to send OTP email", "error", err)
		return usecases.ErrCannotSendEmail
	}

	return nil
}
