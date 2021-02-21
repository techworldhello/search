package search

import (
	"encoding/json"
	"github.com/techworldhello/search/pkg/schema"
	"strconv"
	"strings"
)

// UserResult holds the user search result
type UserResult struct {
	data []schema.User
	size int
}

// GetSize returns the size of user result
func (u UserResult) GetSize() int {
	return u.size
}

// Format formats all user results
func (u UserResult) Format() (all [][]string) {
	for _, v := range u.data {
		var row []string

		row = append(row,
			strconv.Itoa(v.ID),
			v.Name,
			v.Alias,
			strconv.FormatBool(v.Active),
			strconv.FormatBool(v.Verified),
			strconv.FormatBool(v.Shared),
			v.Locale,
			v.Timezone,
			v.Email,
			v.Phone,
			v.Signature,
			strconv.Itoa(v.OrganizationID),
			strings.Join(v.Tags, ",\n"),
			strconv.FormatBool(v.Suspended),
			v.Role,
		)
		all = append(all, row)
	}
	return all
}

func (u UserResult) ToString() (string, error) {
	bytes, err := json.Marshal(u.data)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
