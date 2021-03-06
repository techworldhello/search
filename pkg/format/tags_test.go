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
			expected: []string{"ID", "Type", "Subject", "Description", "Priority", "Status", "SubmitterID", "AssigneeID", "OrganizationID", "Tags", "HasIncidents", "Via"},
		},
		{
			name: "Organization record header tags are correctly returned",
			entity: schema.Organization{},
			expected: []string{"ID", "Name", "DomainNames", "Details", "SharedTickets", "Tags"},
		},
		{
			name: "No tags are returned as no properties exist on struct",
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

func TestGetHeaderTagsInMap(t *testing.T) {
	expectations := []struct {
		name     string
		entity   interface{}
		expected map[string][]string
	}{
		{
			name: "User record json tags are correctly returned",
			entity: schema.User{},
			expected: map[string][]string{"User":{"_id", "name", "alias", "active", "verified", "shared", "locale", "timezone", "email", "phone", "signature", "organization_id", "tags", "suspended", "role"}},
		},
		{
			name: "Ticket record json tags are correctly returned",
			entity: schema.Ticket{},
			expected: map[string][]string{"Ticket":{"_id", "type", "subject", "description", "priority", "status", "submitter_id", "assignee_id", "organization_id", "tags", "has_incidents", "via"}},
		},
		{
			name: "Organization record json tags are correctly returned",
			entity: schema.Organization{},
			expected: map[string][]string{"Organization":{"_id", "name", "domain_names", "details", "shared_tickets", "tags"}},
		},
		{
			name: "No tags are returned as no properties exist on struct",
			entity: struct{}{},
			expected: map[string][]string(nil),
		},
	}

	for _, expect := range expectations {
		t.Run(expect.name, func(t *testing.T) {
			assert.Equal(t, expect.expected, getHeaderTagsInMap(expect.entity))
		})
	}
}
