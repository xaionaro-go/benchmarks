package main

import (
	"testing"

	"github.com/xaionaro-go/benchmarks/internal/helpers"
)

func BenchmarkToPointerSimple(b *testing.B) {
	var result *int
	for n := 0; n < b.N; n++ {
		{
			tempVariable := n
			result = &tempVariable
		}
	}

	helpers.Dummy(result)
}

func BenchmarkToPointerSlice(b *testing.B) {
	var result *int
	for n := 0; n < b.N; n++ {
		result = &[]int{n}[0]
	}

	helpers.Dummy(result)
}

func BenchmarkToPointerMake(b *testing.B) {
	var result *int
	for n := 0; n < b.N; n++ {
		result = new(int)
		*result = n
	}

	helpers.Dummy(result)
}
