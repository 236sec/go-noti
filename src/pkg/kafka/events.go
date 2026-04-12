package kafka

type UserCreatedEvent struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	Role   string `json:"role"`
}

type UserLoginEvent struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
}
