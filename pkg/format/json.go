package format

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Json struct {}

func (j Json) Render(result string) {
	var prettyJSON bytes.Buffer

	json.Indent(&prettyJSON, []byte(result), "", "  ")

	fmt.Println(prettyJSON.String())
}
