goos: linux
goarch: amd64
BenchmarkCompositeKeyMap-8              	300000000	         4.24 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompositeWithStringKeyMap-8    	50000000	        36.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkUint64KeyMap-8                 	300000000	         3.91 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringKeyMap-8                 	300000000	         5.40 ns/op	       0 B/op	       0 allocs/op
BenchmarkCompositeKeyToStringKeyMap-8   	50000000	        26.1 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	_/home/xaionaro/benchmarks/composite-map-key	8.659s
