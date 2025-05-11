package main

func Repeat(character string) string {
	// 文字列を5回繰り返す
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += character
	}
	return repeated
}
