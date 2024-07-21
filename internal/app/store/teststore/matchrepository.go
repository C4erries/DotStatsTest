package teststore

import "github.com/c4erries/server/internal/app/matchmodel"

type MatchRepository struct {
	store   *Store
	matches map[int]*matchmodel.Match
}

func (r *MatchRepository) Add(m *matchmodel.Match) error {
	if err := m.Validate(); err != nil {
		return err
	}

	r.matches[len(r.matches)] = m
	return nil
}
