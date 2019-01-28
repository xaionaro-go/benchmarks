```
goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/benchmarks/value-into-interface
BenchmarkSliceToInterfaceSimple-4   	 1000000	      1512 ns/op
BenchmarkSliceToInterfaceTricky-4   	30000000	        42.5 ns/op
BenchmarkUint64ToInterface-4        	2000000000	         1.97 ns/op
BenchmarkUint64ToUint64-4           	2000000000	         1.93 ns/op
```
