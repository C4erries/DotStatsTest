package sqlstore_test

import (
	"testing"

	"github.com/c4erries/server/internal/app/model"
	"github.com/c4erries/server/internal/app/statsmodel"
	"github.com/c4erries/server/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestStatsRepository_PushStats(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("stats")
	defer teardown("users")

	s := sqlstore.New(db)
	tes := statsmodel.TestStats(t)
	p := model.TestUser(t)
	assert.NoError(t, s.User().Create(p))
	assert.NoError(t, s.Stats().NewPlayer(tes))
	assert.NoError(t, s.Stats().UpdateStats(tes))
	assert.NotNil(t, tes.PlayerID)
}
