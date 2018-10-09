package bench

import (
	"testing"

	"github.com/xaionaro-go/benchmarks/internal/helpers"
)

var (
	a = 1
)

func getA() int {
	return 1
}

func BenchmarkAllocatedSlice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := []int{}
		helpers.Dummy(s)
	}
}

func BenchmarkNilSlice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var s []int
		helpers.Dummy(s)
	}
}
