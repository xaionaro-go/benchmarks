```
goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/benchmarks/defer
BenchmarkDefer-8         	30380739	        35.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkTwoDefers-8     	19684813	        59.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkNoDefer-8       	615593348	         1.90 ns/op	       0 B/op	       0 allocs/op
BenchmarkPseudoDefer-8   	621293611	         1.92 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/xaionaro-go/benchmarks/defer	5.122s
```
