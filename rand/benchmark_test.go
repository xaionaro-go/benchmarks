package rand

import (
	"testing"

	mrand "math/rand"
	crand "crypto/rand"
)

func BenchmarkMathRandIntnSingle(b *testing.B) {
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		mrand.Intn(65536)
	}
}

func BenchmarkCryptoRandIntnSingle(b *testing.B) {
	var v [2]byte
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		crand.Read(v[:])
	}
}

func BenchmarkMathRandIntnParallel(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mrand.Intn(65536)
		}
	})
}

func BenchmarkCryptoRandIntnParallel(b *testing.B) {
	var v [2]byte
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			crand.Read(v[:])
		}
	})
}
