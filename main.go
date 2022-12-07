package go_math

import (
	"fmt"
	"errors"
)

func Add(a, b int) int {
	fmt.Println("Add on my gitee repo");
	return a+b
}

func Sub(a, b int) int {
	fmt.Println("Log from my github code-space. Sub()call ...")
	return a-b
}

func Mul(a, b int) int {
	fmt.Println("Log from assist‘s github code-space. Mul()call ...")
	return a*b
}

func Div(a, b int) (int, error) {
	if 0 == b {
		return 0, errors.New("Not allow divided by zero")
	}
	return a/b, nil
}

func Power(a, b int) int {
	if 1 == b {
		return  a
	}

	if 0 == b {
		return 1
	}

	if 0 == b&1  {
		half := Power(a, b/2)
		return half*half
	}
	return Power(a, b-1) * a
}
