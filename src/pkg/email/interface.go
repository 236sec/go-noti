package email

import "context"

type IEmail interface {
	Send(ctx context.Context, to, subject, body string) error
}
