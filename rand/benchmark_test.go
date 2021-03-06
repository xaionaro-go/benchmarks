package rand

import (
	"sync/atomic"
	"testing"

	crand "crypto/rand"
	mrand "math/rand"
)

var count uint64

func counterIntn(module uint64) uint64 {
	return (atomic.AddUint64(&count, 6700417) / 115249) % module
}

func BenchmarkCounterIntnSingle(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		counterIntn(65536)
	}
}

func BenchmarkMathRandIntnSingle(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mrand.Intn(65536)
	}
}

func BenchmarkCryptoRandIntnSingle(b *testing.B) {
	var v [2]byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		crand.Read(v[:])
	}
}

func BenchmarkCounterIntnParallel(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			counterIntn(65536)
		}
	})
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
