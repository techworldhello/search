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
				Entity: "users",
				Term:   "_id",
				Value:  "4",
			},
			expected: expectedUserRecord,
		},
		{
			name: "User record is not found on search value",
			params: param.Params{
				Entity: "users",
				Term:   "_id",
				Value:  "1234",
			},
			expected: "No results found",
		},
		{
			name: "Ticket record is found",
			params: param.Params{
				Entity: "tickets",
				Term:   "priority",
				Value:  "high",
			},
			expected: expectedTicketRecord,
		},
		{
			name: "Ticket record is not found on search key",
			params: param.Params{
				Entity: "tickets",
				Term:   "not_found",
				Value:  "true",
			},
			expected: "No results found",
		},
		{
			name: "Ticket record is not found on search value",
			params: param.Params{
				Entity: "tickets",
				Term:   "created_at",
				Value:  "not_found",
			},
			expected: "No results found",
		},
		{
			name: "Organization record is found",
			params: param.Params{
				Entity: "organizations",
				Term:   "tags",
				Value:  "Cherry",
			},
			expected: expectedOrgRecord,
		},
		{
			name: "Organization record is not found on search value",
			params: param.Params{
				Entity: "organizations",
				Term:   "_id",
				Value:  "1234",
			},
			expected: "No results found",
		},
		{
			name: "Record not found — entity does not exist",
			params: param.Params{
				Entity: "usr",
				Term:   "_id",
				Value:  "4",
			},
			expected: "Data type usr does not exist",
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

const expectedUserRecord = `{data:[{ID:4 URL:http://initech.zendesk.com/api/v2/users/4.json ExternalID:37c9aef5-cf01-4b07-af24-c6c49ac1d1c7 Name:Rose Newton Alias:Mr Cardenas CreatedAt:2016-02-09T07:52:10 -11:00 Active:true Verified:true Shared:true Locale:de-CH Timezone:Netherlands LastLoginAt:2012-09-25T01:32:46 -10:00 Email:cardenasnewton@flotonic.com Phone:8685-482-450 Signature:Don't Worry Be Happy! OrganizationID:122 Tags:[Gallina Glenshaw Rowe Babb] Suspended:true Role:end-user}] size:1}`

const expectedTicketRecord = `{data:[{ID:436bf9b0-1147-4c0a-8439-6f79833bff5b URL:http://initech.zendesk.com/api/v2/tickets/436bf9b0-1147-4c0a-8439-6f79833bff5b.json ExternalID:9210cdc9-4bee-485f-a078-35396cd74063 CreatedAt:2016-04-28T11:19:34 -10:00 Type:incident Subject:A Catastrophe in Korea (North) Description:Nostrud ad sit velit cupidatat laboris ipsum nisi amet laboris ex exercitation amet et proident. Ipsum fugiat aute dolore tempor nostrud velit ipsum. Priority:high Status:pending SubmitterID:38 AssigneeID:24 OrganizationID:116 Tags:[Ohio Pennsylvania American Samoa Northern Mariana Islands] HasIncidents:false DueAt:2016-07-31T02:37:50 -10:00 Via:web}] size:1}`

const expectedOrgRecord = `{data:[{ID:102 URL:http://initech.zendesk.com/api/v2/organizations/102.json ExternalID:7cd6b8d4-2999-4ff2-8cfd-44d05b449226 Name:Nutralab DomainNames:[trollery.com datagen.com bluegrain.com dadabase.com] CreatedAt:2016-04-07T08:21:44 -10:00 Details:Non profit SharedTickets:false Tags:[Cherry Collier Fuentes Trevino]}] size:1}`
