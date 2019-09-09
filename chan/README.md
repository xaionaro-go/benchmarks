```
goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/benchmarks/chan
BenchmarkChan1-8                	 5229620	       239 ns/op	       0 B/op	       0 allocs/op
BenchmarkChan2-8                	 7536854	       188 ns/op	       0 B/op	       0 allocs/op
BenchmarkChan1024-8             	11750379	        98.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkChan1024WithSelect-8   	 5879042	       178 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/xaionaro-go/benchmarks/chan	5.591s
```