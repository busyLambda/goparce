package combinators

import kombine "github.com/busylambda/combine"

func Opt[T any](inner kombine.Parser[T]) kombine.Parser[T] {
	return func(input *kombine.Input) (*T, error) {
		result, err := inner(input)
		if err != nil {
			return nil, nil
		}

		return result, nil
	}
}
