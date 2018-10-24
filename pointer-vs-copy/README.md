```
go test -bench=. . -benchmem -benchtime 2
...
BenchmarkModifyThroughPointer-4   	200000000	        12.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkModifyUsingCopy-4        	50000000	        64.3 ns/op	       0 B/op	       0 allocs/op
```
