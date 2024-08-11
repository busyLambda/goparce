package combinators

import (
	. "github.com/busylambda/goparce"
)

func SepList[T any, S any](item Parser[T], sep Parser[S]) Parser[[]T] {
	return func(input *Input) (*[]T, error) {
		results := []T{}

		initial_item, err := item(input)
		if err != nil {
			return nil, err
		}

		results = append(results, *initial_item)

		for {
			_, err := sep(input)
			if err != nil {
				break
			}

			new_item, err := item(input)
			if err != nil {
				return nil, err
			}

			results = append(results, *new_item)
		}

		return &results, nil
	}
}
