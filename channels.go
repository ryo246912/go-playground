package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	// チャネルの作成
	c := make(chan int)
	// スライス内の数値を合算し、2つのgoroutine間で作業を分配します。 両方のgoroutineで計算が完了すると、最終結果が計算されます。
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	// チャネルは、ゴルーチン間でデータをやり取りするための仕組みです。
	// make(chan int) でチャネルを作成し、c <- value で値を送信、<-c で値を受信します。
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
