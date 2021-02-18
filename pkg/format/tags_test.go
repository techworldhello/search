package format

import (
	"github.com/stretchr/testify/assert"
	"github.com/techworldhello/search/pkg/schema"
	"testing"
)

func TestGetHeaderTags(t *testing.T) {
	expectations := []struct {
		name     string
		entity   interface{}
		expected []string
	}{
		{
			name: "User record header tags are correctly returned",
			entity: schema.User{},
			expected: []string{"ID", "Name", "Alias", "Active", "Verified", "Shared", "Locale", "Timezone", "Email", "Phone", "Signature", "OrganizationID", "Tags", "Suspended", "Role"},
		},
		{
			name: "Ticket record header tags are correctly returned",
			entity: schema.Ticket{},
			expected: []string{"ID", "Type", "Subject", "Description", "Priority", "Status", "SubmitterID", "AssigneeID", "OrganizationID", "Tags", "HasIncidents", "DueAt", "Via"},
		},
		{
			name: "Organization record header tags are correctly returned",
			entity: schema.Organization{},
			expected: []string{"ID", "Name", "DomainNames", "CreatedAt", "Details", "SharedTickets", "Tags"},
		},
		{
			name: "Organization record header tags are correctly returned",
			entity: struct{}{},
			expected: []string(nil),
		},
	}

	for _, expect := range expectations {
		t.Run(expect.name, func(t *testing.T) {
			assert.Equal(t, expect.expected, getHeaderTags(expect.entity))
		})
	}
}
