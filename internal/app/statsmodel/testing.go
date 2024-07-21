package statsmodel

import "testing"

//Тестовый пользователь для удобства тестировки
func TestStats(t *testing.T) *Stats {
	return &Stats{
		PlayerID:     1442,
		PlayerStats:  PropertyMap{"tom": 1, "bob": 2},
		HeroStats:    PropertyMap{"ter": 1},
		MatchesStats: PropertyMap{"ter": 13, "nig": 14},
		WordCloud:    PropertyMap{"sex": 436},
	}
}
