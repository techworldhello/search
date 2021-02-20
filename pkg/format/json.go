package format

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Json represents the type to render
type Json struct {}

// Render writes the search result to stdout in JSON format
func (j Json) Render(result string) {
	var prettyJSON bytes.Buffer

	json.Indent(&prettyJSON, []byte(result), "", "  ")

	fmt.Println(prettyJSON.String())
}
