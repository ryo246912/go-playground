package main

import "fmt"

func do(i interface{}) {
	// 型switch はいくつかの型アサーションを直列に使用できる構造です。
	// 型switchは通常のswitch文と似ていますが、型switchのcaseは型(値ではない)を指定し、それらの値は指定されたインターフェースの値が保持する値の型と比較されます。

	// switch v := i.(type) {
	// case T:
	//     // here v has type T
	// case S:
	//     // here v has type S
	// default:
	//     // no match; here v has the same type as i
	// }

	// 型switchの宣言は、型アサーション i.(T) と同じ構文を持ちますが、特定の型 T はキーワード type に置き換えられます。
	// このswitch文は、インターフェースの値 i が 型 T または S の値を保持するかどうかをテストします。
	// T および S の各caseにおいて、変数 v はそれぞれ 型 T または S であり、 i によって保持される値を保持します。
	// defaultの場合(一致するものがない場合)、変数 v は同じインターフェース型で値は i となります。

	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
