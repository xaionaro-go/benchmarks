```
$ go test -bench=. . -benchmem
...
BenchmarkShortSliceCheck-4   	100000000	        22.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkBoolMapCheck-4      	10000000	       193 ns/op	       0 B/op	       0 allocs/op
BenchmarkStructMapCheck-4    	10000000	       189 ns/op	       0 B/op	       0 allocs/op
```
