package sqlstore_test

import (
	"testing"

	"github.com/c4erries/server/internal/app/matchmodel"
	"github.com/c4erries/server/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestMatchRepository_Add(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("matches")

	s := sqlstore.New(db)
	m := matchmodel.TestMatch(t)
	assert.NoError(t, s.Match().Add(m))
	assert.NotNil(t, m.MatchID)
}
