package replacer

import (
	"sync"
	"testing"

	"github.com/trafficstars/spinlock"
)

func BenchmarkMutexLockUnlock(b *testing.B) {
	locker := sync.Mutex{}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			locker.Lock()
			locker.Unlock()
		}
	})
}

func BenchmarkSpinlockLockUnlock(b *testing.B) {
	locker := &spinlock.Locker{}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			locker.Lock()
			locker.Unlock()
		}
	})
}
