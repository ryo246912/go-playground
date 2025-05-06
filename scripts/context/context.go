package main

import (
	"context"
	"fmt"
	"net/http"
)

// type Store interface {
// 	Fetch() string
// 	Cancel()
// }

// func Server(store Store) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		ctx := r.Context()

// 		data := make(chan string, 1)

// 		go func() {
// 			data <- store.Fetch()
// 		}()

// 		select {
// 		// cancel前にdataを取得した場合はこちらを実行
// 		case d := <-data:
// 			fmt.Fprint(w, d)
// 		// contextにはメソッドDone()があり、コンテキストが「完了」または「キャンセル」されたときに信号を送信するチャネルを返す
// 		case <-ctx.Done():
// 			store.Cancel()
// 		}
// 	}
// }

type Store interface {
	// contextの主なポイントの1つは、キャンセルを提供する一貫した方法
	// コンテキストがキャンセルされると、そのコンテキストから派生したすべてのコンテキストもキャンセルされます。
	// コンテキストがダウンストリームのStoreに送信されることと、キャンセルされたときにStoreから発生するエラーを処理すること
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return // todo: log error however you like
		}

		fmt.Fprint(w, data)
	}
}
