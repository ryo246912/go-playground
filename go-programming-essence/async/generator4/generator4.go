package main

import (
	"fmt"
	"sync"
	"time"
)

func downloadJSON(u string) {
	fmt.Println(u)
	time.Sleep(time.Second * 6)
}

func main() {
	before := time.Now()

	limit := make(chan struct{}, 20) // 同時に3つのゴルーチンを実行可能
	var wg sync.WaitGroup
	for i := 1; i <= 100; i++ {
		wg.Add(1) // ゴルーチンの数をカウントアップ
		go func(i int) {
			limit <- struct{}{} // チャネルに空の構造体を送信して、ゴルーチンの数を制限
			defer wg.Done()     // ゴルーチンが終了したらカウントダウン
			downloadJSON(fmt.Sprintf("https://example.com/%d.json", i))
			defer func() { <-limit }() // ゴルーチンが終了したらチャネルから空の構造体を取り出す
		}(i)
	}
	wg.Wait()

	fmt.Printf("%v\n", time.Since(before))
}
