package handler

import (
	"regexp"
)

// Params holds the params needed to process search
type Params struct {
	format string
	entity string
	term   string
	value  string
}

// ParseParams returns the search params required by Searcher
func ParseParams(params string) Params {
	reg, err := regexp.Compile("^(table|json)-(users|tickets|organizations)=(.*):(.*)$")
	if err != nil {
		return Params{}
	}

	match := reg.FindStringSubmatch(params)
	if len(match) == 5 {
		return Params{
			format: match[1],
			entity: match[2],
			term:   match[3],
			value:  match[4],
		}
	}

	return Params{}
}
