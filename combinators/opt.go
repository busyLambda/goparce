package combinators

import (
	. "github.com/busylambda/goparce"
)

func Opt[T any](inner Parser[T]) Parser[T] {
	return func(input *Input) (*T, error) {
		before := input.Eaten()

		result, err := inner(input)

		after := input.Eaten()

		ate := before - after

		if err != nil {
			for i := 0; i < ate; i++ {
				input.UnreadRune()
			}

			return nil, nil
		}
		return result, nil
	}
}
