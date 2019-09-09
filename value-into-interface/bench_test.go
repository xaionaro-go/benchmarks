package bench

import (
	"testing"

	"github.com/xaionaro-go/benchmarks/internal/helpers"
)

var sl = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var a [1000]int
var i = uint64(0)
var s = `a string`

func sliceToInterface() uint64 {
	helpers.Dummy(sl)
	return 0
}

func arrayToInterfaceSimple() uint64 {
	helpers.Dummy(a)
	return 0
}

func arrayToInterfaceTricky() uint64 {
	helpers.Dummy(a[:])
	return 0
}

func uint64ToInterface() uint64 {
	helpers.Dummy(i)
	return 0
}

func uint64ToUint64() uint64 {
	helpers.DummyUint64(i)
	return 0
}

func stringToInterface() uint64 {
	helpers.Dummy(s)
	return 0
}

func BenchmarkSliceToInterface(b *testing.B) {
	helpers.Benchmark(b, sliceToInterface)
}

func BenchmarkArrayToInterfaceSimple(b *testing.B) {
	helpers.Benchmark(b, arrayToInterfaceSimple)
}

func BenchmarkArrayToInterfaceTricky(b *testing.B) {
	helpers.Benchmark(b, arrayToInterfaceTricky)
}

func BenchmarkUint64ToInterface(b *testing.B) {
	helpers.Benchmark(b, uint64ToInterface)
}

func BenchmarkUint64ToUint64(b *testing.B) {
	helpers.Benchmark(b, uint64ToUint64)
}

func BenchmarkStringToInterface(b *testing.B) {
	helpers.Benchmark(b, stringToInterface)
}

func BenchmarkConcurrent32StringToInterface(b *testing.B) {
	helpers.BenchmarkConcurrent(b, 32, stringToInterface)
}
