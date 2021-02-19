package text

import "fmt"

const colourReset = "\033[0m"

// ColourGreen writes green coloured text to stdout
func ColourGreen(msg string) {
	fmt.Println("\033[32m", msg, colourReset)
}

// ColourYellow writes yellow coloured text to stdout
func ColourYellow(msg string) {
	fmt.Println("\u001B[33m", msg, colourReset)
}

// ColourCyan writes cyan coloured text to stdout
func ColourCyan(msg string) {
	fmt.Println("\033[36m", msg, colourReset)
}

// ColourPurple writes purple coloured text to stdout
func ColourPurple(msg string) {
	fmt.Println("\033[35m", msg, colourReset)
}

// ColourWhite writes white coloured text to stdout
func ColourWhite(msg string) {
	fmt.Println("\033[37m", msg, colourReset)
}
