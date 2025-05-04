package main

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(key string) (string, error) {
	value, exists := d[key]
	if !exists {
		return "", ErrNotFound
	}
	return value, nil
}

// ・マップの興味深い特性は、マップをポインタとして渡さなくても変更できることです。
// これは、 mapが参照型であるためです。
// ・参照型がもたらす落とし穴は、マップがnil値になる可能性があることです。
// nilマップは読み取り時に空のマップのように動作しますが、nilマップに書き込もうとすると、ランタイムパニックが発生します。
// したがって、空のマップ変数を初期化しないでください。
// var m map[string]string
// 代わりに、上記のように空のマップを初期化するか、makeキーワードを使用してマップを作成できます。
// var dictionary = map[string]string{}
// OR
// var dictionary = make(map[string]string)
func (d Dictionary) Add(key, value string) {
	d[key] = value
}

func Search(target map[string]string, key string) string {
	value, exists := target[key]
	if !exists {
		return ""
	}
	return value
}
