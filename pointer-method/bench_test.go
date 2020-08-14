package pointer_method

import (
	"testing"
)

type someStruct0 struct {
	A int
}

func (s someStruct0) Add(a int) someStruct0 {
	return someStruct0{
		A: s.A + a,
	}
}

type someStruct1 struct {
	A int
}

func (s *someStruct1) Add(a int) *someStruct1 {
	return &someStruct1{
		A: s.A + a,
	}
}

func Benchmark(b *testing.B) {
	b.Run("noPointer", func(b *testing.B) {
		s := someStruct0{}
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s = s.Add(1)
		}
	})
	b.Run("pointer", func(b *testing.B) {
		s := &someStruct1{}
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s = s.Add(1)
		}
	})
}
