The code of this benchmark is copied from: [https://stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go/23857998#23857998](https://stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go/23857998#23857998)
```
goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/benchmarks/string-builder
BenchmarkConcat-4          	  200000	     17832 ns/op	  103608 B/op	       1 allocs/op
BenchmarkBuffer-4          	20000000	        10.6 ns/op	       3 B/op	       0 allocs/op
BenchmarkCopy-4            	30000000	         5.79 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringBuilder-4   	30000000	         3.63 ns/op	       5 B/op	       0 allocs/op
PASS
ok  	github.com/xaionaro-go/benchmarks/string-builder	4.186s
```
