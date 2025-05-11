package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}
	// handler := http.HandlerFunc(PlayerServer)
	// http.Handlerインターフェースを満たすため、代入できる
	log.Fatal(http.ListenAndServe(":5000", server))
	// log.Fatal(http.ListenAndServe(":5000", handler))
}
