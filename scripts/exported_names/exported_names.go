package main

import (
	"fmt"
	"math"
)

func main() {
	// Goでは、最初の文字が大文字で始まる名前は、外部のパッケージから参照できるエクスポート（公開）された名前( exported name )です。
	// math.Pi は math パッケージでエクスポートされており、math.piはエラーになります。
	fmt.Println(math.Pi)
}
