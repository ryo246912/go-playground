package main

import "fmt"

type IPAddr [4]byte

// IPAddr 型を実装してみましょう IPアドレスをドットで4つに区切った( dotted quad )表現で出力するため、 fmt.Stringer インタフェースを実装してください。
// 例えば、 IPAddr{1, 2, 3, 4} は、 "1.2.3.4" として出力するようにします。
// なお、レシーバの型を*IPAddrにした場合、型*IPAddrはString()インタフェースを実装していますが、型IPAddrはString()インタフェースを実装していません。
func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
