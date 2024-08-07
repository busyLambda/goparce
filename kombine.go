package kombine

type Parser[T any] func(input *Input) (*T, error)
