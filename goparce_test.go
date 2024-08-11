package goparce_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"
	"unicode"

	goparce "github.com/busylambda/goparce"
	. "github.com/busylambda/goparce/combinators"
)

func TestString(t *testing.T) {
	input := goparce.NewInput("matchme")
	str, err := String("matchme")(input)
	if err != nil {
		log.Fatalf(err.Error())
	}

	if *str != "matchme" {
		log.Fatalf(`Result did not match "matchme"`)
	}
}

func TestDelimited(t *testing.T) {
	input := goparce.NewInput("[settings]")

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
	input := goparce.NewInput("array_number_1")

	isIdentCont := func(r rune) bool {
		return unicode.IsLetter(r) || r == '_' || unicode.IsNumber(r)
	}

	isIdentStart := func(r rune) bool {
		return unicode.IsLetter(r) || r == '_'
	}

	type Identifier struct {
		literal string
	}

	identifier := func() goparce.Parser[Identifier] {
		return func(input *goparce.Input) (*Identifier, error) {
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

func TestStripWhitespace(t *testing.T) {
	input := goparce.NewInput("  hello   ")
	_, err := StripWhitespace(String("hello"))(input)
	if err != nil {
		log.Println(err.Error())
		log.Fatalf(`Result did not match "   hello   "`)
	}
}

func Test1millionStrings(t *testing.T) {
	data, err := ioutil.ReadFile("output.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	content := string(data)

	input := goparce.NewInput(content)

	_, err = MultOne(StripWhitespace(String("hello1")))(input)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func TestSeparatedList(t *testing.T) {
	input := goparce.NewInput("item, item, item, item, item, item")

	parser := SepList(StripWhitespace(String("item")), String(","))

	_, err := parser(input)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func TestOpt(t *testing.T) {
	input := goparce.NewInput("function")

	parser := Sequence((Opt(String("public"))), StripWhitespace(String("function")))

	_, err := parser(input)
	if err != nil {
		panic(err.Error())
	}
}
