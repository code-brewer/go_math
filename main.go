package go_math

import (
	"errors"
)

func Add(a, b int) int {
	return a+b
}

func Sub(a, b int) int {
	return a-b
}

func Mul(a, b int) int {
	return a*b
}

func Div(a, b int) (int, error) {
	if 0 == b {
		return 0, errors.New("Not allow divided by zero")
	}
	return a/b, nil
}