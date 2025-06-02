package main

import "fmt"

type I int

func (i I) Add(n int) I {
	return i + I(n)
}

func main() {
	var n I = 0

	n = n.Add(1).Add(2)
	fmt.Println(n)

	fmt.Printf("%T\n", n.Add) // Output: main.I.func(int) main.I
	fmt.Printf("%T\n", I.Add) // Output: func(main.I, int) main.I
}
