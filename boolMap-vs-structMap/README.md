```
$ go test -bench=. . -benchmem -benchtime 2s
...
BenchmarkBoolMap-4             5         209179986 ns/op        54897150 B/op      38423 allocs/op
BenchmarkStructMap-4           5         210535947 ns/op        49830084 B/op      38387 allocs/op
```
