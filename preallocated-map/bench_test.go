package preallocatedmap

import (
	"fmt"
	"testing"
)

func BenchmarkMap(b *testing.B) {
	for _, itemsCount := range []int{0, 1, 8, 256, 65536, 1024*1024} {
		b.Run(fmt.Sprintf("itemsCount%d", itemsCount), func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i:=0; i<b.N; i++ {
				m := map[int]struct{}{}
				for v:=0; v<itemsCount; v++ {
					m[v] = struct{}{}
				}
			}
		})
	}
}

func BenchmarkPreallocatedMap(b *testing.B) {
	for _, itemsCount := range []int{8, 256, 65536, 1024 * 1024} {
		b.Run(fmt.Sprintf("itemsCount%d", itemsCount), func(b *testing.B) {
			for _, preallocatedSpace := range []int{0, itemsCount, itemsCount*3/2, itemsCount*2, itemsCount*3, itemsCount*8} {
				b.Run(fmt.Sprintf("preallocated%d", preallocatedSpace), func(b *testing.B) {
					b.ReportAllocs()
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						m := make(map[int]struct{}, preallocatedSpace)
						for v := 0; v < itemsCount; v++ {
							m[v] = struct{}{}
						}
					}
				})
			}
		})
	}
}