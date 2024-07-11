package sqlstore

import (
	"database/sql"

	"github.com/c4erries/server/internal/app/model"
	"github.com/c4erries/server/internal/app/store"
)

// Репозиторий пользователей
type UserRepository struct {
	store *Store
}

// Создание пользователя (ТОЛЬКО ЕСЛИ НЕ СУЩЕСТВУЕТ)
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO users (nickname, email, encrypted_password, player_id) VALUES ($1, $2, $3, $4) RETURNING id",
		u.Nickname, u.Email, u.EncryptedPassword, u.PlayerID,
	).Scan(&u.ID)

}

// Поиск пользователя по адресу почты
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {

	u := &model.User{}

	if err := r.store.db.QueryRow(
		"SELECT id, nickname, email, encrypted_password, player_id FROM users WHERE email=$1",
		email).Scan(&u.ID, &u.Nickname, &u.Email, &u.EncryptedPassword, &u.PlayerID); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}
func (r *UserRepository) FindByNickname(nickname string) (*model.User, error) {

	u := &model.User{}

	if err := r.store.db.QueryRow(
		"SELECT id, nickname, email, encrypted_password, player_id FROM users WHERE nickname=$1",
		nickname).Scan(&u.ID, &u.Nickname, &u.Email, &u.EncryptedPassword, &u.PlayerID); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}

// Поиск пользователя по id
func (r *UserRepository) Find(id int) (*model.User, error) {

	u := &model.User{}

	if err := r.store.db.QueryRow(
		"SELECT id, nickname, email, encrypted_password, player_id FROM users WHERE id=$1",
		id).Scan(&u.ID, &u.Nickname, &u.Email, &u.EncryptedPassword, &u.PlayerID); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}

// Выдать всех пользователей
func (r *UserRepository) ListAll() ([]*model.User, error) {

	var Us []*model.User

	rows, err := r.store.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		u := &model.User{}
		if err := rows.Scan(&u.ID, &u.Nickname, &u.Email, &u.EncryptedPassword, &u.PlayerID); err != nil {
			return nil, err
		}
		Us = append(Us, u)
	}

	return Us, nil
}
