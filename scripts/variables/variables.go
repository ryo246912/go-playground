package main

import "fmt"

// 関数の引数リストと同様に、複数の変数の最後に型を書くことで、変数のリストを宣言できます。
var c, python, java bool

// var 宣言では、変数毎に初期化子( initializer )を与えることができます。
// 初期化子が与えられている場合、型を省略できます。その変数は初期化子が持つ型になります。
var c1, python1, java1 = true, false, "no!"

func main() {
	var i int
	// var 宣言の代わりに、短い := の代入文を使い、暗黙的な型宣言ができます。
	// なお、関数の外では、キーワードではじまる宣言( var, func, など)が必要で、 := での暗黙的な宣言は利用できません。
	k := 3

	fmt.Println(i, c, python, java)
	fmt.Println(k, c1, python1, java1)
}
