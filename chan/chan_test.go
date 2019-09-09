package chan_test

import (
	"testing"
)

func BenchmarkChan(b *testing.B) {
	ch := make(chan struct{})

	go func() {
		for {
			<-ch
		}
	}()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch <- struct{}{}
	}
}
