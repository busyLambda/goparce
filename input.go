package goparce

import (
	"bufio"
	"io"
	"strings"
)

type Input struct {
	reader bufio.Reader
	length int
	eaten  int
}

func NewInput(input string) *Input {
	return &Input{
		reader: *bufio.NewReader(strings.NewReader(input)),
	}
}

func (self *Input) PopFrontN(n int) (string, error) {
	buf := make([]byte, n)
	_, err := io.ReadFull(&self.reader, buf)
	if err == nil {
		self.eaten += n
	}
	return string(buf), err
}

func (self *Input) PopFront() (rune, error) {
	buf := make([]byte, 1)
	_, err := io.ReadFull(&self.reader, buf)
	if err == nil {
		self.eaten += 1
	}
	return rune(buf[0]), err
}

func (self *Input) PeekN(n int) (string, error) {
	b, err := self.reader.Peek(n)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (self *Input) Peek() (rune, error) {
	b, err := self.reader.Peek(1)
	if err != nil {
		return '0', err
	}

	return rune(b[0]), nil
}

func (self *Input) UnreadRune() error {
	err := self.reader.UnreadRune()
	if err == nil {
		self.eaten -= 1
	}
	return err
}

func (self *Input) Eaten() int {
	return self.eaten
}

func (self *Input) Length() int {
	return self.length
}
