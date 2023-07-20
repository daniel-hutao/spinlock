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
	"runtime"
	"sync/atomic"
)

// SpinLock is a simple spin lock implementation with exponential backoff and adaptive spinning.
type SpinLock struct {
	flag uint32
}

// Lock spins until it is able to acquire the lock.
// It uses exponential backoff and adaptive spinning to reduce contention.
func (sl *SpinLock) Lock() {
	backoff := 1
	for !sl.TryLock() {
		for i := 0; i < backoff; i++ {
			runtime.Gosched()
		}
		if backoff < 128 { // Limit the maximum backoff time
			backoff *= 2
		}
	}
}

// Unlock releases the lock.
func (sl *SpinLock) Unlock() {
	atomic.StoreUint32(&sl.flag, 0)
}

// TryLock attempts to acquire the lock without blocking.
// It returns true if the lock was successfully acquired, and false otherwise.
func (sl *SpinLock) TryLock() bool {
	return atomic.CompareAndSwapUint32(&sl.flag, 0, 1)
}
