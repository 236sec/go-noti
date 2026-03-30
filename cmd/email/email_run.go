package main

import (
	"context"
	"log"

	"goboilerplate.com/src/di"
	"goboilerplate.com/src/usecases/email"
)

func main() {
	sendEmailUseCase := di.GetSendEmailUseCase()

	req := email.SendEmailRequest{
		To:      "Email Here @ hostname.com",
		Subject: "Test Email",
		Body:    "Hello, this is a test email!",
	}

	err := sendEmailUseCase.Apply(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to send email: %v", err)
	}

	log.Println("Email sent successfully!")
}
