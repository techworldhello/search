package search

import (
	"encoding/json"
	"github.com/techworldhello/search/pkg/schema"
	"strconv"
	"strings"
)

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

func (o OrganizationResult) ToString() (string, error) {
	bytes, err := json.Marshal(o.data)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
