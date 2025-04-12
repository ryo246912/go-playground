package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

// ポインタレシーバーで定義されたメソッドは、構造体のポインタ型（*MyError）に対してのみ呼び出せます。
// そのため、run 関数で &MyError{...} を返す必要。
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

// error 型は fmt.Stringer に似た組み込みのインタフェースです:
//
//	type error interface {
//		Error() string
//	}
//
// error はインターフェース型であり、Error() string メソッドを持つ型が実装とみなされます。
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
