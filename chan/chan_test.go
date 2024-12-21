package chan_test

import (
	"testing"

	"github.com/eapache/queue/v2"
)

func BenchmarkChan0(b *testing.B) {
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

func BenchmarkChan1(b *testing.B) {
	ch := make(chan struct{}, 1)

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

func BenchmarkChan2(b *testing.B) {
	ch := make(chan struct{}, 2)

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

func BenchmarkChan1024(b *testing.B) {
	ch := make(chan struct{}, 1024)

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

func BenchmarkChan1024WithSelect(b *testing.B) {
	ch0 := make(chan struct{}, 1024)
	ch1 := make(chan struct{}, 1024)

	go func() {
		for {
			select {
			case <-ch0:
			case <-ch1:
			}
		}
	}()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch0 <- struct{}{}
	}
}

func BenchmarkChan1PutGet(b *testing.B) {
	b.ReportAllocs()
	ch := make(chan struct{}, 1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch <- struct{}{}
		<-ch
	}
}

func BenchmarkRingBufferPutGet(b *testing.B) {
	b.ReportAllocs()
	ch := queue.New[struct{}]()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch.Add(struct{}{})
		ch.Peek()
	}
}
