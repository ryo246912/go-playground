package main

import "fmt"

func main() {
	// チャネルは、 バッファ ( buffer )として使えます。 バッファを持つチャネルを初期化するには、 make の２つ目の引数にバッファの長さを与えます:
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// バッファが詰まった時は、チャネルへの送信をブロックします。
	// ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// バッファが空の時には、チャネルの受信をブロックします。
	// fmt.Println(<-ch)
}
