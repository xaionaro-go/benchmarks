```
goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/benchmarks/chan
cpu: AMD Ryzen 9 5900X 12-Core Processor            
BenchmarkChan0-24                 	 9622108	       137.6 ns/op
BenchmarkChan1-24                 	13006867	        94.60 ns/op
BenchmarkChan2-24                 	15598532	        78.56 ns/op
BenchmarkChan1024-24              	37977878	        32.02 ns/op
BenchmarkChan1024WithSelect-24    	17415333	        75.30 ns/op
BenchmarkChan1PutGet-24           	46017588	        26.66 ns/op	       0 B/op	       0 allocs/op
BenchmarkRingBufferPutGet-24      	41865273	        40.88 ns/op	      25 B/op	       0 allocs/op
PASS
ok  	github.com/xaionaro-go/benchmarks/chan	9.795s
```