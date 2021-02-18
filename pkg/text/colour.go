package text

import "fmt"

// ColourGreen writes green coloured text to stdout
func ColourGreen(msg string) {
	fmt.Println("\033[32m", msg)
}

// ColourYellow writes yellow coloured text to stdout
func ColourYellow(msg string) {
	fmt.Println("\u001B[33m", msg)
}

// ColourCyan writes cyan coloured text to stdout
func ColourCyan(msg string) {
	fmt.Println("\033[36m", msg)
}

// ColourPurple writes purple coloured text to stdout
func ColourPurple(msg string) {
	fmt.Println("\033[35m", msg)
}

// ColourWhite writes white coloured text to stdout
func ColourWhite(msg string) {
	fmt.Println("\033[37m", msg)
}
