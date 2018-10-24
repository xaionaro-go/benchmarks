package bench

import (
	"sync"
	"testing"
)

func benchmarkWriteCGOAtomic(b *testing.B, routines int) {
	var wg sync.WaitGroup

	someData = 0

	b.ResetTimer()
	for c := 0; c < routines; c++ {
		wg.Add(1)
		go func() {
			for n := 0; n < b.N; n++ {
				increment(&someData)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	b.StopTimer()

	if someData != uint64(b.N)*uint64(routines) {
		b.Errorf("someData != b.N: %v != %v", someData, b.N)
	}
}

func BenchmarkWriteCGOAtomic_1goroutine(b *testing.B) {
	benchmarkWriteCGOAtomic(b, 1)
}

func BenchmarkWriteCGOAtomic_2goroutine(b *testing.B) {
	benchmarkWriteCGOAtomic(b, 2)
}

func BenchmarkWriteCGOAtomic_4goroutine(b *testing.B) {
	benchmarkWriteCGOAtomic(b, 4)
}

func BenchmarkWriteCGOAtomic_16goroutine(b *testing.B) {
	benchmarkWriteCGOAtomic(b, 16)
}

func BenchmarkWriteCGOAtomic_64goroutine(b *testing.B) {
	benchmarkWriteCGOAtomic(b, 64)
}

func BenchmarkWriteCGOAtomic_256goroutine(b *testing.B) {
	benchmarkWriteCGOAtomic(b, 256)
}
