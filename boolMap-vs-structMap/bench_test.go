package bench

import (
	"testing"

	"github.com/xaionaro-go/benchmarks/internal/helpers"
)

var m0 map[int]bool
var m1 map[int]struct{}

func boolMap() uint64 {
	m0 = map[int]bool{}
	for i:=0; i<1000000; i++ {
		m0[i] = true
	}
	return 0
}

func structMap() uint64 {
	m1 = map[int]struct{}{}
	for i:=0; i<1000000; i++ {
		m1[i] = struct{}{}
	}
	return 0
}

func BenchmarkBoolMap(b *testing.B) {
	helpers.Benchmark(b, boolMap)
}

func BenchmarkStructMap(b *testing.B) {
	helpers.Benchmark(b, structMap)
}
