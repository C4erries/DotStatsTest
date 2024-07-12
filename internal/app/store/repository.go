package store

import (
	"github.com/c4erries/server/internal/app/matchmodel"
	"github.com/c4erries/server/internal/app/model"
)

// Интерфейс репозитория пользователей
type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
	FindByNickname(string) (*model.User, error)
	ListAll() ([]*model.User, error)
}

type MatchRepository interface {
	Add(*matchmodel.Match) error
}
