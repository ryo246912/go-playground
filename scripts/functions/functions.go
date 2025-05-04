package main

import "fmt"

// 型は後ろに書く
func add(x int, y int) int {
	return x + y
}

// 関数の２つ以上の引数が同じ型である場合には、最後の型を残して省略して記述できます。
// x int, y int→x, y int
func add2(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

// Goでの戻り値となる変数に名前をつける( named return value )ことができます。
// 戻り値に名前をつけると、関数の最初で定義した変数名として扱われます。
// この戻り値の名前は、戻り値の意味を示す名前とすることで、関数のドキュメントとして表現するようにしましょう。

// 名前をつけた戻り値の変数を使うと、 return ステートメントに何も書かずに戻すことができます。
// これを "naked" return と呼びます。
// 例のコードのように、naked returnステートメントは、短い関数でのみ利用すべきです。長い関数で使うと読みやすさ( readability )に悪影響があります。
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(add2(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))
}
