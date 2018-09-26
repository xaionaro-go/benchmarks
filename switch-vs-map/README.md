```
$ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/benchmarks/switch-vs-map
BenchmarkMap2-4         100000000               16.4 ns/op
BenchmarkMap64-4        50000000                29.5 ns/op
BenchmarkSwitch2-4      200000000                8.58 ns/op
BenchmarkSwitch64-4     100000000               10.5 ns/op
```
