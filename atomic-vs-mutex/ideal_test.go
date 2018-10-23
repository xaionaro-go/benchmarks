package bench

import (
	"testing"
)

func idealWrite() uint64 {
	someData++
	return 0
}

func idealRead() uint64 {
	return someData
}

func BenchmarkWriteIdeal_1goroutine(b *testing.B) {
	benchmark(b, 1, idealWrite)
}

func BenchmarkWriteIdeal_2goroutine(b *testing.B) {
	benchmark(b, 2, idealWrite)
}

func BenchmarkWriteIdeal_4goroutine(b *testing.B) {
	benchmark(b, 4, idealWrite)
}

func BenchmarkWriteIdeal_16goroutine(b *testing.B) {
	benchmark(b, 16, idealWrite)
}

func BenchmarkWriteIdeal_64goroutine(b *testing.B) {
	benchmark(b, 64, idealWrite)
}

func BenchmarkWriteIdeal_256goroutine(b *testing.B) {
	benchmark(b, 256, idealWrite)
}

func BenchmarkReadIdeal_1goroutine(b *testing.B) {
	benchmark(b, 1, idealRead)
}

func BenchmarkReadIdeal_2goroutine(b *testing.B) {
	benchmark(b, 2, idealRead)
}

func BenchmarkReadIdeal_4goroutine(b *testing.B) {
	benchmark(b, 4, idealRead)
}

func BenchmarkReadIdeal_16goroutine(b *testing.B) {
	benchmark(b, 16, idealRead)
}

func BenchmarkReadIdeal_64goroutine(b *testing.B) {
	benchmark(b, 64, idealRead)
}

func BenchmarkReadIdeal_256goroutine(b *testing.B) {
	benchmark(b, 256, idealRead)
}
