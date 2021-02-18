package main

import (
	"bufio"
	"fmt"
	"github.com/techworldhello/search/pkg/param"
	"github.com/techworldhello/search/pkg/run"
	"github.com/techworldhello/search/pkg/setup"
	"log"
	"os"
	"strings"

	"github.com/techworldhello/search/pkg/enum"
	"github.com/techworldhello/search/pkg/text"
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

	run := run.NewSearchRepo(data)

	fmt.Println(text.GetStartMsg())
	input := readUserInput()

	var isStart = true
	for input != enum.Quit.String() {
		if isStart {
			isStart = false
			fmt.Println(text.GetMenu())
		}

		input = readUserInput()

		switch input {
		case enum.Search.String():
			fmt.Println(text.GetSearchInstructions())

			params := param.Parse(readUserInput())
			if params == (param.Params{}) {
				fmt.Println("Invalid params, please use the format [entity]=[term]:[value].")
				continue
			}
			fmt.Println(run.ProcessSearch(params))
		case enum.List.String():
			fmt.Println("_____ listing _____")
		case enum.Enter.String():
			fmt.Println(text.GetMenu())
		case enum.Help.String():
			fmt.Println(text.GetMenu())
		case enum.Quit.String():
			fmt.Println(text.GetEndMsg())
			return
		}
	}
	fmt.Println(text.GetEndMsg())
}

func readUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.Replace(text, "\n", "", -1)
}
