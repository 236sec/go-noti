package user

type LoginUserRequest struct {
	Username string `binding:"required" json:"username"`
	Password string `binding:"required" json:"password"`
}

type LoginUserResponse struct {
	Token string `json:"token"`
}
