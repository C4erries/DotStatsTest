package matchmodel

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// Структура для операций с матчами из БД
type Match struct {
	MatchID int      `json:"matchid"`
	Team_R  [5]int64 `json:"team_r"`
	Team_D  [5]int64 `json:"team_d"`
	TimeLen string   `json:"timelength"`
	Result  string   `json:"result"`
}

// Валидация данных о матче
func (m *Match) Validate() error {
	return validation.ValidateStruct(m, validation.Field(&m.MatchID, validation.Required),
		validation.Field(&m.TimeLen, validation.Required), validation.Field(&m.Result, validation.Required))
}
