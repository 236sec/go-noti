package email

type SendEmailRequest struct {
	To      string `binding:"required,email" json:"to"`
	Subject string `binding:"required"       json:"subject"`
	Body    string `binding:"required"       json:"body"`
}
