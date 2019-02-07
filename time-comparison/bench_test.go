package test

import (
	`testing`
	`time`
)

var (
	SomeBool bool
	SomeDuration0 = time.Second
	SomeDuration1 = time.Second
)

func BenchmarkTimeDurationIsEqual(b *testing.B) {
	for i:=0; i<b.N; i++ {
		SomeBool = SomeDuration0 == SomeDuration1
	}
}

var (
	SomeTime0 time.Time
	SomeTime1 time.Time
)

func BenchmarkTimeTimeIsEqual(b *testing.B) {
	for i:=0; i<b.N; i++ {
		SomeBool = SomeTime0 == SomeTime1
	}
}

var (
	SomeUint64_0 = uint64(1)
	SomeUint64_1 = uint64(1)
)

func BenchmarkUint64IsEqual(b *testing.B) {
	for i:=0; i<b.N; i++ {
		SomeBool = SomeUint64_0 == SomeUint64_1
	}
}
