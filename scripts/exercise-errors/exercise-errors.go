package main

import (
	"fmt"
	"math"
)

// Sqrt は、複素数をサポートしていないので、負の値が与えられたとき、nil以外のエラー値を返す必要があります。
// 新しい型:type ErrNegativeSqrt float64を作成してください。
// そして、 ErrNegativeSqrt(-2).Error() で、 "cannot Sqrt negative number: -2" を返すような:
// func (e ErrNegativeSqrt) Error() stringメソッドを実装し、 error インタフェースを満たすようにします。

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// Error メソッドの中で、 fmt.Sprint(e) を呼び出すことは、無限ループのプログラムになることでしょう。
	// 最初に fmt.Sprint(float64(e)) として e を変換しておくことで、これを避けることができます。
	// 変数eの型はErrNegativeSqrtです。
	// 型ErrNegativeSqrtはerrorインタフェースを満たす関数Error()を実装しています。
	// ErrNegativeSqrt.Error()でfmt.Sprint(e)を呼び出すと、fmt.Sprint()処理内で引数eのe.Error()を呼び出します。
	// return fmt.Sprintf("cannot Sqrt negative number: %v", e)
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
