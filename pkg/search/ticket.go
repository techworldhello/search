package search

import (
	"github.com/techworldhello/search/pkg/schema"
	"strconv"
)

// Ticket holds the ticket data to search against
type Ticket struct {
	Records []schema.Ticket
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
