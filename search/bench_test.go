package search

import (
	"fmt"
	"reflect"
	"runtime"
	"sort"
	"sync"
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
			reflect.ValueOf(&s[idx]).Elem().SetString(fmt.Sprintf("%08d", idx))
		default:
			panic(fmt.Sprintf("%T", s[0]))
		}
	}
	return s
}

func benchmarkGetLinear[T constraints.Ordered](b *testing.B, size uint) {
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

func benchmarkGetBinary[T constraints.Ordered](b *testing.B, size uint) {
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

func benchmarkGetMap[T constraints.Ordered](b *testing.B, size uint) {
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

func benchmarkGetSyncMap[T constraints.Ordered](b *testing.B, size uint) {
	rng := mathrand.NewWithSeed(0)
	s := makeSlice[T](size)
	var m sync.Map
	for i := 0; i < int(size); i++ {
		m.Store(s[i], i)
	}
	runtime.GC()
	runtime.GC()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := s[mathrand.ReduceUint32(rng.Uint32Xorshift(), uint32(size))]

		// find:
		_, _ = m.Load(v)
	}
}

func benchmarkAddLinear[T constraints.Ordered](b *testing.B, size uint) {
	values := makeSlice[T](size)
	runtime.GC()
	runtime.GC()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var s []T
		for i := uint(0); i < size; i++ {
			s = append(s, values[i])
		}
		_ = s
	}
}

func benchmarkAddBinary[T constraints.Ordered](b *testing.B, size uint) {
	values := makeSlice[T](size)
	runtime.GC()
	runtime.GC()
	b.Run("sortOnce", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var s []T
			for i := uint(0); i < size; i++ {
				s = append(s, values[i])
			}
			_ = s
			sort.Slice(s, func(i, j int) bool {
				return s[i] < s[j]
			})
		}
	})
	runtime.GC()
	runtime.GC()
	b.Run("sortEach", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var s []T
			for i := uint(0); i < size; i++ {
				s = append(s, values[i])
				sort.Slice(s, func(i, j int) bool {
					return s[i] < s[j]
				})
			}
			_ = s
		}
	})
}

func benchmarkAddMap[T constraints.Ordered](b *testing.B, size uint) {
	values := makeSlice[T](size)
	m := map[T]uint{}
	runtime.GC()
	runtime.GC()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := uint(0); i < size; i++ {
			m[values[i]] = i
		}
	}
}

func benchmarkAddSyncMap[T constraints.Ordered](b *testing.B, size uint) {
	values := makeSlice[T](size)
	for _, withLoad := range []bool{false, true} {
		b.Run(fmt.Sprintf("withLoad-%v", withLoad), func(b *testing.B) {
			var m sync.Map
			runtime.GC()
			runtime.GC()
			b.ReportAllocs()
			b.ResetTimer()
			if withLoad {
				for i := 0; i < b.N; i++ {
					for i := uint(0); i < size; i++ {
						_, _ = m.LoadOrStore(values[i], i)
					}
				}
			} else {
				for i := 0; i < b.N; i++ {
					for i := uint(0); i < size; i++ {
						m.Store(values[i], i)
					}
				}
			}
		})
	}
}

func Benchmark(b *testing.B) {
	type benchFunc func(b *testing.B, size uint)
	for _, actionType := range []string{"get", "add"} {
		b.Run(actionType, func(b *testing.B) {
			for _, keyType := range []reflect.Kind{reflect.Int, reflect.String} {
				var _benchmarkLinear, _benchmarkBinary, _benchmarkMap, _benchmarkSyncMap benchFunc
				limit := uint(1024 * 1024)
				switch actionType {
				case "get":
					switch keyType {
					case reflect.Int:
						_benchmarkLinear = benchmarkGetLinear[int]
						_benchmarkBinary = benchmarkGetBinary[int]
						_benchmarkMap = benchmarkGetMap[int]
						_benchmarkSyncMap = benchmarkGetSyncMap[int]
					case reflect.String:
						_benchmarkLinear = benchmarkGetLinear[string]
						_benchmarkBinary = benchmarkGetBinary[string]
						_benchmarkMap = benchmarkGetMap[string]
						_benchmarkSyncMap = benchmarkGetSyncMap[string]
					default:
						panic(keyType.String())
					}
				case "add":
					limit = 2048
					switch keyType {
					case reflect.Int:
						_benchmarkLinear = benchmarkAddLinear[int]
						_benchmarkBinary = benchmarkAddBinary[int]
						_benchmarkMap = benchmarkAddMap[int]
						_benchmarkSyncMap = benchmarkAddSyncMap[int]
					case reflect.String:
						_benchmarkLinear = benchmarkAddLinear[string]
						_benchmarkBinary = benchmarkAddBinary[string]
						_benchmarkMap = benchmarkAddMap[string]
						_benchmarkSyncMap = benchmarkAddSyncMap[string]
					default:
						panic(keyType.String())
					}
				default:
					panic(actionType)
				}
				b.Run(keyType.String(), func(b *testing.B) {
					for _, testCase := range []struct {
						Name      string
						BenchFunc benchFunc
					}{
						{Name: "Linear", BenchFunc: _benchmarkLinear},
						{Name: "Binary", BenchFunc: _benchmarkBinary},
						{Name: "Map", BenchFunc: _benchmarkMap},
						{Name: "SyncMap", BenchFunc: _benchmarkSyncMap},
					} {
						b.Run(testCase.Name, func(b *testing.B) {
							for size := uint(1); size < limit; size *= 2 {
								b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
									testCase.BenchFunc(b, size)
								})
							}
						})
					}
				})
			}
		})
	}
}
