package combinators

import (
	. "github.com/busylambda/goparce"
)

func Opt[T any](inner Parser[T]) Parser[T] {
	return func(input *Input) (*T, error) {
		result, err := inner(input)
		if err != nil {
			return nil, nil
		}

		return result, nil
	}
}
