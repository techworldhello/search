package text

import "fmt"

// GetStartMsg returns a welcome msg
func GetStartMsg() string {
	return fmt.Sprint(">>> Welcome to Zendesk Search <<<\n" +
		"Type 'quit' to exit at any time, press 'Enter' to continue")
}

// GetMenu returns a search menu
func GetMenu() string {
	return fmt.Sprint("Select search options:\n" +
		"  * Press 1 to search Zendesk\n" +
		"  * Press 2 to view a list of searchable fields\n" +
		"  * Type 'quit' to exit")
}

// GetEndMsg returns a goodbye msg
func GetEndMsg() string {
	return "Thanks for visiting. Enjoy your day!"
}

// GetSearchInstructions returns search instructions
func GetSearchInstructions() string {
	return "Please search using the following format [entity]=[term]:[value]\n" +
		"  Eg. users=_id:1\n" +
		"  Eg. tickets=tags:Alaska\n" +
		"  Eg. organizations=domain_names:zentix.com"
}
