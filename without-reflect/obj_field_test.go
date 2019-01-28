package bench

import (
	"reflect"
	"testing"

	"github.com/xaionaro-go/benchmarks/internal/helpers"
)

type objT struct {
	A uint64
}

var (
	obj = &objT{A: 1}
)

func fieldViaReflect() uint64 {
	return reflect.ValueOf(obj).Elem().FieldByName(`A`).Interface().(uint64)
}

func fieldDirect() uint64 {
	return obj.A
}

func BenchmarkGetFieldValueWithoutReflect(bench *testing.B) {
	helpers.Benchmark(bench, fieldDirect)
}

func BenchmarkGetFieldValueWithReflect(bench *testing.B) {
	helpers.Benchmark(bench, fieldViaReflect)
}
