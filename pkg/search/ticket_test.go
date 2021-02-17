package search

import (
	"github.com/stretchr/testify/assert"
	"github.com/techworldhello/search/pkg/mocks"
	"github.com/techworldhello/search/pkg/schema"
	"testing"
)

func TestSearchTicket(t *testing.T) {
	expectations := []struct {
		name     string
		keyTerm  string
		value    string
		expected TicketResult
	}{
		{
			name:    "Found 1 result by ID",
			keyTerm: "_id",
			value:   "436bf9b0-1147-4c0a-8439-6f79833bff5b",
			expected: TicketResult{
				[]schema.Ticket{
					{
						ID:             "436bf9b0-1147-4c0a-8439-6f79833bff5b",
						URL:            "http://initech.zendesk.com/api/v2/tickets/436bf9b0-1147-4c0a-8439-6f79833bff5b.json",
						ExternalID:     "9210cdc9-4bee-485f-a078-35396cd74063",
						CreatedAt:      "2016-04-28T11:19:34 -10:00",
						Type:           "incident",
						Subject:        "A Catastrophe in Korea (North)",
						Description:    "Nostrud ad sit velit cupidatat laboris ipsum nisi amet laboris ex exercitation amet et proident. Ipsum fugiat aute dolore tempor nostrud velit ipsum.",
						Priority:       "high",
						Status:         "pending",
						SubmitterID:    38,
						AssigneeID:     24,
						OrganizationID: 116,
						Tags:           []string{"Ohio", "Pennsylvania", "American Samoa", "Northern Mariana Islands"},
						HasIncidents:   false,
						DueAt:          "2016-07-31T02:37:50 -10:00",
						Via:            "web",
					},
				},
			},
		},
		{
			name:    "Found 1 result by Tags",
			keyTerm: "tags",
			value:   "American Samoa",
			expected: TicketResult{
				[]schema.Ticket{
					{
						ID:             "436bf9b0-1147-4c0a-8439-6f79833bff5b",
						URL:            "http://initech.zendesk.com/api/v2/tickets/436bf9b0-1147-4c0a-8439-6f79833bff5b.json",
						ExternalID:     "9210cdc9-4bee-485f-a078-35396cd74063",
						CreatedAt:      "2016-04-28T11:19:34 -10:00",
						Type:           "incident",
						Subject:        "A Catastrophe in Korea (North)",
						Description:    "Nostrud ad sit velit cupidatat laboris ipsum nisi amet laboris ex exercitation amet et proident. Ipsum fugiat aute dolore tempor nostrud velit ipsum.",
						Priority:       "high",
						Status:         "pending",
						SubmitterID:    38,
						AssigneeID:     24,
						OrganizationID: 116,
						Tags:           []string{"Ohio", "Pennsylvania", "American Samoa", "Northern Mariana Islands"},
						HasIncidents:   false,
						DueAt:          "2016-07-31T02:37:50 -10:00",
						Via:            "web",
					},
				},
			},
		},
		{
			name:    "Found 1 result by SubmitterID",
			keyTerm: "submitter_id",
			value:   "71",
			expected: TicketResult{
				[]schema.Ticket{
					{
						ID:             "1a227508-9f39-427c-8f57-1b72f3fab87c",
						URL:            "http://initech.zendesk.com/api/v2/tickets/1a227508-9f39-427c-8f57-1b72f3fab87c.json",
						ExternalID:     "3e5ca820-cd1f-4a02-a18f-11b18e7bb49a",
						CreatedAt:      "2016-04-14T08:32:31 -10:00",
						Type:           "incident",
						Subject:        "A Catastrophe in Micronesia",
						Description:    "Aliquip excepteur fugiat ex minim ea aute eu labore. Sunt eiusmod esse eu non commodo est veniam consequat.",
						Priority:       "low",
						Status:         "hold",
						SubmitterID:    71,
						AssigneeID:     38,
						OrganizationID: 112,
						Tags:           []string{"Puerto Rico", "Idaho", "Oklahoma", "Louisiana"},
						HasIncidents:   false,
						DueAt:          "2016-08-15T05:37:32 -10:00",
						Via:            "chat",
					},
				},
			},
		},
		{
			name:    "Found 2 results by Type",
			keyTerm: "type",
			value:   "incident",
			expected: TicketResult{
				[]schema.Ticket{
					{
						ID:             "436bf9b0-1147-4c0a-8439-6f79833bff5b",
						URL:            "http://initech.zendesk.com/api/v2/tickets/436bf9b0-1147-4c0a-8439-6f79833bff5b.json",
						ExternalID:     "9210cdc9-4bee-485f-a078-35396cd74063",
						CreatedAt:      "2016-04-28T11:19:34 -10:00",
						Type:           "incident",
						Subject:        "A Catastrophe in Korea (North)",
						Description:    "Nostrud ad sit velit cupidatat laboris ipsum nisi amet laboris ex exercitation amet et proident. Ipsum fugiat aute dolore tempor nostrud velit ipsum.",
						Priority:       "high",
						Status:         "pending",
						SubmitterID:    38,
						AssigneeID:     24,
						OrganizationID: 116,
						Tags:           []string{"Ohio", "Pennsylvania", "American Samoa", "Northern Mariana Islands"},
						HasIncidents:   false,
						DueAt:          "2016-07-31T02:37:50 -10:00",
						Via:            "web",
					},
					{
						ID:             "1a227508-9f39-427c-8f57-1b72f3fab87c",
						URL:            "http://initech.zendesk.com/api/v2/tickets/1a227508-9f39-427c-8f57-1b72f3fab87c.json",
						ExternalID:     "3e5ca820-cd1f-4a02-a18f-11b18e7bb49a",
						CreatedAt:      "2016-04-14T08:32:31 -10:00",
						Type:           "incident",
						Subject:        "A Catastrophe in Micronesia",
						Description:    "Aliquip excepteur fugiat ex minim ea aute eu labore. Sunt eiusmod esse eu non commodo est veniam consequat.",
						Priority:       "low",
						Status:         "hold",
						SubmitterID:    71,
						AssigneeID:     38,
						OrganizationID: 112,
						Tags:           []string{"Puerto Rico", "Idaho", "Oklahoma", "Louisiana"},
						HasIncidents:   false,
						DueAt:          "2016-08-15T05:37:32 -10:00",
						Via:            "chat",
					},
				},
			},
		},
		{
			name:     "Not found — matched key, unmatched value",
			keyTerm:  "type",
			value:    "unmatched",
			expected: TicketResult{[]schema.Ticket(nil)},
		},
		{
			name:     "Not found — unmatched key, matched value",
			keyTerm:  "not_found",
			value:    "1234",
			expected: TicketResult{[]schema.Ticket(nil)},
		},
		{
			name:     "Not found — unmatched key & value",
			keyTerm:  "not_found",
			value:    "not_found",
			expected: TicketResult{[]schema.Ticket(nil)},
		},
	}

	tix := Ticket{mocks.GetTicketTestData()}

	for _, expect := range expectations {
		t.Run(expect.name, func(t *testing.T) {

			searchResult := tix.Search(expect.keyTerm, expect.value)

			assert.Equal(t, expect.expected, searchResult)
		})
	}
}