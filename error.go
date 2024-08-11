package goparce

import "fmt"

type MismatchErr struct {
  expected string
  found    string
}

func (self *MismatchErr) Error() string {
  return fmt.Sprintf("Expected: `%s` but found `%s`", self.expected, self.found)
}

func NewMismatchErr(expected string, found string) error {
  return &MismatchErr{
    expected,
    found,
  }
}
