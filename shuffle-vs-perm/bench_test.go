package shuffleVsPerm

import (
	"math/rand"
	"testing"
)

var (
	shortSlice []int
	mediumSlice []int
	longSlice []int
)

func init() {
	shortSlice = rand.Perm(4)
	mediumSlice = rand.Perm(256)
	longSlice = rand.Perm(65536)
}

func shuffleBenchmark(b *testing.B, slice []int) {
	for i:=0; i<b.N; i++ {
		rand.Shuffle(len(slice), func(i, j int) {
			slice[i], slice[j] = slice[j], slice[i]
		})
	}
}

func BenchmarkShuffleShort(b *testing.B) {
	shuffleBenchmark(b, shortSlice)
}

func BenchmarkShuffleMedium(b *testing.B) {
	shuffleBenchmark(b, mediumSlice)
}

func BenchmarkShuffleLong(b *testing.B) {
	shuffleBenchmark(b, longSlice)
}


func permBenchmark(b *testing.B, slice []int) {
	for i:=0; i<b.N; i++ {
		perm := rand.Perm(len(slice))
		for idx0, idx1 := range perm {
			perm[idx0] = slice[idx1]
		}
		slice = perm
	}
}

func BenchmarkPermShort(b *testing.B) {
	permBenchmark(b, shortSlice)
}

func BenchmarkPermMedium(b *testing.B) {
	permBenchmark(b, mediumSlice)
}

func BenchmarkPermLong(b *testing.B) {
	permBenchmark(b, longSlice)
}
