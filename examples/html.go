package main

import (
	"fmt"

	"github.com/busylambda/goparce"
	C "github.com/busylambda/goparce/combinators"
)

type HtmlNode struct {
	TagName    string
	Attributes []Attribute
	Children   []HtmlNode
}

type Attribute struct {
	Key   string
	Value string
}

func main() {
	input := goparce.NewInput(`<h1 class="text-2xl"></h1>`)

	node, err := htmlNode()(input)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(node)
}

func htmlNode() goparce.Parser[HtmlNode] {
	return func(input *goparce.Input) (*HtmlNode, error) {
		_, err := C.StripWhitespace(C.String("<"))(input)
		if err != nil {
			return nil, err
		}

		tagName, err := C.StripWhitespace(identifier())(input)
		if err != nil {
			return nil, err
		}

		attributes, err := C.MultZero(attribute())(input)
		if err != nil {
			return nil, err
		}

		_, err = C.StripWhitespace(C.String(">"))(input)
		if err != nil {
			return nil, err
		}

		closingTag, err := C.StripWhitespace(C.Delimited(C.String("</"), identifier(), C.String(">")))(input)
		if err != nil {
			return nil, err
		}

		if *closingTag != *tagName {
			return nil, fmt.Errorf("opening and closing tag did not match god %s but wanted %s", *closingTag, *tagName)
		}

		return &HtmlNode{
			TagName:    *tagName,
			Children:   []HtmlNode{},
			Attributes: *attributes,
		}, nil
	}
}

func attribute() goparce.Parser[Attribute] {
	return func(input *goparce.Input) (*Attribute, error) {
		key, err := C.StripWhitespace(identifier())(input)
		if err != nil {
			return nil, err
		}

		_, err = C.StripWhitespace(C.String("="))(input)
		if err != nil {
			return nil, err
		}

		text := C.MultZero(C.Rune(func(r rune) bool { return r != '"' }))
		value, err := C.StripWhitespace(C.Delimited(C.String("\""), text, C.String("\"")))(input)
		if err != nil {
			return nil, err
		}

		return &Attribute{
			Key:   *key,
			Value: string(*value),
		}, nil
	}
}

func identifier() goparce.Parser[string] {
	return func(input *goparce.Input) (*string, error) {
		result := ""

		first_rune, err := C.Rune(isIdentifierStart)(input)
		if err != nil {
			return nil, err
		}

		result += string(*first_rune)

		rest, err := C.MultZero(C.Rune(isIdentifier))(input)
		if err != nil {
			return nil, err
		}

		result += string(*rest)

		return &result, nil
	}
}

func isIdentifier(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r == '_') || (r >= '0' && r <= '9')
}
func isIdentifierStart(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}
