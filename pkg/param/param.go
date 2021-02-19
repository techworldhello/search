package param

import (
	"regexp"
)

// Params holds the params needed to process search
type Params struct {
	Format string
	Entity string
	Term   string
	Value  string
}

// Parse returns the search params required by Searcher
func Parse(params string) Params {
	reg, err := regexp.Compile("^(table|json)-(users|tickets|organizations)=(.*):(.*)$")
	if err != nil {
		return Params{}
	}

	match := reg.FindStringSubmatch(params)
	if len(match) == 5 {
		return Params{
			Format: match[1],
			Entity: match[2],
			Term:   match[3],
			Value:  match[4],
		}
	}

	return Params{}
}
