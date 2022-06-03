package search

import (
	"fmt"
	"reflect"
	"runtime"
	"sort"
	"testing"

	"github.com/xaionaro-go/rand/mathrand"
	"golang.org/x/exp/constraints"
)

func makeSlice[T constraints.Ordered](size uint) []T {
	s := make([]T, size)
	for idx := range s {
		switch reflect.ValueOf(s[0]).Kind() {
		case reflect.Int:
			reflect.ValueOf(&s[idx]).Elem().SetInt(int64(idx))
		case reflect.String:
			reflect.ValueOf(&s[idx]).Elem().SetString(fmt.Sprintf("%d", idx))
		default:
			panic(fmt.Sprintf("%T", s[0]))
		}
	}
	return s
}

func benchmarkLinear[T constraints.Ordered](b *testing.B, size uint) {
	rng := mathrand.NewWithSeed(0)
	s := makeSlice[T](size)
	runtime.GC()
	runtime.GC()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := s[mathrand.ReduceUint32(rng.Uint32Xorshift(), uint32(size))]

		// find:
		for _, cmp := range s {
			if cmp == v {
				break // found
			}
		}
	}
}

func benchmarkBinary[T constraints.Ordered](b *testing.B, size uint) {
	rng := mathrand.NewWithSeed(0)
	s := makeSlice[T](size)
	runtime.GC()
	runtime.GC()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := s[mathrand.ReduceUint32(rng.Uint32Xorshift(), uint32(size))]

		// find:
		sort.Search(int(size), func(i int) bool {
			return s[i] >= v
		})
	}
}

func benchmarkMap[T constraints.Ordered](b *testing.B, size uint) {
	rng := mathrand.NewWithSeed(0)
	s := makeSlice[T](size)
	m := make(map[T]int, size)
	for i := 0; i < int(size); i++ {
		m[s[i]] = i
	}
	runtime.GC()
	runtime.GC()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := s[mathrand.ReduceUint32(rng.Uint32Xorshift(), uint32(size))]

		// find:
		_ = m[v]
	}
}

func BenchmarkSearch(b *testing.B) {
	type benchFunc func(b *testing.B, size uint)
	for _, keyType := range []reflect.Kind{reflect.Int, reflect.String} {
		var _benchmarkLinear, _benchmarkBinary, _benchmarkMap benchFunc
		switch keyType {
		case reflect.Int:
			_benchmarkLinear = benchmarkLinear[int]
			_benchmarkBinary = benchmarkBinary[int]
			_benchmarkMap = benchmarkMap[int]
		case reflect.String:
			_benchmarkLinear = benchmarkLinear[string]
			_benchmarkBinary = benchmarkBinary[string]
			_benchmarkMap = benchmarkMap[string]
		}
		b.Run(keyType.String(), func(b *testing.B) {
			for _, testCase := range []struct {
				Name      string
				BenchFunc benchFunc
			}{
				{Name: "Linear", BenchFunc: _benchmarkLinear},
				{Name: "Binary", BenchFunc: _benchmarkBinary},
				{Name: "Map", BenchFunc: _benchmarkMap},
			} {
				b.Run(testCase.Name, func(b *testing.B) {
					for size := uint(1); size < 1024*1024; size *= 2 {
						b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
							testCase.BenchFunc(b, size)
						})
					}
				})
			}
		})
	}
}
