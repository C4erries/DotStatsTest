package sqlstore

import (
	"github.com/c4erries/server/internal/app/matchmodel"
	"github.com/lib/pq"
)

type MatchRepository struct {
	store *Store
}

func (r *MatchRepository) Add(m *matchmodel.Match) error {
	if err := m.Validate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO matches (match_id, team_R, team_D, time_length, result) VALUES ($1, $2, $3, $4, $5) RETURNING match_id",
		m.MatchID, pq.Array(m.Team_R), pq.Array(m.Team_D), m.TimeLen, m.Result,
	).Scan(&m.MatchID)

}
