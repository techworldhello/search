package text

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMessages(t *testing.T) {
	expectations := []struct {
		name     string
		function func() string
		expected string
	}{
		{
			name:     "Gets start message",
			function: GetStartMsg,
			expected: ">>> Welcome to Zendesk Search <<<\nType 'quit' to exit at any time, press 'Enter' to continue",
		},
		{
			name:     "Gets menu message",
			function: GetMenu,
			expected: "Select search options:\n  * Press 1 to search Zendesk\n  * Press 2 to view a list of searchable fields\n  * Type 'quit' to exit",
		},
		{
			name:     "Gets end message",
			function: GetEndMsg,
			expected: "Thanks for visiting. Enjoy your day!",
		},
		{
			name:     "Gets search instructions",
			function: GetSearchInstructions,
			expected: "Please search using the following format [entity]=[term]:[value]\n  Eg. users=_id:1\n  Eg. tickets=tags:Alaska\n  Eg. organizations=domain_names:zentix.com",
		},
	}

	for _, expect := range expectations {
		t.Run(expect.name, func(t *testing.T) {
			assert.Equal(t, expect.expected, expect.function())
		})
	}
}
