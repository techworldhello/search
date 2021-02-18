package enum

type cmd int

const (
	Search cmd = iota
	List
	Enter
	Help
	Quit
)

// String returns the string type of valid command enums
func (c cmd) String() string {
	var commands = [...]string {
		"1",
		"2",
		"",
		"help",
		"quit",
	}
	if c < Search || c > Quit {
		return "unknown command"
	}
	return commands[c]
}
