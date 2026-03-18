package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"goboilerplate.com/src/models"
)

func TestUser_GetFullName(t *testing.T) {
	testCases := []struct {
		name     string
		firstName string
		lastName  string
		expected  string
	}{
		{"Should return full name when both first and last names are provided", "John", "Doe", "John Doe"},
		{"Should return first name when last name is empty", "John", "", "John "},
		{"Should return last name when first name is empty", "", "Doe", " Doe"},
		{"Should return empty string when both first and last names are empty", "", "", " "},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			user := User{
				firstName: tt.firstName,
				lastName:  tt.lastName,
			}
			assert.Equal(t, tt.expected, user.GetFullName())
		})
	}
}

func TestUser_IsAdmin(t *testing.T) {
	testCases := []struct {
		name     string
		role     string
		expected bool
	}{
		{"Should be admin when role is admin", "admin", true},
		{"Should not be admin when role is user", "user", false},
		{"Should not be admin when role is empty", "", false},
		{"Should not be admin when role is guest", "guest", false},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			user := User{role: tt.role}
			assert.Equal(t, tt.expected, user.IsAdmin())
		})
	}
}

func TestUser_IsAbleToLogin(t *testing.T) {
	testCases := []struct {
		name     string
		role     string
		expected bool
	}{
		{"Admin role Should Be Able to Login", "admin", true},
		{"User role Should Be Able to Login", "user", true},
		{"Guest role Should Not Be Able to Login", "guest", false},
		{"Empty role Should Not Be Able to Login", "", false},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			user := User{role: tt.role}
			assert.Equal(t, tt.expected, user.IsAbleToLogin())
		})
	}
}

func TestFromModel(t *testing.T) {
	model := models.User{
		ID:          1,
		FirstName:   "Jane",
		LastName:    "Doe",
		Username:    "janedoe",
		Password:    "secret",
		Role:        "admin",
		DateOfBirth: "1990-01-01",
	}

	domainUser := FromModel(model)

	assert.Equal(t, model.ID, domainUser.id)
	assert.Equal(t, model.FirstName, domainUser.firstName)
	assert.Equal(t, model.LastName, domainUser.lastName)
	assert.Equal(t, model.Username, domainUser.username)
	assert.Equal(t, model.Password, domainUser.password)
	assert.Equal(t, model.Role, domainUser.role)
	assert.Equal(t, model.DateOfBirth, domainUser.dateOfBirth)
}

func TestUser_ToModel(t *testing.T) {
	user := User{
		id:          1,
		firstName:   "Jane",
		lastName:    "Doe",
		username:    "janedoe",
		password:    "secret",
		role:        "admin",
		dateOfBirth: "1990-01-01",
	}

	model := user.ToModel()

	assert.Equal(t, user.id, model.ID)
	assert.Equal(t, user.firstName, model.FirstName)
	assert.Equal(t, user.lastName, model.LastName)
	assert.Equal(t, user.username, model.Username)
	assert.Equal(t, user.password, model.Password)
	assert.Equal(t, user.role, model.Role)
	assert.Equal(t, user.dateOfBirth, model.DateOfBirth)
}
