package sqlstore

import (
	"database/sql"

	"github.com/c4erries/server/internal/app/store"
	_ "github.com/lib/pq"
)

//Основной файл, отвечающий за хранилище

// Структура хранилища (конфигурация, БД, репозиторий пользователей (для доступа к хранилищу через пользователей))
type Store struct {
	db             *sql.DB
	UserRepository *UserRepository
}

// Открытие (создание) БД
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// Проверка на существование/добавление пользователя
func (s *Store) User() store.UserRepository {
	if s.UserRepository != nil {
		return s.UserRepository
	}

	s.UserRepository = &UserRepository{
		store: s,
	}

	return s.UserRepository
}
