package bench

import (
	"testing"

	"github.com/xaionaro-go/benchmarks/internal/helpers"
)

var a [1000]int

func sliceToInterfaceSimple() uint64 {
	helpers.Dummy(a)
	return 0
}

func sliceToInterfaceTricky() uint64 {
	helpers.Dummy(a[:])
	return 0
}

func uint64ToInterface() uint64 {
	helpers.Dummy(uint64(0))
	return 0
}

func uint64ToUint64() uint64 {
	helpers.DummyUint64(uint64(0))
	return 0
}

func BenchmarkSliceToInterfaceSimple(b *testing.B) {
	helpers.Benchmark(b, sliceToInterfaceSimple)
}

func BenchmarkSliceToInterfaceTricky(b *testing.B) {
	helpers.Benchmark(b, sliceToInterfaceTricky)
}

func BenchmarkUint64ToInterface(b *testing.B) {
	helpers.Benchmark(b, uint64ToInterface)
}

func BenchmarkUint64ToUint64(b *testing.B) {
	helpers.Benchmark(b, uint64ToUint64)
}
