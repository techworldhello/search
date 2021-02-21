package handler

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
			params: "table-users=_id:1",
			expected: Params{
				format: "table",
				entity: "users",
				term:   "_id",
				value:  "1",
			},
		},
		{
			name: "Ticket params are correctly parsed",
			params: "json-tickets=status:pending",
			expected: Params{
				format: "json",
				entity: "tickets",
				term:   "status",
				value:  "pending",
			},
		},
		{
			name:     "Params not parsed due to format",
			params:   "users:_id=1",
			expected: Params{},
		},
		{
			name:     "Params not parsed due to invalid entity",
			params:   "another_format-ticket=_id:1",
			expected: Params{},
		},
	}

	for _, expect := range expectations {
		t.Run(expect.name, func(t *testing.T) {
			assert.Equal(t, expect.expected, ParseParams(expect.params))
		})
	}
}
