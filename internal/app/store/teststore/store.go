package teststore

import (
	"github.com/c4erries/server/internal/app/matchmodel"
	"github.com/c4erries/server/internal/app/model"
	"github.com/c4erries/server/internal/app/statsmodel"
	"github.com/c4erries/server/internal/app/store"
)

//Служебное тестовое хранилище

// (ТЕСТ) Структура хранилища (конфигурация, БД, репозиторий пользователей (для доступа к хранилищу через пользователей))
type Store struct {
	UserRepository  *UserRepository
	MatchRepository *MatchRepository
	StatsRepository *StatsRepository
}

// Открытие (создание) БД
func New() *Store {
	return &Store{}
}

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

func (s *Store) Match() store.MatchRepository {
	if s.MatchRepository != nil {
		return s.MatchRepository
	}

	s.MatchRepository = &MatchRepository{
		store:   s,
		matches: make(map[int]*matchmodel.Match),
	}

	return s.MatchRepository
}

func (s *Store) Stats() store.StatsRepository {
	if s.StatsRepository != nil {
		return s.StatsRepository
	}

	s.StatsRepository = &StatsRepository{
		store:     s,
		statistic: make(map[int]*statsmodel.Stats),
	}

	return s.StatsRepository
}
