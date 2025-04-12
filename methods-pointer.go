package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// ポインタレシーバでメソッドを宣言できます。
// これはレシーバの型が、ある型 T への構文 *T があることを意味します。 （なお、 T は *int のようなポインタ自身を取ることはできません）
// ポインタレシーバを持つメソッド(ここでは Scale )は、レシーバが指す変数を変更できます。
// レシーバ自身を更新することが多いため、変数レシーバよりもポインタレシーバの方が一般的です。
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// 変数レシーバでは、 Scale メソッドの操作は元の Vertex 変数のコピーを操作します。 （これは関数の引数としての振るまいと同じです）。 つまり main 関数で宣言した Vertex 変数を変更するためには、Scale メソッドはポインタレシーバにする必要があるのです。
func (v Vertex) Scale2(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// 関数
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())

	v2 := Vertex{3, 4}
	v2.Scale2(10)
	fmt.Println(v2.Abs())

	// 関数の場合は正しく、ポインタを渡す必要がある
	var v3 Vertex
	// ScaleFunc(v3, 5)  // Compile error!
	ScaleFunc(&v3, 5) // OK

	// メソッドがポインタレシーバである場合、呼び出し時に、変数、または、ポインタのいずれかのレシーバとして取ることができます:
	var v4 Vertex
	// Goは利便性のため、 v.Scale(5) のステートメントを (&v).Scale(5) として解釈します。
	v4.Scale(5) // OK
	p := &v4
	p.Scale(10) // OK

	// メソッドが変数レシーバである場合、呼び出し時に、変数、または、ポインタのいずれかのレシーバとして取ることができます:
	// この場合、 p.Abs() は (*p).Abs() として解釈されます。
	var v5 Vertex
	fmt.Println(v5.Abs()) // OK
	p2 := &v5
	fmt.Println(p2.Abs()) // OK

	// ポインタレシーバを使う2つの理由があります。
	// ひとつは、メソッドがレシーバが指す先の変数を変更するためです。
	// ふたつに、メソッドの呼び出し毎に変数のコピーを避けるためです。 例えば、レシーバが大きな構造体である場合に効率的です。
	// 例では、 Abs メソッドはレシーバ自身を変更する必要はありませんが、 Scale と Abs は両方とも *Vertex 型のレシーバです。
	// 一般的には、値レシーバ、または、ポインタレシーバのどちらかですべてのメソッドを与え、混在させるべきではありません。 (この理由は数ページ後にわかります)
}
