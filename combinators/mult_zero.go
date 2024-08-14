package combinators

import (
	. "github.com/busylambda/goparce"
)

func MultZero[T any](inner Parser[T]) Parser[[]T] {
	return func(input *Input) (*[]T, error) {
		results := []T{}

		for {
			result, err := inner(input)
			if err != nil || result == nil {
				break
			}

			results = append(results, *result)
		}

		return &results, nil
	}
}
