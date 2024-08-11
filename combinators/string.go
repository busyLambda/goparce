package combinators

import (
	. "github.com/busylambda/goparce"
)

func String(str string) Parser[string] {
	return func(input *Input) (*string, error) {
		n := len(str)

		p, err := input.PeekN(n)
		if err != nil {
			return nil, err
		}

		if p != str {
      err = NewMismatchErr(str, p)

			return nil, err
		}

		result, _ := input.PopFrontN(n)

		return &result, nil
	}
}
