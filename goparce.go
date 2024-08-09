package goparce

// A function that may be passed around as a parser.
type Parser[T any] func(input *Input) (*T, error)
