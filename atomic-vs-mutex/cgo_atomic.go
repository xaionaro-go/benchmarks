package bench

/*
#include <stdint.h>

void increment(void *a) {
	__sync_fetch_and_add((uint64_t *)a, 1);
}
*/
import "C"

import (
	"unsafe"
)

func increment(a *uint64) {
	C.increment((unsafe.Pointer)(a))
}
