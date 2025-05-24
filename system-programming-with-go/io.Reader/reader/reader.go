package main

import (
	"bytes"
	"io"
	"os"
	"strings"
)

func main() {
	var b bytes.Buffer
	b.Write([]byte("test bytes"))

	// 入力用バッファをメモリ上で確保していく
	buffer := make([]byte, 1024)
	// Read(p []bytes)のpは読みこんだ内容を一次的に保存するバッファ
	for i := 0; i < 3; i++ {
		size, _ := b.Read(buffer)
		println(size)
	}
	println(string(buffer))

	reader := strings.NewReader("テストデータ")
	io.Copy(os.Stdout, reader)
}
