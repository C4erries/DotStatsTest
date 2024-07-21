package teststore_test

import (
	"testing"

	"github.com/c4erries/server/internal/app/statsmodel"
	"github.com/c4erries/server/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestStatsRepository_UpdateStats(t *testing.T) {
	s := teststore.New()
	stat := statsmodel.TestStats(t)
	assert.NoError(t, s.Stats().NewPlayer(stat))
	assert.NoError(t, s.Stats().UpdateStats(stat))
	assert.NotNil(t, stat)
}
