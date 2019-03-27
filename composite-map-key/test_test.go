package test

import (
	"testing"
)

func BenchmarkCompositeKeyMap(b *testing.B) {
	type k struct {
		a uint32
		b uint32
	}
	m := map[k]uint64{k{1,2}: 3}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = m[k{1,2}]
	}
	b.StopTimer()
}

func BenchmarkCompositeWithStringKeyMap(b *testing.B) {
	type k struct {
		a uint32
		b string
	}
	m := map[k]uint64{k{1,"\000\000\000\002"}: 3}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = m[k{1,"\000\000\000\002"}]
	}
	b.StopTimer()
}

func BenchmarkUint64KeyMap(b *testing.B) {
	m := map[uint64]uint64{(1 << 32) + 2: 3}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = m[(1 << 32) + 2]
	}
	b.StopTimer()
}

func BenchmarkStringKeyMap(b *testing.B) {
	m := map[string]uint64{"\000\000\000\001\000\000\000\002": 3}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = m["\000\000\000\001\000\000\000\002"]
	}
	b.StopTimer()
}

var k0 = "\000\000\000\001"
var k1 = "\000\000\000\002"
func BenchmarkCompositeKeyToStringKeyMap(b *testing.B) {
	m := map[string]uint64{"\000\000\000\001\000\000\000\002": 3}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = m[k0 + k1]
	}
	b.StopTimer()
}
