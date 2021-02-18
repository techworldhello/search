package search

// Result represents the search result
type Result interface {
	GetSize() int
	Formatter
}

type Formatter interface {
	Format() [][]string
}
