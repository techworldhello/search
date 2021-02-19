package run

import (
	"github.com/techworldhello/search/pkg/param"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/techworldhello/search/pkg/mocks"
	"github.com/techworldhello/search/pkg/search"
	"github.com/techworldhello/search/pkg/setup"
)

func TestProcessSearch(t *testing.T) {
	expectations := []struct {
		name     string
		params   param.Params
		expected string
	}{
		{
			name: "User record is found",
			params: param.Params{
				Format: "table",
				Entity: "users",
				Term:   "_id",
				Value:  "4",
			},
			expected: "",
		},
		{
			name: "User record is not found on search value",
			params: param.Params{
				Format: "table",
				Entity: "users",
				Term:   "_id",
				Value:  "1234",
			},
			expected: "No results found",
		},
		{
			name: "Ticket record is found",
			params: param.Params{
				Format: "json",
				Entity: "tickets",
				Term:   "priority",
				Value:  "high",
			},
			expected: "",
		},
		{
			name: "Ticket record is not found on search key",
			params: param.Params{
				Format: "json",
				Entity: "tickets",
				Term:   "not_found",
				Value:  "true",
			},
			expected: "No results found",
		},
		{
			name: "Ticket record is not found on search value",
			params: param.Params{
				Format: "table",
				Entity: "tickets",
				Term:   "created_at",
				Value:  "not_found",
			},
			expected: "No results found",
		},
		{
			name: "Organization record is found",
			params: param.Params{
				Format: "table",
				Entity: "organizations",
				Term:   "tags",
				Value:  "Cherry",
			},
			expected: "",
		},
		{
			name: "Organization record is not found on search value",
			params: param.Params{
				Format: "json",
				Entity: "organizations",
				Term:   "_id",
				Value:  "1234",
			},
			expected: "No results found",
		},
		{
			name: "Record not found — entity does not exist",
			params: param.Params{
				Format: "table",
				Entity: "usr",
				Term:   "_id",
				Value:  "4",
			},
			expected: "No results found",
		},
		{
			name: "Output format is not supported",
			params: param.Params{
				Format: "xml",
				Entity: "users",
				Term:   "_id",
				Value:  "4",
			},
			expected: "Cannot present data in xml format",
		},
	}

	d := setup.Data{
		User:         search.User{Records: mocks.GetUserTestData()},
		Ticket:       search.Ticket{Records: mocks.GetTicketTestData()},
		Organization: search.Organization{Records: mocks.GetOrganizationTestData()},
	}

	a := NewSearchRepo(d)

	for _, expect := range expectations {
		t.Run(expect.name, func(t *testing.T) {
			assert.Equal(t, expect.expected, a.ProcessSearch(expect.params))
		})
	}
}
