package user

type CreateUserRequest struct {
	FirstName   string `binding:"required" json:"firstName"`
	LastName    string `binding:"required" json:"lastName"`
	Username    string `binding:"required" json:"username"`
	Password    string `binding:"required" json:"password"`
	DateOfBirth string `binding:"required" json:"dateOfBirth"`
}

type CreateUserResponse struct {
	ID int `json:"id"`
}
