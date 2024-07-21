package model_test

import (
	"testing"

	model "github.com/c4erries/server/internal/app/model"
	"github.com/stretchr/testify/assert"
)

// Конфигурация тестов валидации пользователей
func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.User {
				u := model.TestUser(t)
				return u
			},
			isValid: true,
		},

		{
			name: "With enc. pass",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				u.EncryptedPassword = "Encrypted"
				return u
			},
			isValid: true,
		},

		{
			name: "Empty Email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = ""
				return u
			},
			isValid: false,
		},

		{
			name: "Invalid Email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = "pen is black"
				return u
			},
			isValid: false,
		},

		{
			name: "Empty pass",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				return u
			},
			isValid: false,
		},

		{
			name: "Short pass",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = "short"
				return u
			},
			isValid: false,
		},

		{
			name: "No nick",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Nickname = ""
				return u
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}
