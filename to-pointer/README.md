```
$ go test -bench=. -benchmem
goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/benchmarks/to-pointer
BenchmarkToPointerSimple-4      100000000               18.6 ns/op             8 B/op          1 allocs/op
BenchmarkToPointerSlice-4       100000000               18.9 ns/op             8 B/op          1 allocs/op
BenchmarkToPointerMake-4        100000000               18.8 ns/op             8 B/op          1 allocs/op
```
