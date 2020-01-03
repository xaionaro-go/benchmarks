// This is an adapted version of https://gist.githubusercontent.com/mariomac/8518c01d44402244b2595360b867d60a/raw/a1023e92258ee9661069d9f64defd9dd223cebb7/quicksort.go
package javatest

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
)

const (
	itemsAmount = 1e8
)

func mariomacQuickSort(arr []int32, begin, end int) {
	if begin < end {
		partitionIndex := mariomacPartition(arr, begin, end)

		// Recursively sort elements of the 2 sub-arrays
		mariomacQuickSort(arr, begin, partitionIndex-1)
		mariomacQuickSort(arr, partitionIndex+1, end)
	}
}

func mariomacPartition(arr []int32, begin, end int) int {
	pivot := arr[end]
	i := (begin - 1)

	for j := begin; j < end; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[end] = arr[end], arr[i+1]

	return i + 1
}

type int32s []int32

func (s int32s) Len() int { return len(s) }
func (s int32s) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s int32s) Less(i, j int) bool { return s[i] < s[j] }

func mariomacNewSlice() int32s {
	return make([]int32, itemsAmount)
}

func mariomacFillRandom(s int32s) {
	for i:=0; i<itemsAmount; i++ {
		s[i] = int32(rand.Int())
	}
}

func BenchmarkMariomacAllocateSlice(b *testing.B) {
	b.ReportAllocs()

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		mariomacNewSlice()
	}
}

func BenchmarkMariomacFillRandom(b *testing.B) {
	b.ReportAllocs()

	s := mariomacNewSlice()

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		mariomacFillRandom(s)
	}
}

func BenchmarkMariomacQuickSort(b *testing.B) {
	b.ReportAllocs()

	s := mariomacNewSlice()

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		b.StopTimer()
		mariomacFillRandom(s)
		b.StartTimer()
		mariomacQuickSort(s, 0, itemsAmount-1)
	}
}

func TestMariomacQuickSort(t *testing.T) {
	s := mariomacNewSlice()
	mariomacFillRandom(s)
	mariomacQuickSort(s, 0, itemsAmount-1)
	mariomacCheck(s)
}


func mariomacCheck(s []int32) {
	for i := 1; i < itemsAmount; i++ {
		if s[i-1] > s[i] {
			fmt.Printf("Array not ordered in position %d: %d > %d", i,
				s[i-1], s[i])
			os.Exit(-1)
		}
	}
}
