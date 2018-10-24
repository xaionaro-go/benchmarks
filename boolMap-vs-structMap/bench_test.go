package bench

import (
	"testing"

	"github.com/xaionaro-go/benchmarks/internal/helpers"
)

func boolMap() uint64 {
	m := map[int]bool{}
	for i := 0; i < 1000000; i++ {
		m[i] = true
	}
	return 0
}

func structMap() uint64 {
	m := map[int]struct{}{}
	for i := 0; i < 1000000; i++ {
		m[i] = struct{}{}
	}
	return 0
}

func BenchmarkBoolMap(b *testing.B) {
	helpers.Benchmark(b, boolMap)
}

func BenchmarkStructMap(b *testing.B) {
	helpers.Benchmark(b, structMap)
}
