goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/benchmarks/defer
BenchmarkDefer_1goroutine-4     	30000000	        53.5 ns/op
BenchmarkNoDefer_1goroutine-4   	2000000000	         1.94 ns/op
PASS
ok  	github.com/xaionaro-go/benchmarks/defer	5.743s
