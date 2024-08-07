package combinators

import (
	"fmt"

	. "github.com/busylambda/combine"
)

func String(str string) Parser[string] {
	return func(input *Input) (*string, error) {
		n := len(str)

		p, err := input.PeekN(n)
		if err != nil {
			return nil, err
		}

		if p != str {
			return nil, fmt.Errorf("Error here later.")
		}

		result, _ := input.PopFrontN(n)

		return &result, nil
	}
}
