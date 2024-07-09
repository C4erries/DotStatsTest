package matchmodel_test

import (
	"testing"

	"github.com/c4erries/server/internal/app/matchmodel"
	"github.com/stretchr/testify/assert"
)

func TestMatch_Add(t *testing.T) {
	testCases := []struct {
		name    string
		m       func() *matchmodel.Match
		isValid bool
	}{
		{
			name: "valid",
			m: func() *matchmodel.Match {
				m := matchmodel.TestMatch(t)
				return m
			},
			isValid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.m().Validate())
			} else {
				assert.Error(t, tc.m().Validate())
			}
		})
	}
}
