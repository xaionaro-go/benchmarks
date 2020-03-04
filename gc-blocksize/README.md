```
goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/benchmarks/gc-blocksize
BenchmarkBlockSize/blockSize_1-8         	    3085	    205063 ns/op	       9 B/op	       0 allocs/op
BenchmarkBlockSize/blockSize_1024-8      	    2479	    206715 ns/op	       7 B/op	       0 allocs/op
BenchmarkBlockSize/blockSize_32768-8     	    2748	    207603 ns/op	       7 B/op	       0 allocs/op
BenchmarkBlockSize/blockSize_10485760-8  	    2167	    273000 ns/op	       1 B/op	       0 allocs/op
BenchmarkBlockSize/blockSize_20971520-8  	    1933	    305989 ns/op	       2 B/op	       0 allocs/op
BenchmarkBlockSize/blockSize_41943040-8  	    1864	    316502 ns/op	       5 B/op	       0 allocs/op
PASS
ok  	github.com/xaionaro-go/benchmarks/gc-blocksize	27.056s
```