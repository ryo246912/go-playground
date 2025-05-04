package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// 定数は、 Helloが呼び出されるたびに"Hello、"文字列インスタンスを作成する手間を省く
const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Greet(writer io.Writer, name string) {
	// fmt.Fprintf(writer, fmt.Sprintf("%s%s", englishHelloPrefix, name))
	// バッファに送信
	fmt.Fprintf(writer, "%s%s", englishHelloPrefix, name)
}

// http.ResponseWriterと、リクエストの作成に使用された http.Request
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func Hello(name string, language string) string {
	// return fmt.Sprintf("%s%s", englishHelloPrefix, name)
	if name == "" {
		name = "World"
	}

	return greeingPrefix(language) + name
}

// 関数のシグネチャでは、 named return value (prefix string)を作成しました。
// これにより、関数に prefix という変数が作成されます。
// - "zero" 値が割り当てられます。これはタイプによって異なります。たとえば、intは0で、stringの場合は""です。
// - これは関数のGo Docに表示されるので、コードの意図をより明確にすることができます。
func greeingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}

func main() {
	fmt.Println(Hello("Chris", ""))
	Greet(os.Stdout, "Elodie")
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
