package teststore

import (
	"github.com/c4erries/server/internal/app/model"
	"github.com/c4erries/server/internal/app/store"
)

//Служебное тестовое хранилище

// (ТЕСТ) Структура хранилища (конфигурация, БД, репозиторий пользователей (для доступа к хранилищу через пользователей))
type Store struct {
	UserRepository *UserRepository
}

// Открытие (создание) БД
func New() *Store {
	return &Store{}
}

// Проверка на существование/добавление пользователя
func (s *Store) User() store.UserRepository {
	if s.UserRepository != nil {
		return s.UserRepository
	}

	s.UserRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}

	return s.UserRepository
}
