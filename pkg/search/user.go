package search

import (
	"github.com/techworldhello/search/pkg/schema"
	"strconv"
)

// User holds the user data to search against
type User struct {
	Records []schema.User
}

type UserResult struct {
	data []schema.User
}

// Search returns all user results matching on term and value
func (user User) Search(term, value string) Result {
	var users UserResult

	for _, usr := range user.Records {
		if dataValue, found := toMap(usr)[term]; found {
			var searchValue interface{}

			if value == "" {
				searchValue = ""
			} else {
				switch term {
				case "_id", "organization_id":
					searchValue, _ = strconv.Atoi(value)
				case "active", "verified", "shared", "suspended":
					if value == "true" || value == "false" {
						searchValue, _ = strconv.ParseBool(value)
					}
				case "tags":
					for _, tag := range usr.Tags {
						if tag == value {
							dataValue = tag
							searchValue = tag
							break
						}
					}
				default:
					searchValue = value
				}

			}

			if dataValue == searchValue {
				users.data = append(users.data, usr)
			}
		}
	}

	return users
}
