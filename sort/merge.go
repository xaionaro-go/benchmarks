package sort

import (
	"github.com/TheAlgorithms/Go/constraints"
)

var (
	buffer interface{}
	bufIdx int
)

func mergePreallocated[T constraints.Ordered](a []T, b []T) []T {
	resultLen := len(a) + len(b)
	buf, ok := buffer.([]T)
	if !ok {
		buf = make([]T, resultLen)
		buffer = buf
	}
	if cap(buf[bufIdx:]) < resultLen {
		newBuf := make([]T, cap(buf)*2)
		copy(newBuf, buf)
		buf = newBuf
		buffer = buf
	}
	r := buf[bufIdx : bufIdx+resultLen]
	bufIdx += resultLen
	var i = 0
	var j = 0

	for i < len(a) && j < len(b) {

		if a[i] <= b[j] {
			r[i+j] = a[i]
			i++
		} else {
			r[i+j] = b[j]
			j++
		}

	}

	for i < len(a) {
		r[i+j] = a[i]
		i++
	}
	for j < len(b) {
		r[i+j] = b[j]
		j++
	}

	return r
}

// mergeSortPreallocated performs merge sort on a slice
func mergeSortPreallocated[T constraints.Ordered](items []T) []T {

	if len(items) < 2 {
		return items
	}

	var middle = len(items) / 2
	var a = mergeSortPreallocated(items[:middle])
	var b = mergeSortPreallocated(items[middle:])
	return mergePreallocated(a, b)
}

func MergeSortPreallocated[T constraints.Ordered](items []T) []T {
	result := mergeSortPreallocated(items)
	bufIdx = 0
	return result
}
