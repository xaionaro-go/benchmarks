package reduce_number

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func simple32(src, mod uint32) uint32 {
	return src % mod
}

// http://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction/
func lemire32(src, mod uint32) uint32 {
	return uint32(((uint64)(src) * uint64(mod)) >> 32)
}

func okunev32(src, mod uint32) uint32 {
	mask := mod
	mask |= mask >> 1
	mask |= mask >> 2
	mask |= mask >> 4
	mask |= mask >> 8
	mask |= mask >> 16
	src &= mask
	if src < mod {
		return src
	}
	return src & (mask >> 1)
}

func simple64(src, mod uint64) uint64 {
	return src % mod
}

func okunev64(src, mod uint64) uint64 {
	mask := mod
	mask |= mask >> 1
	mask |= mask >> 2
	mask |= mask >> 4
	mask |= mask >> 8
	mask |= mask >> 16
	mask |= mask >> 32
	src &= mask
	if src < mod {
		return src
	}
	return src & (mask >> 1)
}

func benchmark32(b *testing.B, fn func(src, mod uint32) uint32) {
	e := b.N + 1
	b.ResetTimer()
	for i := 1; i < e; i++ {
		fn(3948558707, uint32(i))
	}
}

func benchmark64(b *testing.B, fn func(src, mod uint64) uint64) {
	e := b.N + 1
	b.ResetTimer()
	for i := 1; i < e; i++ {
		fn(8963315421273233617, uint64(i))
	}
}

func BenchmarkIntegerReductionSimple32(b *testing.B) {
	benchmark32(b, simple32)
}

func BenchmarkIntegerReductionLemire32(b *testing.B) {
	benchmark32(b, lemire32)
}

func BenchmarkIntegerReductionOkunev32(b *testing.B) {
	benchmark32(b, okunev32)
}

func BenchmarkIntegerReductionSimple64(b *testing.B) {
	benchmark64(b, simple64)
}

func BenchmarkIntegerReductionOkunev64(b *testing.B) {
	benchmark64(b, okunev64)
}

func Test(t *testing.T) {
	for idx, reduceFunc := range []func(uint32, uint32) uint32{
		simple32,
		lemire32,
		okunev32,
	} {
		for try:=0; try<100; try++ {
			mod := uint32(rand.Int31n(256))+1
			var countMax, countZero int
			for i:=0; i<100000; i++ {
				r := reduceFunc(uint32(rand.Int63n(math.MaxUint32)), mod)
				if r == mod-1 {
					countMax++
				}
				if r == 0 {
					countZero++
				}
			}

			assert.True(t, float64(countMax)/float64(countZero) > 0.80, fmt.Sprint(idx, countMax, countZero))
			assert.True(t, float64(countMax)/float64(countZero) < 1.20, fmt.Sprint(idx, countMax, countZero))
		}
	}
}