package gc_blocksize

import (
	"fmt"
	"runtime"
	"testing"
)

func BenchmarkBlockSize(b *testing.B) {
	for _, blockSize := range []int{1, 1<<10, 1<<15, 10<<20, 10<<21, 10<<22} {
		b.Run(fmt.Sprintf("blockSize_%d", blockSize), func(b *testing.B) {
			for i:=0; i<b.N; i++ {
				b.StopTimer()
				_ = make([]byte, blockSize)
				b.StartTimer()
				runtime.GC()
				runtime.Gosched()
				runtime.GC()
			}
		})
	}
}
