package text

import "fmt"

// GetStartMsg returns a welcome msg
func GetStartMsg() string {
	return fmt.Sprint("👋 Welcome to Zendesk Search!\n\n" +
		"Type 'quit' to exit at any time, press 'Enter' to continue")
}

// GetMenu returns a search menu
func GetMenu() string {
	return fmt.Sprint("Select search options:\n" +
		"  > Press 1 to search Zendesk\n" +
		"  > Press 2 to view a list of searchable fields\n" +
		"  > Type 'quit' to exit")
}

// GetEndMsg returns a goodbye msg
func GetEndMsg() string {
	return "Thanks for visiting. Enjoy your day!"
}

// GetSearchInstructions returns search instructions
func GetSearchInstructions() string {
	return "🔎 Please search using the following format [view]-[entity]=[field]:[value]\n\n" +
		"  > [view] can either be json or table — it is the way you'd like your data displayed\n" +
		"  > [entity] can either be users, tickets or organizations\n\n" +
		"  Eg. table-users=_id:1\n" +
		"      table-tickets=tags:Alaska\n" +
		"      json-organizations=domain_names:zentix.com"
}

// GetInvalidParamMsg returns an invalid search param warning
func GetInvalidParamMsg() string {
	return "Invalid search params, please use the format [view]-[entity]=[field]:[value]"
}