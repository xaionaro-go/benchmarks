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

func withTwoDefers() (r uint64) {
	defer func() {
		r = 1
	}()
	defer func() {
	}()
	return
}

func noDefer() uint64 {
	return 1
}

func BenchmarkDefer(b *testing.B) {
	helpers.Benchmark(b, withDefer)
}

func BenchmarkTwoDefers(b *testing.B) {
	helpers.Benchmark(b, withTwoDefers)
}

func BenchmarkNoDefer(b *testing.B) {
	helpers.Benchmark(b, noDefer)
}
