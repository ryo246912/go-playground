package main

import (
	"log"
	"net/http"
)

func main() {
	// Handlerインターフェースがサーバーを作るために実装する必要があるものであることを探りました。
	// 通常は、structを作成してそれを行い、独自のServeHTTPメソッドを実装してインターフェースを実装します。
	// ただし、構造体のユースケースはデータを保持するためのものですが、currently には状態がないため、データを作成するのは適切ではありません。
	// HandlerFuncを使用すると、これを回避できます。

	// HandlerFuncタイプは、通常の関数をHTTPハンドラーとして使用できるようにするアダプターです。
	// fが適切なシグネチャを持つ関数である場合、HandlerFunc(f) はfを呼び出すハンドラーです。
	handler := http.HandlerFunc(PlayerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
