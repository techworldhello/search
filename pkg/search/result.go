package search

// Result provides the ability to format and inspect size of result
type Result interface {
	GetSize() int
	Formatter
}

// Formatter provides the ability to format
type Formatter interface {
	Format() [][]string
}
