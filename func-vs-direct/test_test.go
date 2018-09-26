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

func BenchmarkDirect(b *testing.B) {
	for n := 0; n < b.N; n++ {
		helpers.Dummy(a)
	}
}

func BenchmarkFunc(b *testing.B) {
	for n := 0; n < b.N; n++ {
		helpers.Dummy(getA())
	}
}
