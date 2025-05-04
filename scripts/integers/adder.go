package main

// ２つの数値を足した値を返します
func Add(x, y int) int {
	return x + y
}

func main() {
	// 2 + 3 = 5
	sum := Add(2, 3)
	println(sum)
}
