package main

import (
	"bufio"
	"fmt"
	"github.com/techworldhello/search/pkg/setup"
	"log"
	"os"
	"strings"

	"github.com/techworldhello/search/pkg/command"
)

func main() {
	files := setup.NewFilePaths(
		"./pkg/data/users.json",
		"./pkg/data/tickets.json",
		"./pkg/data/organizations.json",
	)

	data, err := files.PrepareJSONData()
	if err != nil {
		log.Fatal("Error preparing files, please check them and try again.")
	}

	fmt.Print(data)

	command.PrintStartMessage()
	input := readUserInput()

	var isStart = true

	for input != command.Quit.String() {
		if isStart {
			isStart = false
			command.PrintMenu()
		}

		input = readUserInput()

		switch input {
		case command.Search.String():
			fmt.Println("Enter Search Term")
		case command.List.String():
			fmt.Println("_____ listing _____")
		case command.Enter.String():
			command.PrintMenu()
		case command.Help.String():
			command.PrintMenu()
		case command.Quit.String():
			command.PrintEndMsg()
			return
		}
	}
	command.PrintEndMsg()
}

func readUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.Replace(text, "\n", "", -1)
}
