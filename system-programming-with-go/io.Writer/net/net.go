package main

import (
	"io"
	"net"
	"os"
)

func main() {
	// net.Connという通信コネクションを表すインターフェイスであり、
	// io.WriterReaderを満たす
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	io.WriteString(conn, "GET / HTTP/1.0\r\nHOST: ascii.jp\r\n\r\n")
	io.Copy(os.Stdout, conn)
}
