```
goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/benchmarks/slice-allocation
BenchmarkAllocatedSliceIfaceBoxing-8           	30361909	        39.7 ns/op	      32 B/op	       1 allocs/op
BenchmarkNilSliceIfaceBoxing-8                 	531446181	         2.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkMakeSlice8IfaceBoxing-8               	16989213	        64.2 ns/op	 124.68 MB/s	      40 B/op	       2 allocs/op
BenchmarkMakeSliceUnsafeString8IfaceBoxing-8   	26745309	        84.7 ns/op	  94.41 MB/s	      24 B/op	       2 allocs/op
BenchmarkArray8IfaceBoxing-8                   	10359627	       106 ns/op	  75.60 MB/s	      40 B/op	       2 allocs/op
BenchmarkArrayUnsafeString8IfaceBoxing-8       	16159585	        68.4 ns/op	 116.99 MB/s	      24 B/op	       2 allocs/op
BenchmarkAllocatedSlice-8                      	1000000000	         0.459 ns/op	       0 B/op	       0 allocs/op
BenchmarkNilSlice-8                            	1000000000	         0.483 ns/op	       0 B/op	       0 allocs/op
BenchmarkMakeSlice8-8                          	1000000000	         1.16 ns/op	6910.32 MB/s	       0 B/op	       0 allocs/op
BenchmarkArray8-8                              	1000000000	         0.499 ns/op	16018.67 MB/s	       0 B/op	       0 allocs/op
BenchmarkArrayUnsafeString8-8                  	913250084	         1.27 ns/op	6310.25 MB/s	       0 B/op	       0 allocs/op
PASS
ok  	github.com/xaionaro-go/benchmarks/slice-allocation	13.600s
```