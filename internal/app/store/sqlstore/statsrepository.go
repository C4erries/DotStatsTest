package sqlstore

import (
	"github.com/c4erries/server/internal/app/statsmodel"
)

type StatsRepository struct {
	store *Store
}

func (r *StatsRepository) NewPlayer(s *statsmodel.Stats) error {
	return r.store.db.QueryRow(
		"INSERT INTO stats (player_id) VALUES ($1) RETURNING player_id",
		s.PlayerID,
	).Scan(&s.PlayerID)
}

func (r *StatsRepository) UpdateStats(s *statsmodel.Stats) error {
	return r.store.db.QueryRow(
		"UPDATE stats SET player_stats = $1,  heroes_stats=$2, matches_stats=$3, wordcloud_stats=$4 where player_id = $5 RETURNING player_id",
		s.PlayerStats, s.HeroStats, s.MatchesStats, s.WordCloud, s.PlayerID,
	).Scan(&s.PlayerID)
}
