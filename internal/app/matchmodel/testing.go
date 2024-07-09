package matchmodel

import "testing"

//Тестовый пользователь для удобства тестировки
func TestMatch(t *testing.T) *Match {
	return &Match{
		MatchID: 1234567,
		Team_R:  [5]int64{12, 24, 36, 48, 60},
		Team_D:  [5]int64{22, 33, 77, 88, 99},
		TimeLen: "23:41",
		Result:  "Dire",
	}
}
