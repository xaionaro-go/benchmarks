package panic_recover

import (
	"errors"
	"testing"
)

var err = errors.New("sample-error")

// anti-optimizer placeholder
var DummyVariable interface{}

func panicRecover() {
	defer func() {
		DummyVariable = recover()
	}()

	panic("someArgument")
}

var panicRecoverDeep func()
var errorFunctionDeep func() error

func init() {
	panicRecoverDeep = func() {
		panic("someArgument")
	}
	for i := 0; i < 99; i++ {
		oldF := panicRecoverDeep
		panicRecoverDeep = func() {
			oldF()
		}
	}
	{
		oldF := panicRecoverDeep
		panicRecoverDeep = func() {
			defer func() {
				recover()
			}()
			oldF()
		}
	}

	errorFunctionDeep = func() error {
		return err
	}
	for i := 0; i < 100; i++ {
		oldF := errorFunctionDeep
		errorFunctionDeep = func() error {
			return oldF()
		}
	}
}

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
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := errorFunction(); err != nil {
			DummyVariable = err
		}
	}
}

func BenchmarkPanicRecoverDeep(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		panicRecoverDeep()
	}
}

func BenchmarkErrorFuncDeep(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := errorFunctionDeep(); err != nil {
			DummyVariable = err
		}
	}
}
