package setup

import (
	"encoding/json"
	"fmt"
	"sync"

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
	var wg sync.WaitGroup

	userChan := make(chan search.User, 1)
	userErrChan := make(chan error, 1)
	wg.Add(1)
	go fp.loadUsers(JSONReader{fp.User}, userChan, userErrChan, &wg)
	userErr := <-userErrChan

	ticketChan := make(chan search.Ticket, 1)
	ticketErrChan := make(chan error, 1)
	wg.Add(1)
	go fp.loadTickets(JSONReader{fp.Ticket}, ticketChan, ticketErrChan, &wg)
	ticketErr := <-ticketErrChan

	orgChan := make(chan search.Organization, 1)
	orgErrChan := make(chan error, 1)
	wg.Add(1)
	go fp.loadOrganizations(JSONReader{fp.Organization}, orgChan, orgErrChan, &wg)
	orgErr := <- orgErrChan

	wg.Wait()

	if userErr != nil || ticketErr != nil || orgErr != nil {
		return d, fmt.Errorf("Error loading files.\n" +
			"User: %+v,\n" +
			"Ticket: %+v,\n" +
			"Organization: %+v\n", userErr, ticketErr, orgErr)
	}

	return Data{
		<- userChan,
		<- ticketChan,
		<- orgChan,
	}, nil
}

func (fp FilePaths) loadUsers(jr JSONReader, outChan chan search.User, errChan chan error, wg *sync.WaitGroup) {
	defer close(outChan)
	defer close(errChan)
	defer wg.Done()

	var users search.User

	content, err := jr.ReadFile()
	if err != nil {
		errChan <- err
		return
	}
	if err := json.Unmarshal(content, &users.Records); err != nil {
		errChan <- err
		return
	}

	outChan <- users
}

func (fp FilePaths) loadTickets(jr JSONReader, outChan chan search.Ticket, errChan chan error, wg *sync.WaitGroup) {
	defer close(outChan)
	defer close(errChan)
	defer wg.Done()

	var tickets search.Ticket

	content, err := jr.ReadFile()
	if err != nil {
		errChan <- err
		return
	}
	if err := json.Unmarshal(content, &tickets.Records); err != nil {
		errChan <- err
		return
	}

	outChan <- tickets
}

func (fp FilePaths) loadOrganizations(jr JSONReader, outChan chan search.Organization, errChan chan error, wg *sync.WaitGroup) {
	defer close(outChan)
	defer close(errChan)
	defer wg.Done()

	var orgs search.Organization

	content, err := jr.ReadFile()
	if err != nil {
		errChan <- err
		return
	}
	if err := json.Unmarshal(content, &orgs.Records); err != nil {
		errChan <- err
		return
	}

	outChan <- orgs
}
