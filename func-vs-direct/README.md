```
$ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/benchmarks/func-vs-direct
BenchmarkDirect-4       100000000               19.1 ns/op
BenchmarkFunc-4         100000000               19.0 ns/op
```
