package main

import (
	"bufio"
	"fmt"
	"github.com/techworldhello/search/pkg/format"
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
			text.ColourCyan(text.GetMenu())
		}

		input = readUserInput()

		switch input {
		case enum.Search.String():
			text.ColourGreen(text.GetSearchInstructions())

			params := param.Parse(readUserInput())
			if params == (param.Params{}) {
				fmt.Println("Invalid params, please use the format [entity]=[term]:[value].")
				continue
			}
			text.ColourWhite(run.ProcessSearch(params))
		case enum.List.String():
			f := format.Fields{}
			text.ColourYellow(f.List())
		case enum.Enter.String():
			text.ColourCyan(text.GetMenu())
		case enum.Help.String():
			text.ColourCyan(text.GetMenu())
		case enum.Quit.String():
			text.ColourPurple(text.GetEndMsg())
			return
		}
	}
	text.ColourPurple(text.GetEndMsg())
}

func readUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.Replace(text, "\n", "", -1)
}
