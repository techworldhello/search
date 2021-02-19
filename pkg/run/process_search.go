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
	json  format.Json
}

// NewSearchRepo returns an initialised instance of SearchRepo
func NewSearchRepo(d setup.Data) *SearchRepo {
	return &SearchRepo{Data: d}
}

// ProcessSearch returns search results based on the entity parameter
func (s SearchRepo) ProcessSearch(params param.Params) string {
	result := s.runSearch(params)
	if result == nil || result.GetSize() == 0 {
		return "No results found"
	}

	switch params.Format {
	case "table":
		s.table.Render(result)
	case "json":
		str, err := result.ToString()
		if err != nil {
			return "Error parsing result"
		}
		s.json.Render(str)
	default:
		return fmt.Sprintf("Cannot present data in %s format", params.Format)
	}

	return ""
}

func (s SearchRepo) runSearch(params param.Params) (sr search.Result) {
	var h Handler

	switch params.Entity {
	case "users":
		h = s.User
	case "tickets":
		h = s.Ticket
	case "organizations":
		h = s.Organization
	default:
		return sr
	}

	return h.Search(params.Term, params.Value)
}
