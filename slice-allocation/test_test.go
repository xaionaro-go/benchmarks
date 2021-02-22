package bench

import (
	"testing"
	_ "unsafe"

	"github.com/xaionaro-go/benchmarks/internal/helpers"
	"github.com/xaionaro-go/unsafetools"
)

func BenchmarkAllocatedSliceIfaceBoxing(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s := []byte{}
		helpers.Dummy(s)
	}
}

func BenchmarkNilSliceIfaceBoxing(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		var s []byte
		helpers.Dummy(s)
	}
}

func BenchmarkMakeSlice8IfaceBoxing(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(8)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s := make([]byte, 8)
		helpers.Dummy(s)
	}
}

func BenchmarkMakeSliceUnsafeString8IfaceBoxing(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(8)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s := make([]byte, 8)
		helpers.Dummy(unsafetools.CastBytesToString(s))
	}
}

func BenchmarkArray8IfaceBoxing(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(8)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		var s [8]byte
		helpers.Dummy(s[:])
	}
}

func BenchmarkArrayUnsafeString8IfaceBoxing(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(8)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		var s [8]byte
		helpers.Dummy(unsafetools.CastBytesToString(s[:]))
	}
}

func BenchmarkAllocatedSlice(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s := []byte{}
		helpers.DummyBytes(s)
	}
}

func BenchmarkNilSlice(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		var s []byte
		helpers.DummyBytes(s)
	}
}

func BenchmarkMakeSlice8(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(8)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s := make([]byte, 8)
		helpers.DummyBytes(s)
	}
}

func BenchmarkArray8(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(8)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		var s [8]byte
		helpers.DummyBytes(s[:])
	}
}

func BenchmarkArrayUnsafeString8(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(8)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		var s [8]byte
		helpers.DummyString(unsafetools.CastBytesToString(s[:]))
	}
}
