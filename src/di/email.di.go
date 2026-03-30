package di

import (
	"log"
	"sync"

	"goboilerplate.com/config"
	"goboilerplate.com/src/pkg/email"
)

var getEmailService = sync.OnceValue(func() email.IEmail {
	cfg := config.GetConfig()
	emailService, err := email.NewEmail(&cfg.EnvConfig.Email)
	if err != nil {
		log.Fatalf("failed to create email service: %v", err)
	}
	return emailService
})
