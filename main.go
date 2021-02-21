package main

import (
	"bufio"
	"github.com/techworldhello/search/pkg/format"
	"github.com/techworldhello/search/pkg/handler"
	"github.com/techworldhello/search/pkg/setup"
	"log"
	"os"
	"strings"

	"github.com/techworldhello/search/pkg/enum"
	"github.com/techworldhello/search/pkg/text"
)

func main() {
	files := setup.NewFilePaths(
		os.Getenv("USERS_FILE"),
		os.Getenv("TICKETS_FILE"),
		os.Getenv("ORGANIZATIONS_FILE"),
	)

	data, err := files.PrepareJSONData()
	if err != nil {
		log.Fatal("Error preparing files, please check them and try again.")
	}

	repo := handler.NewSearchRepo(data)
	fields := format.NewSearchFields()

	text.ColourPurple(text.GetStartMsg())
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
			params := handler.ParseParams(readUserInput())
			if params == (handler.Params{}) {
				text.ColourWhite(text.GetInvalidParamMsg())
				continue
			}
			text.ColourWhite(repo.ProcessSearch(params))
		case enum.List.String():
			text.ColourYellow(fields.List())
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
