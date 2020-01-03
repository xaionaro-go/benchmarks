```
xaionaro@xaionaro-mbp mariomac % go test ./... -bench=. -benchtime 1ns
goos: darwin
goarch: amd64
pkg: github.com/xaionaro-go/benchmarks/vs-java/mariomac
BenchmarkMariomacAllocateSlice-16              1          14259324 ns/op        400007264 B/op         2 allocs/op
BenchmarkMariomacFillRandom-16                 1        1509445405 ns/op               0 B/op          0 allocs/op
BenchmarkMariomacQuickSort-16                  1        10680925474 ns/op              0 B/op          0 allocs/op
PASS
ok      github.com/xaionaro-go/benchmarks/vs-java/mariomac      78.203s
```
