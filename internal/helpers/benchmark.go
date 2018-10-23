package helpers

import (
	"sync"
	"testing"
)

var someData uint64

func BenchmarkConcurrent(b *testing.B, routines int, fn func() uint64) uint64 {
	var wg sync.WaitGroup

	someData = 0

	b.ResetTimer()
	for c := 0; c < routines; c++ {
		wg.Add(1)
		go func() {
			for n := 0; n < b.N; n++ {
				fn()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	b.StopTimer()

	return someData
}

func Benchmark(b *testing.B, fn func() uint64) uint64 {
	someData = 0

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		fn()
	}
	b.StopTimer()

	return someData
}
