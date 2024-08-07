package combinators

import (
	"fmt"

	. "github.com/busylambda/combine"
)

func Rune(predicate func(rune) bool) Parser[rune] {
	return func(input *Input) (*rune, error) {
		r, err := input.Peek()
		if err != nil {
			return nil, err
		}

		if !predicate(r) {
			return nil, fmt.Errorf("TODO: Put a custom error here.")
		}

		r, _ = input.PopFront()

		return &r, nil
	}
}
