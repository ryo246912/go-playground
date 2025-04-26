package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		// select ステートメントは、goroutineを複数の通信操作で待たせます。
		// select は、複数ある case のいずれかが準備できるようになるまでブロックし、準備ができた case を実行します。
		// もし、複数の case の準備ができている場合、 case はランダムに選択されます。
		select {
		// xの値をcチャネルに送信
		case c <- x:
			x, y = y, x+y
		// quitチャネルを受信したら、後続の処理を実行
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
