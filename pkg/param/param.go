package param

import "regexp"

// Params holds the params needed to process search
type Params struct {
	Entity string
	Term   string
	Value  string
}

// Parse returns 3 search params required by Searcher
func Parse(params string) Params {
	reg, err := regexp.Compile("(users|tickets|organizations)=(.*):(.*)")
	if err != nil {
		return Params{}
	}

	match := reg.FindStringSubmatch(params)
	if len(match) == 4 {
		return Params{
			Entity: match[1],
			Term:   match[2],
			Value:  match[3],
		}
	}

	return Params{}
}
