```
$ go test -bench=. . -benchmem -benchtime 2s
goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/benchmarks/boolMap-vs-structMap
BenchmarkBoolMap-4            10         210457146 ns/op        54897817 B/op      38431 allocs/op
BenchmarkStructMap-4          10         214032472 ns/op        49826865 B/op      38393 allocs/op
```
