package combinators

import (
	"io"
	"unicode"

	. "github.com/busylambda/goparce"
)

func isWhitespace(r rune) bool {
	return unicode.IsSpace(r)
}

func StripWhitespace[T any](innerParser Parser[T]) Parser[T] {
	return func(input *Input) (*T, error) {
		for {
			r, err := input.Peek()
			if err != nil {
				return nil, err
			}

			if isWhitespace(r) {
				input.PopFront()
			} else {
				break
			}
		}

		res, err := innerParser(input)
		if err != nil {
			return nil, err
		}

		for {
			r, err := input.Peek()
			if err != nil {
				if err == io.EOF {
					break
				}

				return nil, err
			}

			if isWhitespace(r) {
				input.PopFront()
			} else {
				break
			}
		}
		return res, nil
	}
}
