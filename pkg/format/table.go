package format

import (
	"github.com/olekukonko/tablewriter"
	"github.com/techworldhello/search/pkg/schema"
	"github.com/techworldhello/search/pkg/search"
	"os"
)

// Table represents the type to render
type Table struct {}

// Render writes the search result to stdout in table format
func (t Table) Render(result search.Result) {
	table := tablewriter.NewWriter(os.Stdout)

	switch result.(type) {
	case search.UserResult:
		table.SetHeader(getHeaderTags(schema.User{}))
	case search.TicketResult:
		table.SetHeader(getHeaderTags(schema.Ticket{}))
	case search.OrganizationResult:
		table.SetHeader(getHeaderTags(schema.Organization{}))
	}

	table.SetCenterSeparator("*")
	table.SetColumnSeparator("╪")
	table.SetRowSeparator("-")

	table.SetAlignment(tablewriter.ALIGN_LEFT)

	table.AppendBulk(result.Format())
	table.Render()
}
