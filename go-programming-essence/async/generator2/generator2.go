package main

import (
	"fmt"
	"time"
)

func fanIn(ch1, ch2 <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			select {
			case msg := <-ch1:
				ch <- msg
			case msg := <-ch2:
				ch <- msg
			}
		}
	}()
	return ch
}

func generator(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- msg
			time.Sleep(time.Second)
		}
		close(ch)
	}()
	return ch
}

func main() {
	ch := fanIn(generator("Hello, World!"), generator("Hello, Go!"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}
