package sha1

import (
	"crypto/sha1"
	"testing"
)

func BenchmarkSHA1Sequential(b *testing.B) {
	hasher := sha1.New()
	b.ReportAllocs()
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		hasher.Write([]byte{0})
		hasher.Sum(nil)
		hasher.Reset()
	}
}

func BenchmarkSHA1Parallel(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		hasher := sha1.New()
		for pb.Next() {
			hasher.Write([]byte{0})
			hasher.Sum(nil)
			hasher.Reset()
		}
	})
}
