package setup

import (
	"encoding/json"

	"github.com/techworldhello/search/pkg/search"
)

// FileManager provides the ability to read a file
type FileManager interface {
	ReadFile() ([]byte, error)
}

// FilePaths holds the location of files
type FilePaths struct {
	User, Ticket, Organization string
}

// NewFilePaths returns an initialised instance of FilePaths
func NewFilePaths(user, ticket, organization string) *FilePaths {
	return &FilePaths{
		User:         user,
		Ticket:       ticket,
		Organization: organization,
	}
}

// Data is the data holder for search functionalities to run against
type Data struct {
	search.User
	search.Ticket
	search.Organization
}

// PrepareJSONData reads the content of each file and unmarshals JSON data to struct
func (fp FilePaths) PrepareJSONData() (d Data, err error) {
	users, err := fp.loadUsers(JSONReader{fp.User})
	if err != nil {
		return d, err
	}
	tickets, err := fp.loadTickets(JSONReader{fp.Ticket})
	if err != nil {
		return d, err
	}
	orgs, err := fp.loadOrganizations(JSONReader{fp.Organization})
	if err != nil {
		return d, err
	}
	return Data{
		users,
		tickets,
		orgs,
	}, nil
}

func (fp FilePaths) loadUsers(jr JSONReader) (search.User, error) {
	var users search.User

	content, err := jr.ReadFile()
	if err != nil {
		return users, err
	}
	if err := json.Unmarshal(content, &users.Records); err != nil {
		return users, err
	}
	return users, nil
}

func (fp FilePaths) loadTickets(jr JSONReader) (search.Ticket, error) {
	var tickets search.Ticket

	content, err := jr.ReadFile()
	if err != nil {
		return tickets, err
	}
	if err := json.Unmarshal(content, &tickets.Records); err != nil {
		return tickets, err
	}
	return tickets, nil
}

func (fp FilePaths) loadOrganizations(jr JSONReader) (search.Organization, error) {
	var orgs search.Organization

	content, err := jr.ReadFile()
	if err != nil {
		return orgs, err
	}
	if err := json.Unmarshal(content, &orgs.Records); err != nil {
		return orgs, err
	}
	return orgs, nil
}
