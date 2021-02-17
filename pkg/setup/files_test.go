package setup

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/techworldhello/search/pkg/schema"
	"github.com/techworldhello/search/pkg/search"
)

func TestPrepareJSONDataWorks(t *testing.T) {
	f := NewFilePaths("../data/test-users.json", "../data/test-tickets.json", "../data/test-organizations.json")
	data, err := f.PrepareJSONData()

	expected := Data{
		User: search.User{Records: []schema.User{
			{
				ID:             1,
				URL:            "http://initech.zendesk.com/api/v2/users/1.json",
				ExternalID:     "74341f74-9c79-49d5-9611-87ef9b6eb75f",
				Name:           "Francisca Rasmussen",
				Alias:          "Miss Coffey",
				CreatedAt:      "2016-04-15T05:19:46 -10:00",
				Active:         true,
				Verified:       true,
				Shared:         false,
				Locale:         "en-AU",
				Timezone:       "Sri Lanka",
				LastLoginAt:    "2013-08-04T01:03:27 -10:00",
				Email:          "coffeyrasmussen@flotonic.com",
				Phone:          "8335-422-718",
				Signature:      "Don't Worry Be Happy!",
				OrganizationID: 119,
				Tags:           []string{"Springville", "Sutton", "Hartsville/Hartley", "Diaperville"},
				Suspended:      true,
				Role:           "admin",
			},
			{
				ID:             4,
				URL:            "http://initech.zendesk.com/api/v2/users/4.json",
				ExternalID:     "37c9aef5-cf01-4b07-af24-c6c49ac1d1c7",
				Name:           "Rose Newton",
				Alias:          "Mr Cardenas",
				CreatedAt:      "2016-02-09T07:52:10 -11:00",
				Active:         true,
				Verified:       true,
				Shared:         true,
				Locale:         "de-CH",
				Timezone:       "Netherlands",
				LastLoginAt:    "2012-09-25T01:32:46 -10:00",
				Email:          "cardenasnewton@flotonic.com",
				Phone:          "8685-482-450",
				Signature:      "Don't Worry Be Happy!",
				OrganizationID: 122,
				Tags:           []string{"Gallina", "Glenshaw", "Rowe", "Babb"},
				Suspended:      true,
				Role:           "end-user",
			},
		}},
		Ticket: search.Ticket{Records: []schema.Ticket{
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
		}},
		Organization: search.Organization{Records: []schema.Organization{
			{
				ID:            101,
				URL:           "http://initech.zendesk.com/api/v2/organizations/101.json",
				ExternalID:    "9270ed79-35eb-4a38-a46f-35725197ea8d",
				Name:          "Enthaze",
				DomainNames:   []string{"kage.com", "ecratic.com", "endipin.com", "zentix.com"},
				CreatedAt:     "2016-05-21T11:10:28 -10:00",
				Details:       "MegaCorp",
				SharedTickets: false,
				Tags:          []string{"Fulton", "West", "Rodriguez", "Farley"},
			},
			{
				ID:            102,
				URL:           "http://initech.zendesk.com/api/v2/organizations/102.json",
				ExternalID:    "7cd6b8d4-2999-4ff2-8cfd-44d05b449226",
				Name:          "Nutralab",
				DomainNames:   []string{"trollery.com", "datagen.com", "bluegrain.com", "dadabase.com"},
				CreatedAt:     "2016-04-07T08:21:44 -10:00",
				Details:       "Non profit",
				SharedTickets: false,
				Tags:          []string{"Cherry", "Collier", "Fuentes", "Trevino"},
			},
		}},
	}

	assert.Equal(t, expected, data)
	assert.Equal(t, nil, err)
}

func TestPrepareJSONDataCannotFindFile(t *testing.T) {
	f := FilePaths{
		User:         "../data/cannot-find.json",
		Ticket:       "../data/test-tickets.json",
		Organization: "../data/test-organizations.json",
	}

	data, err := f.PrepareJSONData()

	assert.Equal(t, Data{}, data)
	assert.Contains(t, err.Error(), "no such file or directory")
}
