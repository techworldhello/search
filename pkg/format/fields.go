package format

import (
	"fmt"
	"github.com/techworldhello/search/pkg/schema"
	"strings"
)

// Fields holds the types with searchable fields
type Fields struct {
	schema.User
	schema.Ticket
	schema.Organization
}

// NewFields returns an initialised instance of Fields
func NewFields() *Fields {
	return &Fields{}
}

// List returns all searchable fields
func (f Fields) List() string {
	tagsMap := getHeaderTagsInMap(f.User, f.Ticket, f.Organization)

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
