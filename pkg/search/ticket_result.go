package search

import (
	"encoding/json"
	"github.com/techworldhello/search/pkg/schema"
	"strconv"
	"strings"
)

// TicketResult holds the ticket search result
type TicketResult struct {
	data []schema.Ticket
	size int
}

// GetSize returns the size of ticket result
func (t TicketResult) GetSize() int {
	return t.size
}

// Format formats all ticket results
func (t TicketResult) Format() (all [][]string) {
	for _, v := range t.data {
		var row []string

		row = append(row,
			v.ID,
			v.Type,
			v.Subject,
			v.Description,
			v.Priority,
			v.Status,
			strconv.Itoa(v.SubmitterID),
			strconv.Itoa(v.AssigneeID),
			strconv.Itoa(v.OrganizationID),
			strings.Join(v.Tags, ",\n"),
			strconv.FormatBool(v.HasIncidents),
			v.Via,
		)
		all = append(all, row)
	}
	return all
}

func (t TicketResult) ToString() (string, error) {
	bytes, err := json.Marshal(t.data)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
