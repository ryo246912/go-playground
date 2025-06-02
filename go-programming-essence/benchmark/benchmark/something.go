package something

import "fmt"

func fanIn(ch1, ch2 <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		for {
			select {
			case s := <-ch1:
				out <- s
			case s := <-ch2:
				out <- s
			default:
				fmt.Println("No messages received, waiting...")
			}
		}
	}()
	return out
}

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

func doSomething() {
	ch := fanIn(generator("Hello, World!"), generator("Hello, Go!"))
	for i := 0; i < 10; i++ {
		println(<-ch)
	}
}
