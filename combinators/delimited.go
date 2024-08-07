package combinators

import (
	. "github.com/busylambda/combine"
)

func Delimited[T any, D any](left Parser[D], inner Parser[T], right Parser[D]) Parser[T] {
	return func(input *Input) (*T, error) {
		_, err := left(input)
		if err != nil {
			return nil, err
		}

		result, err := inner(input)
		if err != nil {
			return nil, err
		}

		_, err = right(input)
		if err != nil {
			return nil, err
		}

		return result, nil
	}
}
