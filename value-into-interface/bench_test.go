package bench

import (
	"testing"

	"github.com/xaionaro-go/benchmarks/internal/helpers"
)

var a [1000]int

func simple() uint64 {
	helpers.Dummy(a)
	return 0
}

func tricky() uint64 {
	helpers.Dummy(a[:])
	return 0
}

func BenchmarkSimple(b *testing.B) {
	helpers.Benchmark(b, simple)
}

func BenchmarkTricky(b *testing.B) {
	helpers.Benchmark(b, tricky)
}

