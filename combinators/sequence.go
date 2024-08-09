package combinators

import (
	. "github.com/busylambda/goparce"
)

func Sequence[T any](parsers ...Parser[T]) Parser[[]T] {
	return func(input *Input) (*[]T, error) {
		results := []T{}

		for _, parser := range parsers {
			result, err := parser(input)
			if err != nil {
				return nil, err
			}

			results = append(results, *result)
		}

		return &results, nil
	}
}
