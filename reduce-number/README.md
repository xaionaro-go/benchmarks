```
goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/benchmarks/reduce-number
BenchmarkIntegerReductionSimple32-8   	216929269	         5.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkIntegerReductionLemire32-8   	544098690	         2.14 ns/op	       0 B/op	       0 allocs/op
BenchmarkIntegerReductionOkunev32-8   	374748714	         3.21 ns/op	       0 B/op	       0 allocs/op
BenchmarkIntegerReductionSimple64-8   	100000000	        10.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkIntegerReductionOkunev64-8   	312919128	         3.45 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/xaionaro-go/benchmarks/reduce-number	8.333s
```
