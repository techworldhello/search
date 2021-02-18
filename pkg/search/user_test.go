package search

import (
	"github.com/stretchr/testify/assert"
	"github.com/techworldhello/search/pkg/mocks"
	"github.com/techworldhello/search/pkg/schema"
	"testing"
)

func TestSearchUser(t *testing.T) {
	expectations := []struct {
		name     string
		keyTerm  string
		value    string
		expected UserResult
	}{
		{
			name:    "Found 1 result by ID",
			keyTerm: "_id",
			value:   "4",
			expected: UserResult{
				data:[]schema.User{
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
				},
				size: 1,
			},
		},
		{
			name:    "Found 1 result by Name",
			keyTerm: "name",
			value:   "Francisca Rasmussen",
			expected: UserResult{
				data: []schema.User{
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
				},
				size: 1,
			},
		},
		{
			name:    "Found 1 result by Tags",
			keyTerm: "tags",
			value:   "Glenshaw",
			expected: UserResult{
				data: []schema.User{
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
				},
				size: 1,
			},
		},
		{
			name:    "Found 2 results by Suspended",
			keyTerm: "suspended",
			value:   "true",
			expected: UserResult{
				data: []schema.User{
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
				},
				size: 2,
			},
		},
		{
			name:     "Not found — matched key, unmatched value by casing",
			keyTerm:  "timezone",
			value:    "netherlands",
			expected: UserResult{data: []schema.User(nil), size: 0},
		},
		{
			name:     "Not found — matched key, unmatched value",
			keyTerm:  "tags",
			value:    "unmatched1234",
			expected: UserResult{data: []schema.User(nil), size: 0},
		},
		{
			name:     "Not found — unmatched key, matched value",
			keyTerm:  "not_found",
			value:    "1234",
			expected: UserResult{data: []schema.User(nil), size: 0},
		},
		{
			name:     "Not found — unmatched key & value",
			keyTerm:  "not_found",
			value:    "not_found",
			expected: UserResult{data: []schema.User(nil), size: 0},
		},
	}

	user := User{mocks.GetUserTestData()}

	for _, expect := range expectations {
		t.Run(expect.name, func(t *testing.T) {

			searchResult := user.Search(expect.keyTerm, expect.value)

			assert.Equal(t, expect.expected, searchResult)
		})
	}
}

func TestFormatUserWorks(t *testing.T) {
	userResult := UserResult{data: mocks.GetUserTestData(), size: 4}

	result := [][]string{{"1", "Francisca Rasmussen", "Miss Coffey", "true", "true", "false", "en-AU", "Sri Lanka", "coffeyrasmussen@flotonic.com", "8335-422-718", "Don't Worry Be Happy!", "119", "Springville,\nSutton,\nHartsville/Hartley,\nDiaperville", "true", "admin"}, {"2", "Cross Barlow", "Miss Joni", "true", "true", "false", "zh-CN", "Armenia", "jonibarlow@flotonic.com", "9575-552-585", "Don't Worry Be Happy!", "106", "Foxworth,\nWoodlands,\nHerlong,\nHenrietta", "false", "admin"}, {"3", "Ingrid Wagner", "Miss Buck", "false", "false", "false", "en-AU", "Trinidad and Tobago", "buckwagner@flotonic.com", "9365-482-943", "Don't Worry Be Happy!", "104", "Mulino,\nKenwood,\nWescosville,\nLoyalhanna", "false", "end-user"}, {"4", "Rose Newton", "Mr Cardenas", "true", "true", "true", "de-CH", "Netherlands", "cardenasnewton@flotonic.com", "8685-482-450", "Don't Worry Be Happy!", "122", "Gallina,\nGlenshaw,\nRowe,\nBabb", "true", "end-user"}}

	assert.Equal(t, result, userResult.Format())
}
