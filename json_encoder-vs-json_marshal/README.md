```
goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/benchmarks/json_encoder-vs-json_marshal
BenchmarkJsonEncoder4-8                      	 2000000	       992 ns/op	     144 B/op	       3 allocs/op
BenchmarkJsonEncoder256-8                    	  100000	     23012 ns/op	     656 B/op	       3 allocs/op
BenchmarkJsonEncoder65536-8                  	     300	   4180997 ns/op	  143273 B/op	       3 allocs/op
BenchmarkJsonEncoderWithBufferReuse4-8       	 2000000	      1063 ns/op	      32 B/op	       1 allocs/op
BenchmarkJsonEncoderWithBufferReuse256-8     	  100000	     13779 ns/op	      32 B/op	       1 allocs/op
BenchmarkJsonEncoderWithBufferReuse65536-8   	     300	   3664973 ns/op	    1474 B/op	       1 allocs/op
BenchmarkJsonMarshal4-8                      	 2000000	       952 ns/op	      48 B/op	       2 allocs/op
BenchmarkJsonMarshal256-8                    	  100000	     15635 ns/op	     608 B/op	       2 allocs/op
BenchmarkJsonMarshal65536-8                  	     500	   4067791 ns/op	  140504 B/op	       2 allocs/op
PASS
ok  	github.com/xaionaro-go/benchmarks/json_encoder-vs-json_marshal	20.401s
```
