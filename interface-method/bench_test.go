package interfacemethod

import (
	"testing"
)

type Interface interface {
	PtrMethod()
	Method()
}

type Struct struct {
}

func (s *Struct) PtrMethod() {}
func (s Struct) Method()     {}

func BenchmarkPtrMethodDirect(b *testing.B) {
	s := &Struct{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.PtrMethod()
	}
}

func BenchmarkMethodDirect(b *testing.B) {
	s := &Struct{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Method()
	}
}

func BenchmarkPtrMethodViaInterface(b *testing.B) {
	var s Interface
	s = &Struct{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.PtrMethod()
	}
}

func BenchmarkMethodViaInterface(b *testing.B) {
	var s Interface
	s = &Struct{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Method()
	}
}
