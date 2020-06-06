package gc_blocksize

import (
	"fmt"
	"runtime"
	"testing"
)

func BenchmarkBlockSizeAllocateAndGC(b *testing.B) {
	for _, blockSize := range []int{1, 1<<10, 1<<15, 10<<20, 10<<21, 10<<22} {
		b.Run(fmt.Sprintf("blockSize_%d", blockSize), func(b *testing.B) {
			for i:=0; i<b.N; i++ {
				_ = make([]byte, blockSize)
				runtime.GC()
			}
		})
	}
}

func BenchmarkBlockSizeAllocate(b *testing.B) {
	for _, blockSize := range []int{1, 1<<10, 1<<15, 10<<20, 10<<21, 10<<22} {
		b.Run(fmt.Sprintf("blockSize_%d", blockSize), func(b *testing.B) {
			for i:=0; i<b.N; i++ {
				_ = make([]byte, blockSize)
			}
			runtime.GC()
		})
	}
}

func BenchmarkBlockSizeGC(b *testing.B) {
	for _, blockSize := range []int{1, 1<<10, 1<<15, 10<<20, 10<<21, 10<<22} {
		b.Run(fmt.Sprintf("blockSize_%d", blockSize), func(b *testing.B) {
			for i:=0; i<b.N; i++ {
				b.StopTimer()
				_ = make([]byte, blockSize)
				b.StartTimer()
				runtime.GC()
			}
		})
	}
}
