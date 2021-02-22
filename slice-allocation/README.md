```
goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/benchmarks/slice-allocation
BenchmarkAllocatedSliceIfaceBoxing-8           	95357548	        33.2 ns/op	      32 B/op	       1 allocs/op
BenchmarkNilSliceIfaceBoxing-8                 	1000000000	         2.18 ns/op	       0 B/op	       0 allocs/op
BenchmarkMakeSlice8IfaceBoxing-8               	78462126	        55.5 ns/op	 144.20 MB/s	      40 B/op	       2 allocs/op
BenchmarkMakeSliceUnsafeString8IfaceBoxing-8   	88007739	        50.6 ns/op	 157.99 MB/s	      24 B/op	       2 allocs/op
BenchmarkArray8IfaceBoxing-8                   	60843411	        56.4 ns/op	 141.83 MB/s	      40 B/op	       2 allocs/op
BenchmarkArrayUnsafeString8IfaceBoxing-8       	76773693	        50.9 ns/op	 157.23 MB/s	      24 B/op	       2 allocs/op
BenchmarkAllocatedSlice-8                      	1000000000	         0.268 ns/op	       0 B/op	       0 allocs/op
BenchmarkNilSlice-8                            	1000000000	         0.262 ns/op	       0 B/op	       0 allocs/op
BenchmarkMakeSlice8-8                          	1000000000	         0.576 ns/op	13888.09 MB/s	       0 B/op	       0 allocs/op
BenchmarkArray8-8                              	1000000000	         0.300 ns/op	26650.38 MB/s	       0 B/op	       0 allocs/op
BenchmarkArrayUnsafeString8-8                  	1000000000	         0.703 ns/op	11381.98 MB/s	       0 B/op	       0 allocs/op
PASS
ok  	github.com/xaionaro-go/benchmarks/slice-allocation	24.253s
```