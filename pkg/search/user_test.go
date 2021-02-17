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
				[]schema.User{
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
			},
		},
		{
			name:    "Found 1 result by Name",
			keyTerm: "name",
			value:   "Francisca Rasmussen",
			expected: UserResult{
				[]schema.User{
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
			},
		},
		{
			name:    "Found 1 result by Tags",
			keyTerm: "tags",
			value:   "Glenshaw",
			expected: UserResult{
				[]schema.User{
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
			},
		},
		{
			name:    "Found 2 results by Suspended",
			keyTerm: "suspended",
			value:   "true",
			expected: UserResult{
				[]schema.User{
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
			},
		},
		{
			name:     "Not found — matched key, unmatched value by casing",
			keyTerm:  "timezone",
			value:    "netherlands",
			expected: UserResult{[]schema.User(nil)},
		},
		{
			name:     "Not found — matched key, unmatched value",
			keyTerm:  "tags",
			value:    "unmatched1234",
			expected: UserResult{[]schema.User(nil)},
		},
		{
			name:     "Not found — unmatched key, matched value",
			keyTerm:  "not_found",
			value:    "1234",
			expected: UserResult{[]schema.User(nil)},
		},
		{
			name:     "Not found — unmatched key & value",
			keyTerm:  "not_found",
			value:    "not_found",
			expected: UserResult{[]schema.User(nil)},
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
