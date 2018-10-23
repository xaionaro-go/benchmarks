package bench

import (
	"sync"
	"testing"
)

var m = &sync.Mutex{}

func mutexWrite() uint64 {
	m.Lock()
	someData++
	m.Unlock()
	return 0
}

func BenchmarkWriteMutex_1goroutine(b *testing.B) {
	benchmark(b, 1, mutexWrite)
}

func BenchmarkWriteMutex_2goroutine(b *testing.B) {
	benchmark(b, 2, mutexWrite)
}

func BenchmarkWriteMutex_4goroutine(b *testing.B) {
	benchmark(b, 4, mutexWrite)
}

func BenchmarkWriteMutex_16goroutine(b *testing.B) {
	benchmark(b, 16, mutexWrite)
}

func BenchmarkWriteMutex_64goroutine(b *testing.B) {
	benchmark(b, 64, mutexWrite)
}

func BenchmarkWriteMutex_256goroutine(b *testing.B) {
	benchmark(b, 256, mutexWrite)
}

func mutexRead() uint64 {
	m.Lock()
	defer m.Unlock()
	return someData
}

func BenchmarkReadMutex_1goroutine(b *testing.B) {
	benchmark(b, 1, mutexRead)
}

func BenchmarkReadMutex_2goroutine(b *testing.B) {
	benchmark(b, 2, mutexRead)
}

func BenchmarkReadMutex_4goroutine(b *testing.B) {
	benchmark(b, 4, mutexRead)
}

func BenchmarkReadMutex_16goroutine(b *testing.B) {
	benchmark(b, 16, mutexRead)
}

func BenchmarkReadMutex_64goroutine(b *testing.B) {
	benchmark(b, 64, mutexRead)
}

func BenchmarkReadMutex_256goroutine(b *testing.B) {
	benchmark(b, 256, mutexRead)
}
