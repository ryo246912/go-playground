package main

import "fmt"

// 定数は、 Helloが呼び出されるたびに"Hello、"文字列インスタンスを作成する手間を省く
const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	// return fmt.Sprintf("%s%s", englishHelloPrefix, name)
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Chris"))
}
