package bench

import (
	"testing"

	"github.com/xaionaro-go/benchmarks/internal/helpers"
)

func withDefer() (r uint64) {
	defer func() {
		r = 1
	}()
	return
}

func noDefer() uint64 {
	return 1
}

func BenchmarkDefer_1goroutine(b *testing.B) {
	helpers.Benchmark(b, withDefer)
}

func BenchmarkNoDefer_1goroutine(b *testing.B) {
	helpers.Benchmark(b, noDefer)
}
