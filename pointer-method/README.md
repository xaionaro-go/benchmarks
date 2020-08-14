```
goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/benchmarks/pointer-method
Benchmark/noPointer-8         	1000000000	         0.526 ns/op	       0 B/op	       0 allocs/op
Benchmark/pointer-8           	85375690	        12.7 ns/op	       8 B/op	       1 allocs/op
PASS
ok  	github.com/xaionaro-go/benchmarks/pointer-method	1.681s
```
