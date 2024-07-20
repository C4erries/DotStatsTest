package teststore

import "github.com/c4erries/server/internal/app/statsmodel"

type StatsRepository struct {
	store     *Store
	statistic map[int]*statsmodel.Stats
}

func (r *StatsRepository) NewPlayer(s *statsmodel.Stats) error {
	s.MatchesStats = nil
	s.HeroStats = nil
	s.PlayerStats = nil
	s.WordCloud = nil
	r.statistic[s.PlayerID] = s
	return nil
}

func (r *StatsRepository) UpdateStats(s *statsmodel.Stats) error {
	r.statistic[s.PlayerID] = s
	return nil
}
