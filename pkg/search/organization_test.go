package search

import (
	"github.com/stretchr/testify/assert"
	"github.com/techworldhello/search/pkg/mocks"
	"github.com/techworldhello/search/pkg/schema"
	"testing"
)

func TestSearch(t *testing.T) {
	expectations := []struct {
		name     string
		keyTerm  string
		value    string
		expected OrganizationResult
	}{
		{
			name:    "Found 1 result by ID",
			keyTerm: "_id",
			value:   "102",
			expected: OrganizationResult{
				data: []schema.Organization{
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
				},
				size: 1,
			},
		},
		{
			name:    "Found 1 result by Name",
			keyTerm: "name",
			value:   "Enthaze",
			expected: OrganizationResult{
				data: []schema.Organization{
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
				},
				size: 1,
			},
		},
		{
			name:    "Found 1 result by DomainNames",
			keyTerm: "domain_names",
			value:   "bluegrain.com",
			expected: OrganizationResult{
				data: []schema.Organization{
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
				},
				size: 1,
			},
		},
		{
			name:    "Found 1 result by Tags — matched key & value",
			keyTerm: "tags",
			value:   "Armstrong",
			expected: OrganizationResult{
				data: []schema.Organization{
					{
						ID:            103,
						URL:           "http://initech.zendesk.com/api/v2/organizations/103.json",
						ExternalID:    "e73240f3-8ecf-411d-ad0d-80ca8a84053d",
						Name:          "Plasmos",
						DomainNames:   []string{"comvex.com", "automon.com", "verbus.com", "gogol.com"},
						CreatedAt:     "2016-05-28T04:40:37 -10:00",
						Details:       "",
						SharedTickets: true,
						Tags:          []string{"Parrish", "Lindsay", "Armstrong", "Vaughn"},
					},
				},
				size: 1,
			},
		},
		{
			name:    "Found 1 result by Details — matched key & empty value",
			keyTerm: "details",
			value:   "",
			expected: OrganizationResult{
				data: []schema.Organization{
					{
						ID:            103,
						URL:           "http://initech.zendesk.com/api/v2/organizations/103.json",
						ExternalID:    "e73240f3-8ecf-411d-ad0d-80ca8a84053d",
						Name:          "Plasmos",
						DomainNames:   []string{"comvex.com", "automon.com", "verbus.com", "gogol.com"},
						CreatedAt:     "2016-05-28T04:40:37 -10:00",
						Details:       "",
						SharedTickets: true,
						Tags:          []string{"Parrish", "Lindsay", "Armstrong", "Vaughn"},
					},
				},
				size: 1,
			},
		},
		{
			name:    "Found 2 results by SharedTickets — matched key & value",
			keyTerm: "shared_tickets",
			value:   "false",
			expected: OrganizationResult{
				data: []schema.Organization{
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
				},
				size: 2,
			},
		},
		{
			name:     "Not found — matched key, unmatched value",
			keyTerm:  "shared_tickets",
			value:    "unmatched",
			expected: OrganizationResult{data: []schema.Organization(nil), size: 0},
		},
		{
			name:     "Not found — unmatched key, matched value",
			keyTerm:  "not_found",
			value:    "101",
			expected: OrganizationResult{data: []schema.Organization(nil), size: 0},
		},
		{
			name:     "Not found — unmatched key & value",
			keyTerm:  "not_found",
			value:    "not_found",
			expected: OrganizationResult{data: []schema.Organization(nil), size: 0},
		},
	}

	org := Organization{mocks.GetOrganizationTestData()}

	for _, expect := range expectations {
		t.Run(expect.name, func(t *testing.T) {

			searchResult := org.Search(expect.keyTerm, expect.value)

			assert.Equal(t, expect.expected, searchResult)
		})
	}
}

func TestFormatOrganizationWorks(t *testing.T) {
	orgsResult := OrganizationResult{data: mocks.GetOrganizationTestData(), size: 3}

	result := [][]string{{"101", "Enthaze", "kage.com,\necratic.com,\nendipin.com,\nzentix.com", "MegaCorp", "false", "Fulton,\nWest,\nRodriguez,\nFarley"}, {"102", "Nutralab", "trollery.com,\ndatagen.com,\nbluegrain.com,\ndadabase.com", "Non profit", "false", "Cherry,\nCollier,\nFuentes,\nTrevino"}, {"103", "Plasmos", "comvex.com,\nautomon.com,\nverbus.com,\ngogol.com", "", "true", "Parrish,\nLindsay,\nArmstrong,\nVaughn"}}

	assert.Equal(t, result, orgsResult.Format())
}
