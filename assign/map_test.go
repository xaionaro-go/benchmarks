package bench

import (
	"testing"

	"github.com/xaionaro-go/benchmarks/internal/helpers"
)

var a = map[string]interface{}{
	"a": "a",
	"b": "b",
	"c": "c",
	"1": 1,
	"i": map[string]interface{}{},
}
var b = map[string]interface{}{}

func dummy() uint64 {
	return 0
}

func assignMap() uint64 {
	b = a
	helpers.Dummy(b)
	return 0
}

func BenchmarkDummy(b *testing.B) {
	helpers.Benchmark(b, dummy)
}

func BenchmarkAssignMap(b *testing.B) {
	helpers.Benchmark(b, assignMap)
}
