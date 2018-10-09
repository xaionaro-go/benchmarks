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

func BenchmarkSprintf(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fmt.Sprintf("%v", -42)
	}
}

func BenchmarkFormatInt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		strconv.FormatInt(-42, 10)
	}
}
