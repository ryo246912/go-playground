package main

import (
	"fmt"
	"sync"
	"time"
)

// コンフリクトを避けるために、一度に1つのgoroutineだけが変数にアクセスできるようにしたい場合はどうでしょうか？
// このコンセプトは、排他制御( mutual exclusion )と呼ばれ、このデータ構造を指す一般的な名前は mutex (ミューテックス)です。

// Goの標準ライブラリは、排他制御をsync.Mutexと次の二つのメソッドで提供します:
// Lock
// Unlock
// Inc メソッドにあるように、 Lock と Unlock で囲むことで排他制御で実行するコードを定義できます。
// Value メソッドのように、mutexがUnlockされることを保証するために defer を使うこともできます。

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := {v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
