package bench

import (
	"reflect"
	"testing"

	"github.com/xaionaro-go/benchmarks/internal/helpers"
)

type A int
type B int

var (
	a = A(123)
	b = B(789)
)

func (b B) M(x int) int {
	return 0
}

func checkWithoutReflect(v interface{}) bool {
	_, has := v.(interface{ M(int) int })
	return has
}

func withoutReflect() uint64 {
	checkWithoutReflect(a)
	checkWithoutReflect(b)
	return 0
}

func checkWithReflect(i interface{}) bool {
	_, ok := reflect.TypeOf(i).MethodByName("M")
	return ok
}

func withReflect() uint64 {
	checkWithReflect(a)
	checkWithReflect(b)
	return 0
}

func BenchmarkCheckMethodWithoutReflect(bench *testing.B) {
	helpers.Benchmark(bench, withoutReflect)
}

func BenchmarkCheckMethodWithReflect(bench *testing.B) {
	helpers.Benchmark(bench, withReflect)
}
