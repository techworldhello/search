package search

import (
	"github.com/techworldhello/search/pkg/schema"
	"strconv"
)

// Organization holds the organization data to search against
type Organization struct {
	Records []schema.Organization
}

// Search returns all organization results matching on term and value
func (organization Organization) Search(term, value string) Result {
	var orgs OrganizationResult

	for _, org := range organization.Records {
		if dataValue, found := toMap(org)[term]; found {
			var searchValue interface{}

			if value == "" {
				searchValue = ""
			} else {
				switch term {
				case "_id":
					searchValue, _ = strconv.Atoi(value)
				case "shared_tickets":
					if value == "true" || value == "false" {
						searchValue, _ = strconv.ParseBool(value)
					}
				case "domain_names":
					for _, domainName := range org.DomainNames {
						if domainName == value {
							dataValue = domainName
							searchValue = domainName
							break
						}
					}
				case "tags":
					for _, tag := range org.Tags {
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
				orgs.data = append(orgs.data, org)
				orgs.size++
			}
		}
	}

	return orgs
}
