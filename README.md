# SpinLock

[![build and test](https://github.com/daniel-hutao/spinlock/workflows/CI/badge.svg)](https://github.com/daniel-hutao/spinlock/actions)
[![go report](https://goreportcard.com/badge/github.com/daniel-hutao/spinlock)](https://goreportcard.com/report/github.com/daniel-hutao/spinlock)
[![release](https://img.shields.io/github/release/daniel-hutao/spinlock.svg)](https://github.com/daniel-hutao/spinlock/releases/)

---

SpinLock is a spin lock implementation in Go with exponential backoff and adaptive spinning.

## Installation

To install the package, run:

```bash
go get github.com/daniel-hutao/spinlock
```

## Usage

Import the package and create a new SpinLock:

```go
import "github.com/daniel-hutao/spinlock"

var sl spinlock.SpinLock
```

Then you can use the `Lock` and `Unlock` methods:

```go
sl.Lock()
// critical section
sl.Unlock()
```

## Performance Testing

We have conducted performance tests to compare the efficiency of our SpinLock implementation with the standard Mutex in Go. The tests were run on a MacBook Pro with an Apple M1 chip, 16GB of RAM.

### SpinLock

```bash
$ go test -benchmem -run=^$ -bench ^BenchmarkSpinLock$ github.com/daniel-hutao/spinlock

goos: darwin
goarch: arm64
pkg: github.com/daniel-hutao/spinlock
=== RUN   BenchmarkSpinLock
BenchmarkSpinLock
BenchmarkSpinLock-10            111107053               10.80 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/daniel-hutao/spinlock        2.973s
```

### Mutex

```bash
$ go test -benchmem -run=^$ -bench ^BenchmarkMutex$ github.com/daniel-hutao/spinlock

goos: darwin
goarch: arm64
pkg: github.com/daniel-hutao/spinlock
=== RUN   BenchmarkMutex
BenchmarkMutex
BenchmarkMutex-10       10366155               115.5 ns/op             0 B/op          0 allocs/op
PASS
ok      github.com/daniel-hutao/spinlock        1.793s
```

### Conclusion

Based on our tests, the SpinLock implementation performs significantly better than the standard Mutex in Go on a MacBook Pro with an Apple M1 chip. Specifically, operations on SpinLock are approximately an order of magnitude faster than those on Mutex.
