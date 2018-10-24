package bench

import (
	"testing"

	"github.com/xaionaro-go/benchmarks/internal/helpers"
)

var slice = []int{0, 3, 5, 7, 9}

func shortSliceCheck() uint64 {
	for i := 0; i < 5; i++ {
		found := false
		for _, v := range slice {
			if v == i {
				found = true
				break
			}
		}
		if !found {
			continue
		}
		helpers.SomeData++
	}

	return 0
}

func boolMapCheck() uint64 {
	m := map[int]bool{}
	for _, v := range slice {
		m[v] = true
	}

	for i := 0; i < 5; i++ {
		if !m[i] {
			continue
		}
		helpers.SomeData++
	}

	return 0
}

func structMapCheck() uint64 {
	m := map[int]struct{}{}
	for _, v := range slice {
		m[v] = struct{}{}
	}

	for i := 0; i < 5; i++ {
		if _, ok := m[i]; !ok {
			continue
		}
		helpers.SomeData++
	}

	return 0
}

func BenchmarkShortSliceCheck(b *testing.B) {
	helpers.Benchmark(b, shortSliceCheck)
}

func BenchmarkBoolMapCheck(b *testing.B) {
	helpers.Benchmark(b, boolMapCheck)
}

func BenchmarkStructMapCheck(b *testing.B) {
	helpers.Benchmark(b, structMapCheck)
}
