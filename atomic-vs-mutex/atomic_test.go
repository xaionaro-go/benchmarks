package bench

import (
	"sync/atomic"
	"testing"
)

func atomicWrite() uint64 {
	atomic.AddUint64(&someData, 1)
	return 0
}

func atomicRead() uint64 {
	return atomic.LoadUint64(&someData)
}

func BenchmarkWriteAtomic_1goroutine(b *testing.B) {
	benchmark(b, 1, atomicWrite)
}

func BenchmarkWriteAtomic_2goroutine(b *testing.B) {
	benchmark(b, 2, atomicWrite)
}

func BenchmarkWriteAtomic_4goroutine(b *testing.B) {
	benchmark(b, 4, atomicWrite)
}

func BenchmarkWriteAtomic_16goroutine(b *testing.B) {
	benchmark(b, 16, atomicWrite)
}

func BenchmarkWriteAtomic_64goroutine(b *testing.B) {
	benchmark(b, 64, atomicWrite)
}

func BenchmarkWriteAtomic_256goroutine(b *testing.B) {
	benchmark(b, 256, atomicWrite)
}

func BenchmarkReadAtomic_1goroutine(b *testing.B) {
	benchmark(b, 1, atomicRead)
}

func BenchmarkReadAtomic_2goroutine(b *testing.B) {
	benchmark(b, 2, atomicRead)
}

func BenchmarkReadAtomic_4goroutine(b *testing.B) {
	benchmark(b, 4, atomicRead)
}

func BenchmarkReadAtomic_16goroutine(b *testing.B) {
	benchmark(b, 16, atomicRead)
}

func BenchmarkReadAtomic_64goroutine(b *testing.B) {
	benchmark(b, 64, atomicRead)
}

func BenchmarkReadAtomic_256goroutine(b *testing.B) {
	benchmark(b, 256, atomicRead)
}
