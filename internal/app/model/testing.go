package model

import "testing"

//Тестовый пользователь для удобства тестировки
func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user@example.org",
		Password: "password",
	}
}
