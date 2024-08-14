package combinators

import (
	"fmt"

	. "github.com/busylambda/goparce"
)

func OneOf[T any](parsers ...Parser[T]) Parser[T] {
	return func(input *Input) (*T, error) {
		for _, parser := range parsers {
			before := input.Eaten()

			result, err := parser(input)
			if err == nil {
				return result, nil
			}

			after := input.Eaten()

			ate := before - after

			for i := 0; i < ate; i++ {
				input.UnreadRune()
			}
		}
		return nil, fmt.Errorf("OneOf: None of the parsers matched")
	}
}
