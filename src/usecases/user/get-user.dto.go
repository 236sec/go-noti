package user

type GetUserResponse struct {
	ID       int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username string `json:"username"`
	DateOfBirth string `json:"date_of_birth"`
}