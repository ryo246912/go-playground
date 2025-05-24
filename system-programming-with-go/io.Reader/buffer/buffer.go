package main

import (
	"bytes"
	"os"
)

func main() {
	var buffer1 bytes.Buffer
	buffer2 := bytes.NewBuffer([]byte{0x10, 0x20, 0x30})
	buffer3 := bytes.NewBufferString("初期文字列")

	os.Stdout.Write(buffer1.Bytes())
	os.Stdout.Write(buffer2.Bytes())
	os.Stdout.Write(buffer3.Bytes())
}
