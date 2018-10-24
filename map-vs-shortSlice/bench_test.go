package bench

import (
	"testing"

	"github.com/xaionaro-go/benchmarks/internal/helpers"
)

var numIterations = 0
var slice = []int{0, 3, 5, 7, 9}

func shortSliceCheck() uint64 {
	for i := 0; i < numIterations; i++ {
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

	for i := 0; i < numIterations; i++ {
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

	for i := 0; i < numIterations; i++ {
		if _, ok := m[i]; !ok {
			continue
		}
		helpers.SomeData++
	}

	return 0
}

func BenchmarkShortSliceCheck_iterations2(b *testing.B) {
	numIterations = 2
	helpers.Benchmark(b, shortSliceCheck)
}

func BenchmarkBoolMapCheck_iterations2(b *testing.B) {
	numIterations = 2
	helpers.Benchmark(b, boolMapCheck)
}

func BenchmarkStructMapCheck_iterations2(b *testing.B) {
	numIterations = 2
	helpers.Benchmark(b, structMapCheck)
}

func BenchmarkShortSliceCheck_iterations10(b *testing.B) {
	numIterations = 10
	helpers.Benchmark(b, shortSliceCheck)
}

func BenchmarkBoolMapCheck_iterations10(b *testing.B) {
	numIterations = 10
	helpers.Benchmark(b, boolMapCheck)
}

func BenchmarkStructMapCheck_iterations10(b *testing.B) {
	numIterations = 10
	helpers.Benchmark(b, structMapCheck)
}

func BenchmarkShortSliceCheck_iterations100(b *testing.B) {
	numIterations = 100
	helpers.Benchmark(b, shortSliceCheck)
}

func BenchmarkBoolMapCheck_iterations100(b *testing.B) {
	numIterations = 100
	helpers.Benchmark(b, boolMapCheck)
}

func BenchmarkStructMapCheck_iterations100(b *testing.B) {
	numIterations = 100
	helpers.Benchmark(b, structMapCheck)
}
