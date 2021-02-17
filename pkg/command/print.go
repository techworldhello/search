package command

import "fmt"

// PrintStartMessage writes welcome msg to standard output
func PrintStartMessage() {
	fmt.Println("Welcome to Zendesk Search")
	fmt.Println("Type 'quit' to exit at any time, press 'Enter' to continue")
}

// PrintMenu writes search menu to standard output
func PrintMenu() {
	fmt.Print("Select search options:\n" +
		"* Press 1 to search Zendesk\n" +
		"* Press 2 to view a list of searchable fields\n" +
		"* Type 'quit' to exit\n")
}

// PrintEndMsg writes goodbye msg to standard output
func PrintEndMsg() {
	fmt.Println("Thanks for visiting. Enjoy your day!")
}
