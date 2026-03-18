package domain

import "goboilerplate.com/src/models"

type User struct {
	id          int
	firstName   string
	lastName    string
	username    string
	password    string
	role        string
	dateOfBirth string
}

// --- Getters ---
func (u User) ID() int            { return u.id }
func (u User) FirstName() string  { return u.firstName }
func (u User) LastName() string   { return u.lastName }
func (u User) Username() string   { return u.username }
func (u User) Password() string   { return u.password }
func (u User) Role() string       { return u.role }
func (u User) DateOfBirth() string { return u.dateOfBirth }

func (u *User) GetFullName() string {
	return u.firstName + " " + u.lastName
}

func (u *User) IsAdmin() bool {
	return u.role == "admin"
}

func (u *User) IsAbleToLogin() bool {
	return u.role == "admin" || u.role == "user"
}

// FromModel converts models.User to domain.User
func FromModel(modelUser models.User) User {
	return User{
		id:          modelUser.ID,
		firstName:   modelUser.FirstName,
		lastName:    modelUser.LastName,
		username:    modelUser.Username,
		password:    modelUser.Password,
		role:        modelUser.Role,
		dateOfBirth: modelUser.DateOfBirth,
	}
}

// ToModel converts domain.User to models.User
func (u *User) ToModel() models.User {
	return models.User{
		ID:          u.id,
		FirstName:   u.firstName,
		LastName:    u.lastName,
		Username:    u.username,
		Password:    u.password,
		Role:        u.role,
		DateOfBirth: u.dateOfBirth,
	}
}