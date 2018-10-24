package bench

import (
	"testing"

	"github.com/xaionaro-go/benchmarks/internal/helpers"
)

var s = "0123456789abcdef0123456789abcdef"
var b = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
var r = []byte{}

func noCast() uint64 {
	r = make([]byte, len(b))
	copy(r, b)
	return 0
}

func doCast() uint64 {
	r = make([]byte, len(b))
	copy(r, []byte(s))
	return 0
}

func doCastSimple() uint64 {
	r = []byte(s)
	return 0
}

func BenchmarkBytesWithoutCast(bench *testing.B) {
	helpers.Benchmark(bench, noCast)
	r[0] = 1
	if b[0] != 0 {
		bench.Errorf("The source was been modified\n")
	}
}

func BenchmarkBytesWithCast(bench *testing.B) {
	helpers.Benchmark(bench, doCast)
	r[0] = 1
	if []byte(s)[0] != '0' {
		bench.Errorf("The source was been modified\n")
	}
}

func BenchmarkBytesWithCastSimple(bench *testing.B) {
	helpers.Benchmark(bench, doCastSimple)
	r[0] = 1
	if []byte(s)[0] != '0' {
		bench.Errorf("The source was been modified\n")
	}
}
