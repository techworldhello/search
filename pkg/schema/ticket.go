package schema

// Ticket represents the shape of ticket data
type Ticket struct {
	ID             string   `json:"_id" header:"ID"`
	URL            string   `json:"url"`
	ExternalID     string   `json:"external_id"`
	CreatedAt      string   `json:"created_at"`
	Type           string   `json:"type" header:"Type"`
	Subject        string   `json:"subject" header:"Subject"`
	Description    string   `json:"description" header:"Description"`
	Priority       string   `json:"priority" header:"Priority"`
	Status         string   `json:"status" header:"Status"`
	SubmitterID    int      `json:"submitter_id" header:"SubmitterID"`
	AssigneeID     int      `json:"assignee_id" header:"AssigneeID"`
	OrganizationID int      `json:"organization_id" header:"OrganizationID"`
	Tags           []string `json:"tags" header:"Tags"`
	HasIncidents   bool     `json:"has_incidents" header:"HasIncidents"`
	DueAt          string   `json:"due_at" header:"DueAt"`
	Via            string   `json:"via" header:"Via"`
}

