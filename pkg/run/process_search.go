package run

import (
	"fmt"
	"github.com/techworldhello/search/pkg/format"
	"github.com/techworldhello/search/pkg/param"
	"github.com/techworldhello/search/pkg/search"
	"github.com/techworldhello/search/pkg/setup"
)

// Searcher provides the ability to search for result
type Searcher interface {
	Search(term, value string) search.Result
}

// Handler is the interface that groups the search method
type Handler interface {
	Searcher
}

// SearchRepo holds the data needed to execute search
type SearchRepo struct {
	setup.Data
	table format.Table
}

// NewSearchRepo returns an initialised instance of SearchRepo
func NewSearchRepo(d setup.Data) *SearchRepo {
	return &SearchRepo{Data: d}
}

// ProcessSearch returns search results based on the entity parameter
func (s SearchRepo) ProcessSearch(params param.Params) string {
	var h Handler

	switch params.Entity {
	case "users":
		h = s.User
	case "tickets":
		h = s.Ticket
	case "organizations":
		h = s.Organization
	default:
		return fmt.Sprintf("Data type %s does not exist", params.Entity)
	}

	result := h.Search(params.Term, params.Value)
	if result.GetSize() == 0 || result == nil {
		return "No results found"
	}

	s.table.Render(result)

	return ""
}
