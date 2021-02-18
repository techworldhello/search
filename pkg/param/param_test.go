package param

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseParams(t *testing.T) {
	expectations := []struct {
		name     string
		params   string
		expected Params
	}{
		{
			name: "User params are correctly parsed",
			params: "users=_id:1",
			expected: Params{
				Entity: "users",
				Term:   "_id",
				Value:  "1",
			},
		},
		{
			name: "Ticket params are correctly parsed",
			params: "tickets=status:pending",
			expected: Params{
				Entity: "tickets",
				Term:   "status",
				Value:  "pending",
			},
		},
		{
			name:     "Params not parsed due to format",
			params:   "users:_id=1",
			expected: Params{},
		},
		{
			name:     "Params not parsed due to invalid entity",
			params:   "ticket=_id:1",
			expected: Params{},
		},
	}

	for _, expect := range expectations {
		t.Run(expect.name, func(t *testing.T) {
			assert.Equal(t, expect.expected, Parse(expect.params))
		})
	}
}
