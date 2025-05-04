package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// 受け手は、受信の式に2つ目のパラメータを割り当てることで、そのチャネルがcloseされているかどうかを確認できます:
	// v, ok := <-ch
	// 受信する値がない、かつ、チャネルが閉じているなら、 ok の変数は、 false になります。

	// 送り手は、これ以上の送信する値がないことを示すため、チャネルを close できます。
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)

	// 受信側は、チャネルが close されるまで、受信を続けます
	for i := range c {
		fmt.Println(i)
	}
}
