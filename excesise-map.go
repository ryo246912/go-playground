package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	var result = make(map[string]int)
	a := strings.Split(s, " ")

	for _, value := range a {
		result[value] += 1
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
