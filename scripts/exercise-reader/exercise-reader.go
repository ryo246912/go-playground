package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader) Read(bytes []byte) (int, error) {
	// ASCII文字 'A' の無限ストリームを出力する Reader 型を実装してください。
	for i := range bytes {
		bytes[i] = 'A'
	}

	return len(bytes), nil
}

func main() {
	reader.Validate(MyReader{})
}
