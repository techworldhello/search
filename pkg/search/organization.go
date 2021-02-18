package search

import (
	"github.com/techworldhello/search/pkg/schema"
	"strconv"
	"strings"
)

// Organization holds the organization data to search against
type Organization struct {
	Records []schema.Organization
}

// OrganizationResult holds the organization search result
type OrganizationResult struct {
	data []schema.Organization
	size int
}

// GetSize returns the size of organization result
func (o OrganizationResult) GetSize() int {
	return o.size
}

// Format formats all organization results
func (o OrganizationResult) Format() (all [][]string) {
	for _, v := range o.data {
		var row []string

		row = append(row,
			strconv.Itoa(v.ID),
			v.Name,
			strings.Join(v.DomainNames, ",\n"),
			v.Details,
			strconv.FormatBool(v.SharedTickets),
			strings.Join(v.Tags, ",\n"),
		)
		all = append(all, row)
	}
	return all
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
