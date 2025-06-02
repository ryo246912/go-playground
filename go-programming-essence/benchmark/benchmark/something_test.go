package something

import "testing"

func BenchmarkDoSometing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doSomething()
	}
}
