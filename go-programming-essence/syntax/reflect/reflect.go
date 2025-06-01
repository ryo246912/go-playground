package main

import (
	"fmt"
	"reflect"
)

func printDetail(v any) {
	rt := reflect.TypeOf(v)
	switch rt.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Println("int:", v)
	case reflect.String:
		fmt.Println("string:", v)
	default:
		fmt.Println("unknown type:", rt.Kind())
	}
}

func main() {
	type V int
	var v V = 10
	printDetail(v)

	s := struct{}{}
	printDetail(s)
}
