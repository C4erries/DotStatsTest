package teststore_test

import (
	"testing"

	"github.com/c4erries/server/internal/app/matchmodel"
	"github.com/c4erries/server/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestMatchRepository_Add(t *testing.T) {
	s := teststore.New()
	m := matchmodel.TestMatch(t)
	assert.NoError(t, s.Match().Add(m))
	assert.NotNil(t, m)
}
