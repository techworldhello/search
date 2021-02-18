package schema

// User represents the shape of user data
type User struct {
	ID             int      `json:"_id" header:"ID"`
	URL            string   `json:"url"`
	ExternalID     string   `json:"external_id"`
	Name           string   `json:"name" header:"Name"`
	Alias          string   `json:"alias" header:"Alias"`
	CreatedAt      string   `json:"created_at"`
	Active         bool     `json:"active" header:"Active"`
	Verified       bool     `json:"verified" header:"Verified"`
	Shared         bool     `json:"shared" header:"Shared"`
	Locale         string   `json:"locale" header:"Locale"`
	Timezone       string   `json:"timezone" header:"Timezone"`
	LastLoginAt    string   `json:"last_login_at"`
	Email          string   `json:"email" header:"Email"`
	Phone          string   `json:"phone" header:"Phone"`
	Signature      string   `json:"signature" header:"Signature"`
	OrganizationID int      `json:"organization_id" header:"OrganizationID"`
	Tags           []string `json:"tags" header:"Tags"`
	Suspended      bool     `json:"suspended" header:"Suspended"`
	Role           string   `json:"role" header:"Role"`
}
