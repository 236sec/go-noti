package otp

type SendOTPPayload struct {
	UserID  string `json:"UserID"`
	Email   string `json:"Email"`
	OTPCode string `json:"OTPCode"`
	RefCode string `json:"RefCode"`
}
