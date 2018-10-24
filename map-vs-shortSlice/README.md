```
$ go test -bench=. . -benchmem
...
BenchmarkShortSliceCheck_iterations2-4          200000000                9.65 ns/op            0 B/op          0 allocs/op
BenchmarkBoolMapCheck_iterations2-4             10000000               168 ns/op               0 B/op          0 allocs/op
BenchmarkStructMapCheck_iterations2-4           10000000               174 ns/op               0 B/op          0 allocs/op
BenchmarkShortSliceCheck_iterations10-4         30000000                41.7 ns/op             0 B/op          0 allocs/op
BenchmarkBoolMapCheck_iterations10-4            10000000               234 ns/op               0 B/op          0 allocs/op
BenchmarkStructMapCheck_iterations10-4          10000000               232 ns/op               0 B/op          0 allocs/op
BenchmarkShortSliceCheck_iterations100-4         3000000               555 ns/op               0 B/op          0 allocs/op
BenchmarkBoolMapCheck_iterations100-4            1000000              1246 ns/op               0 B/op          0 allocs/op
BenchmarkStructMapCheck_iterations100-4          1000000              1125 ns/op               0 B/op          0 allocs/op
```
