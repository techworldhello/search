package schema

// Organization represents the shape of organization data
type Organization struct {
	ID            int      `json:"_id" header:"ID"`
	URL           string   `json:"url"`
	ExternalID    string   `json:"external_id"`
	Name          string   `json:"name" header:"Name"`
	DomainNames   []string `json:"domain_names" header:"DomainNames"`
	CreatedAt     string   `json:"created_at" header:"CreatedAt"`
	Details       string   `json:"details" header:"Details"`
	SharedTickets bool     `json:"shared_tickets" header:"SharedTickets"`
	Tags          []string `json:"tags" header:"Tags"`
}
