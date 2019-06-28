```
goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/benchmarks/interface-method
BenchmarkPtrMethodDirect-8         	2000000000	         0.34 ns/op	       0 B/op	       0 allocs/op
BenchmarkMethodDirect-8            	2000000000	         0.34 ns/op	       0 B/op	       0 allocs/op
BenchmarkPtrMethodViaInterface-8   	1000000000	         1.99 ns/op	       0 B/op	       0 allocs/op
BenchmarkMethodViaInterface-8      	1000000000	         3.16 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/xaionaro-go/benchmarks/interface-method	7.052s
```
