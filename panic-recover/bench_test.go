package panic_recover

import (
	"errors"
	"testing"
)

// anti-optimizer placeholder
var DummyVariable interface{}

func panicRecover() {
	defer func() {
		DummyVariable = recover()
	}()

	panic("someArgument")
}

var err = errors.New("sample-error")

func errorFunction() error {
	return err
}

func BenchmarkPanicRecover(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		panicRecover()
	}
}

func BenchmarkErrorFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := errorFunction(); err != nil {
			DummyVariable = err
		}
	}
}
