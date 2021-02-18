package search

import (
	"github.com/techworldhello/search/pkg/schema"
	"strconv"
	"strings"
)

// Ticket holds the ticket data to search against
type Ticket struct {
	Records []schema.Ticket
}

// TicketResult holds the ticket search result
type TicketResult struct {
	data []schema.Ticket
	size int
}

func (t TicketResult) GetSize() int {
	return t.size
}

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
			v.DueAt,
			v.Via,
		)
		all = append(all, row)
	}
	return all
}

// Search returns all ticket results matching on term and value
func (ticket Ticket) Search(term, value string) Result {
	var tickets TicketResult

	for _, ticket := range ticket.Records {
		if dataValue, found := toMap(ticket)[term]; found {
			var searchValue interface{}

			if value == "" {
				searchValue = ""
			} else {
				switch term {
				case "submitter_id", "assignee_id", "organization_id":
					searchValue, _ = strconv.Atoi(value)
				case "has_incidents":
					if value == "true" || value == "false" {
						searchValue, _ = strconv.ParseBool(value)
					}
				case "tags":
					for _, tag := range ticket.Tags {
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
				tickets.data = append(tickets.data, ticket)
				tickets.size++
			}
		}
	}

	return tickets
}
