/*
Copyright 2023 Daniel Hu.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
