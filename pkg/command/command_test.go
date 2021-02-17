package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	outOfRange := cmd(5)

	expectations := []struct {
		name     string
		command  string
		expected string
	}{
		{
			name:     "Enum exist",
			command:  Search.String(),
			expected: "1",
		},
		{
			name:     "Enum does not exist",
			command:  outOfRange.String(),
			expected: "unknown command",
		},
	}

	for _, expect := range expectations {
		t.Run(expect.name, func(t *testing.T) {
			assert.Equal(t, expect.expected, expect.command)
		})
	}
}
