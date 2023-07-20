package spinlock

import (
	"sync"
	"testing"
)

func TestSpinLock(t *testing.T) {
	var sl SpinLock
	var counter int

	// Start 1000 goroutines, each incrementing the counter 1000 times.
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				sl.Lock()
				counter++
				sl.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()

	// Check that the counter is correct.
	if counter != 1000*1000 {
		t.Errorf("counter = %d, want %d", counter, 1000*1000)
	}
}

func BenchmarkSpinLock(b *testing.B) {
	var sl SpinLock
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sl.Lock()
			_ = 1
			sl.Unlock()
		}
	})
}

func BenchmarkMutex(b *testing.B) {
	var mu sync.Mutex
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu.Lock()
			_ = 1
			mu.Unlock()
		}
	})
}
