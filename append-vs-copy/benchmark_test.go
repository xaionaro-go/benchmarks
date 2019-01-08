package benchmark

import (
	"testing"
)

const (
	size = 256
)

func BenchmarkAppend(b *testing.B) {
	buf := make([]byte, size)
	val := make([]byte, size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf = buf[:0]
		buf = append(buf, val...)
	}
}

func BenchmarkCopy(b *testing.B) {
	buf := make([]byte, size)
	val := make([]byte, size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf = buf[:size]
		copy(buf, val)
	}
}
