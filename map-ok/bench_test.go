package bench

import (
	"testing"
)

func BenchmarkValue(b *testing.B) {
	m := map[int]bool{}
	for i := 0; i<256; i++ {
		m[i<<1] = true
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i :=0; i<b.N; i++ {
		_ = m[i&0xff]
	}
}

func BenchmarkOK(b *testing.B) {
	m := map[int]bool{}
	for i := 0; i<256; i++ {
		m[i<<1] = true
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i :=0; i<b.N; i++ {
		_, _ = m[i&0xff]
	}
}
