package email

import (
	"context"

	"github.com/wneessen/go-mail"
	"goboilerplate.com/config"
)

type Email struct {
	client *mail.Client
	from   string
}

func NewEmail(cfg *config.EmailConfig) (IEmail, error) {
	c, err := mail.NewClient(
		cfg.EmailHost,
		mail.WithTLSPortPolicy(mail.TLSMandatory),
		mail.WithSMTPAuth(mail.SMTPAuthAutoDiscover),
		mail.WithUsername(cfg.EmailUsername),
		mail.WithPassword(cfg.EmailPassword),
	)
	if err != nil {
		return nil, err
	}
	return &Email{
		client: c,
		from:   cfg.EmailFrom,
	}, nil
}

func (e *Email) Send(ctx context.Context, to, subject, body string) error {
	m := mail.NewMsg()
	if err := m.From(e.from); err != nil {
		return err
	}
	if err := m.To(to); err != nil {
		return err
	}
	m.Subject(subject)
	m.SetBodyString(mail.TypeTextPlain, body)
	if err := e.client.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
