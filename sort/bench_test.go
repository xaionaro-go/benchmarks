package sort

import (
	"fmt"
	"math/rand"
	"testing"

	stdsort "sort"

	"github.com/TheAlgorithms/Go/sort"
)

func BenchmarkSort(b *testing.B) {
	unsorted := make([]uint, 1024*1024)
	rng := rand.New(rand.NewSource(0))
	for idx := range unsorted {
		unsorted[idx] = uint(rng.Uint64())
	}
	buf := make([]uint, 1024*1024)
	for _, testCase := range []struct {
		Name string
		Func func([]uint) []uint
	}{
		{Name: "NoSort", Func: func(u []uint) []uint {
			return u
		}},
		{Name: "StdQuicksort", Func: func(u []uint) []uint {
			stdsort.Slice(u, func(i, j int) bool {
				return u[i] < u[j]
			})
			return u
		}},
		{Name: "Bubble", Func: sort.Bubble[uint]},
		{Name: "Quicksort", Func: sort.Quicksort[uint]},
		{Name: "Merge", Func: sort.Merge[uint]},
	} {
		b.Run(testCase.Name, func(b *testing.B) {
			sortFunc := testCase.Func
			for _, sliceSize := range []uint64{0, 1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 10000} {
				b.Run(fmt.Sprintf("%d", sliceSize), func(b *testing.B) {
					buf = buf[:sliceSize]
					b.ReportAllocs()
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						copy(buf, unsorted)
						sortFunc(buf)
					}
				})
			}
		})
	}
}
