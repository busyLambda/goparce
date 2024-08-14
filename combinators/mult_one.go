package combinators

import (
	. "github.com/busylambda/goparce"
)

func MultOne[T any](inner Parser[T]) Parser[[]T] {
	return func(input *Input) (*[]T, error) {
		results := []T{}

		result, err := inner(input)
		if err != nil || result == nil {
			return nil, err
		}

		results = append(results, *result)

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
