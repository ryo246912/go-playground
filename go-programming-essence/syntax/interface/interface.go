package main

import "fmt"

// (*foo)(nil)は「*foo型のnilポインタ」を作る
// 型変換している型(値)
var _ I = (*foo)(nil)

type I interface {
	doSomething()
	doOtherThing()
}

type foo struct{}

func (f *foo) doSomething() {
	fmt.Println("Doing something")
}

func main() {
	f := &foo{}
	f.doSomething()
}
