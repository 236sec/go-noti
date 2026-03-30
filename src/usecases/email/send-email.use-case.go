package email

import (
	"context"

	"goboilerplate.com/src/pkg/email"
	"goboilerplate.com/src/usecases"
)

type ISendEmailUseCase interface {
	Apply(ctx context.Context, req SendEmailRequest) error
}

type SendEmailUseCase struct {
	emailService email.IEmail
}

func NewSendEmailUseCase(emailService email.IEmail) *SendEmailUseCase {
	return &SendEmailUseCase{emailService: emailService}
}

func (u *SendEmailUseCase) Apply(ctx context.Context, req SendEmailRequest) error {
	if err := u.emailService.Send(ctx, req.To, req.Subject, req.Body); err != nil {
		return usecases.ErrCannotSendEmail
	}
	return nil
}
