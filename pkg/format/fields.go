package format

import (
	"fmt"
	"github.com/techworldhello/search/pkg/schema"
	"strings"
)

// SearchFields holds the types with searchable fields
type SearchFields struct {
	schema.User
	schema.Ticket
	schema.Organization
}

// NewSearchFields returns an initialised instance of SearchFields
func NewSearchFields() *SearchFields {
	return &SearchFields{}
}

// List returns all searchable fields
func (sf SearchFields) List() string {
	tagsMap := getHeaderTagsInMap(sf.User, sf.Ticket, sf.Organization)

	var output string
	for k, v := range tagsMap {
		title := fmt.Sprintf("Search %s with", k)
		jointFields := strings.Join(v[:], "\n")
		output += fmt.Sprintf("\n\n" +
			"%s\n" +
			"%s\n" +
			"%s\n" +
			"%s", strings.Repeat("*", len(title)), title, strings.Repeat("-", len(title)), jointFields)
	}

	return output
}
