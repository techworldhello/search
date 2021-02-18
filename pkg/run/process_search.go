package run

import (
	"fmt"
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

// SearchRepo holds the data needed to execute search and list
type SearchRepo struct {
	setup.Data
}

// NewSearchRepo returns an initialised instance of Actions
func NewSearchRepo(d setup.Data) *SearchRepo {
	return &SearchRepo{Data: d}
}

// ProcessSearch returns search results based on the entity parameter
func (a SearchRepo) ProcessSearch(params param.Params) string {
	var h Handler

	switch params.Entity {
	case "users":
		h = a.User
	case "tickets":
		h = a.Ticket
	case "organizations":
		h = a.Organization
	default:
		return fmt.Sprintf("Data type %s does not exist", params.Entity)
	}

	result := h.Search(params.Term, params.Value)
	if result.GetSize() == 0 || result == nil {
		return "No results found"
	}

	return fmt.Sprintf("%+v", result)
}
