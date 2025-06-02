package main

func generator(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- msg
		}
		close(ch)
	}()
	return ch
}

func main() {
	ch := generator("Hello, World!")
	for i := 0; i < 5; i++ {
		println(<-ch)
	}
}
