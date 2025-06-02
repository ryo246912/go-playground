package main

import (
	"path/filepath"
)

func main() {
	// path/filepathは物理的なパスで使用
	// pathは論理的なパスで使用
	println(filepath.Base("C:/path/to/file.txt"))     // file.txt
	println(filepath.Dir("C:/path/to/file.txt"))      // C:/path/to
	println(filepath.Clean("C:/path/to/../file.txt")) // C:/path/file.txt
	println(filepath.Ext("C:/path/to/file.txt"))      // .txt
	println(filepath.IsAbs("C:/path/to/file.txt"))    // true
	println(filepath.IsAbs("./file.txt"))             // false
	println(filepath.Join("C:/path", "/to/file.txt")) // C:/path/to/file.txt

	// 絶対パスを取得
	ab, err := filepath.Abs("./go-programming-essence/../file.txt")
	if err == nil {
		println(ab)
	}

}
