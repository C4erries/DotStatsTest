package sqlstore

import (
	"database/sql"

	"github.com/c4erries/server/internal/app/store"
	_ "github.com/lib/pq"
)

//Основной файл, отвечающий за хранилище

// Структура хранилища (конфигурация, БД, репозиторий пользователей (для доступа к хранилищу через пользователей))
type Store struct {
	db              *sql.DB
	UserRepository  *UserRepository
	MatchRepository *MatchRepository
}

// Открытие (создание) БД
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// Создание UserRepository
func (s *Store) User() store.UserRepository {
	if s.UserRepository != nil {
		return s.UserRepository
	}

	s.UserRepository = &UserRepository{
		store: s,
	}

	return s.UserRepository
}

func (s *Store) Match() store.MatchRepository {
	if s.MatchRepository != nil {
		return s.MatchRepository
	}

	s.MatchRepository = &MatchRepository{
		store: s,
	}

	return s.MatchRepository
}
