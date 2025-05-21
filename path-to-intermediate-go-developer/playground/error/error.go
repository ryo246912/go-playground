package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	_, err0 := strconv.Atoi("a") // "a"は数値に直せないので、エラーが発生する
	fmt.Printf("err0: [%T] %v\n", err0, err0)

	err1 := errors.Unwrap(err0)
	fmt.Printf("err1: [%T] %v\n", err1, err1)
}
