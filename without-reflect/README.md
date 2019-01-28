```
goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/benchmarks/without-reflect
BenchmarkCheckMethodWithoutReflect-4     	100000000	        20.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkCheckMethodWithReflect-4        	 3000000	       493 ns/op	     168 B/op	       6 allocs/op
BenchmarkGetFieldValueWithoutReflect-4   	1000000000	         2.23 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetFieldValueWithReflect-4      	10000000	       130 ns/op	      16 B/op	       2 allocs/op
```
