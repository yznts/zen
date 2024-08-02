package slice

import (
	"testing"
)

func BenchmarkFilter(b *testing.B) {
	s := Range(1, 10000)
	for i := 0; i < b.N; i++ {
		_ = Filter(s, func(v int) bool { return v < 500 })
	}
}
