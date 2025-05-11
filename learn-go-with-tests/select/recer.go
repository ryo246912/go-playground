package main

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string, timeout time.Duration) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

// func Racer(a, b string, timeout time.Duration) (winner string, error error) {
// 	// aDuration := measureResponseTime(a)
// 	// bDuration := measureResponseTime(b)

// 	// if aDuration < bDuration {
// 	// 	return a
// 	// } else {
// 	// 	return b
// 	// }

// 	// selectでできることは、multiple チャネルで待機することです。
// 	select {
// 	case <-ping(a):
// 		return a, nil
// 	case <-ping(b):
// 		return b, nil
// 	// time.Afterは、selectを使用する場合に非常に便利な関数です。
// 	// 今回のケースでは発生しませんでしたが、リッスンしているチャネルが値を返さない場合、永久にブロックするコードを書く可能性があります。
// 	// time.Afterは、chan（ pingのように）を返し、指定した時間が経過すると信号を送ります。
// 	case <-time.After(timeout):
// 		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
// 	}
// }

// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	return time.Since(start)
// }

// chan struct{}を作成して返す関数pingを定義
// - struct{}はメモリの観点から利用できる最小のデータ型
func ping(url string) chan struct{} {
	// varを使用すると、変数は型の「ゼロ」値で初期化されます。したがって、stringの場合は""、intの場合は0になります。
	// チャネルの場合、ゼロ値はnilであり、<-で送信しようとすると、nilチャネルに送信できないため、永久にブロックされます。
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}
