package bench

import (
	"fmt"
	"strconv"
	"testing"
)

var (
	a = 1
)

func getA() int {
	return 1
}

func BenchmarkSprintfValue(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fmt.Sprintf("%v", -42)
	}
}

func BenchmarkSprintfInt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fmt.Sprintf("%v", -42)
	}
}

func BenchmarkFormatInt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		strconv.FormatInt(-42, 10)
	}
}

func BenchmarkFormatFloat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		strconv.FormatFloat(-42, 'f', 0, 64)
	}
}

func BenchmarkFormatFloatConcat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = strconv.FormatFloat(-42, 'f', 0, 64)+`%`
	}
}

func BenchmarkSprintfFloatConcat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fmt.Sprintf(`%f%%`, float64(-42))
	}
}
