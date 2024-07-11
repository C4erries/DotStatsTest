package teststore

import (
	"github.com/c4erries/server/internal/app/model"
	"github.com/c4erries/server/internal/app/store"
)

// (ТЕСТ) Репозиторий пользователей
type UserRepository struct {
	store *Store
	users map[int]*model.User
}

// (ТЕСТ) Создание пользователей
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	u.ID = len(r.users)
	r.users[u.ID] = u

	return nil
}

// (ТЕСТ) Поиск пользователя по почте
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}

	return nil, store.ErrRecordNotFound
}

func (r *UserRepository) Find(id int) (*model.User, error) {
	u, ok := r.users[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}

func (r *UserRepository) ListAll() ([]*model.User, error) {
	var Us []*model.User
	for _, u := range r.users {
		Us = append(Us, u)
	}
	return Us, nil
}
