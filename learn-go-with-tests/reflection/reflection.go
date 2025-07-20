package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	// 指定された変数のValueを返す関数ValueOf
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		if field.Kind() == reflect.String {
			fn(field.String())
		}
	}
}
