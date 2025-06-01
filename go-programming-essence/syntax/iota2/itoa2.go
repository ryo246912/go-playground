package main

import "fmt"

const (
	Apple = iota + iota
	Orange
	Banana = iota + 3
)

func main() {
	fmt.Println(Apple)
	fmt.Println(Orange)
	fmt.Println(Banana)
}
