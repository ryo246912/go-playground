package main

import "fmt"

func fibonacci() func() int {
	n, n1 := 0, 1
	return func() int {
		v := n
		n, n1 = n1, n+n1
		return v
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
