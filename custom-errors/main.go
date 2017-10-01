package main

import (
	"fmt"
)

type customError struct {
	s string
}

type custom_Error struct {
	s string
}

func (e *customError) Error() string {
	return e.s
}

func (e *custom_Error) Error() string {
	return e.s
}

func New(s string) error {
	return &customError{s}
}

func test(index int) (int, error) {
	if index < 10 {
		return index, nil
	}

	return -1, New("Index too high")

}

func main() {
	_, err := test(15)
	if _, ok := err.(*customError); ok {
		fmt.Println(err, "If ok then error type customError")
	}
}
