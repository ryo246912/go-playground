package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	// interface(インタフェース)型は、メソッドのシグニチャの集まりで定義します。
	// そのメソッドの集まりを実装した値を、interface型の変数へ持たせることができます。
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// Abs メソッドが、 Vertex ではなく *Vertex の定義であり、 Vertex が Abser インタフェースを実装していないということになるためエラーとなります。
	a = v

	fmt.Println(a.Abs())
}
