package combinators

import (
	. "github.com/busylambda/combine"
)

func MultZero[T any](inner Parser[T]) Parser[[]T] {
	return func(input *Input) (*[]T, error) {
		results := []T{}

		for {
			result, err := inner(input)
			if err != nil {
				break
			}

			results = append(results, *result)
		}

		return &results, nil
	}
}
