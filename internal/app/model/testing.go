package model

import "testing"

//Тестовый пользователь для удобства тестировки
func TestUser(t *testing.T) *User {
	return &User{
		Nickname: "Usar",
		Email:    "user@ple.org",
		Password: "password",
		PlayerID: 1442,
	}
}
