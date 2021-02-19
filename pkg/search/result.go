package search

// Result provides the ability to format and inspect size of result
type Result interface {
	GetSize() int
	Formatter
	Stringer
}

// Formatter provides the ability to format
type Formatter interface {
	Format() [][]string
}

// Stringer provides the ability to convert to string
type Stringer interface {
	ToString() (string, error)
}
