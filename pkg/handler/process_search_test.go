package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/techworldhello/search/pkg/mocks"
	"github.com/techworldhello/search/pkg/search"
	"github.com/techworldhello/search/pkg/setup"
)

func TestProcessSearch(t *testing.T) {
	expectations := []struct {
		name     string
		params   Params
		expected string
	}{
		{
			name: "User record is found",
			params: Params{
				format: "table",
				entity: "users",
				term:   "_id",
				value:  "4",
			},
			expected: "",
		},
		{
			name: "User record is not found on search value",
			params: Params{
				format: "table",
				entity: "users",
				term:   "_id",
				value:  "1234",
			},
			expected: "No results found",
		},
		{
			name: "Ticket record is found",
			params: Params{
				format: "json",
				entity: "tickets",
				term:   "priority",
				value:  "high",
			},
			expected: "",
		},
		{
			name: "Ticket record is not found on search key",
			params: Params{
				format: "json",
				entity: "tickets",
				term:   "not_found",
				value:  "true",
			},
			expected: "No results found",
		},
		{
			name: "Ticket record is not found on search value",
			params: Params{
				format: "table",
				entity: "tickets",
				term:   "created_at",
				value:  "not_found",
			},
			expected: "No results found",
		},
		{
			name: "Organization record is found",
			params: Params{
				format: "table",
				entity: "organizations",
				term:   "tags",
				value:  "Cherry",
			},
			expected: "",
		},
		{
			name: "Organization record is not found on search value",
			params: Params{
				format: "json",
				entity: "organizations",
				term:   "_id",
				value:  "1234",
			},
			expected: "No results found",
		},
		{
			name: "Record not found — entity does not exist",
			params: Params{
				format: "table",
				entity: "usr",
				term:   "_id",
				value:  "4",
			},
			expected: "No results found",
		},
		{
			name: "Output format is not supported",
			params: Params{
				format: "xml",
				entity: "users",
				term:   "_id",
				value:  "4",
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
