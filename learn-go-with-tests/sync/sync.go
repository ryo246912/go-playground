package main

import "sync"

type Counter struct {
	// publicにしないため小文字
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// 型の埋め込みは見栄えがいいけど、 public インターフェースの一部になってしまうのであまり良くない
// type Counter struct {
// 	sync.Mutex
// 	value int
// }

// func (c *Counter) Inc() {
// 	c.Lock()
// 	defer c.Unlock()
// 	c.value++
// }

func (c *Counter) Value() int {
	return c.value
}
