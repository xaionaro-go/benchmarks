```
$ go test -bench=. . -benchmem -benchtime 1s
...
BenchmarkModifyThroughPointer-4   	100000000	        10.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkModifyUsingCopy-4        	20000000	        65.4 ns/op	       0 B/op	       0 allocs/op
```
