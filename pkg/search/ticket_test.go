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
				data: []schema.Ticket{
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
				size: 1,
			},
		},
		{
			name:    "Found 1 result by Tags",
			keyTerm: "tags",
			value:   "American Samoa",
			expected: TicketResult{
				data: []schema.Ticket{
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
				size: 1,
			},
		},
		{
			name:    "Found 1 result by SubmitterID",
			keyTerm: "submitter_id",
			value:   "71",
			expected: TicketResult{
				data: []schema.Ticket{
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
				size: 1,
			},
		},
		{
			name:    "Found 2 results by Type",
			keyTerm: "type",
			value:   "incident",
			expected: TicketResult{
				data: []schema.Ticket{
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
				size: 2,
			},
		},
		{
			name:     "Not found — matched key, unmatched value",
			keyTerm:  "type",
			value:    "unmatched",
			expected: TicketResult{data: []schema.Ticket(nil), size: 0},
		},
		{
			name:     "Not found — unmatched key, matched value",
			keyTerm:  "not_found",
			value:    "1234",
			expected: TicketResult{data: []schema.Ticket(nil), size: 0},
		},
		{
			name:     "Not found — unmatched key & value",
			keyTerm:  "not_found",
			value:    "not_found",
			expected: TicketResult{data: []schema.Ticket(nil), size: 0},
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

func TestFormatTicketWorks(t *testing.T) {
	ticketResult := TicketResult{data: mocks.GetTicketTestData(), size: 3}

	result := [][]string{{"436bf9b0-1147-4c0a-8439-6f79833bff5b", "incident", "A Catastrophe in Korea (North)", "Nostrud ad sit velit cupidatat laboris ipsum nisi amet laboris ex exercitation amet et proident. Ipsum fugiat aute dolore tempor nostrud velit ipsum.", "high", "pending", "38", "24", "116", "Ohio,\nPennsylvania,\nAmerican Samoa,\nNorthern Mariana Islands", "false", "web"}, {"1a227508-9f39-427c-8f57-1b72f3fab87c", "incident", "A Catastrophe in Micronesia", "Aliquip excepteur fugiat ex minim ea aute eu labore. Sunt eiusmod esse eu non commodo est veniam consequat.", "low", "hold", "71", "38", "112", "Puerto Rico,\nIdaho,\nOklahoma,\nLouisiana", "false", "chat"}, {"2217c7dc-7371-4401-8738-0a8a8aedc08d", "problem", "A Catastrophe in Hungary", "Ipsum fugiat voluptate reprehenderit cupidatat aliqua dolore consequat. Consequat ullamco minim laboris veniam ea id laborum et eiusmod excepteur sint laborum dolore qui.", "normal", "closed", "9", "65", "105", "Massachusetts,\nNew York,\nMinnesota,\nNew Jersey", "true", "web"}}

	assert.Equal(t, result, ticketResult.Format())
}
