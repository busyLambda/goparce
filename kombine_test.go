package kombine_test

import (
	"log"
	"testing"
	"unicode"

	kombine "github.com/busylambda/combine"
	. "github.com/busylambda/combine/combinators"
)

func TestString(t *testing.T) {
	input := kombine.NewInput("matchme")
	str, err := String("matchme")(input)
	if err != nil {
		log.Fatalf(err.Error())
	}

	if *str != "matchme" {
		log.Fatalf(`Result did not match "matchme"`)
	}
}

func TestDelimited(t *testing.T) {
	input := kombine.NewInput("[settings]")

	left := String("[")
	inner := String("settings")
	right := String("]")

	str, err := Delimited(left, inner, right)(input)
	if err != nil {
		log.Fatalf(err.Error())
	}

	if *str != "settings" {
		log.Fatalf(`Result did not match "settings"`)
	}
}

func TestParseIdentifier(t *testing.T) {
	input := kombine.NewInput("array_number_1")

	isIdentCont := func(r rune) bool {
		return unicode.IsLetter(r) || r == '_' || unicode.IsNumber(r)
	}

	isIdentStart := func(r rune) bool {
		return unicode.IsLetter(r) || r == '_'
	}

	type Identifier struct {
		literal string
	}

	identifier := func() kombine.Parser[Identifier] {
		return func(input *kombine.Input) (*Identifier, error) {
			start, err := Rune(isIdentStart)(input)
			if err != nil {
				return nil, err
			}

			rest, err := MultOne(Rune(isIdentCont))(input)
			if err != nil {
				return nil, err
			}

			literal := string(*start) + string(*rest)

			return &Identifier{literal}, nil
		}
	}

	ident, err := identifier()(input)
	if err != nil {
		log.Fatalf(err.Error())
	}

	if ident.literal != "array_number_1" {
		log.Fatalf(`Result did not match "i_am_an_identifier"`)
	}
}

/*
func TestFuncArgs(t *testing.T) {
	input := kombine.NewInput("(a: int, b: int)")

	left := String("[")
	inner := String("settings")
	right := String("]")

	str, err := Delimited(left, inner, right)(input)
	if err != nil {
		log.Fatalf(err.Error())
	}

	if *str != "settings" {
		log.Fatalf(`Result did not match "settings"`)
	}
}
*/
