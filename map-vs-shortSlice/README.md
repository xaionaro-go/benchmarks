```
$ go test -bench=. . -benchmem
...
BenchmarkShortSliceCheck_iterations2-4          100000000               10.3 ns/op             0 B/op          0 allocs/op
BenchmarkBoolMapCheck_iterations2-4             10000000               179 ns/op               0 B/op          0 allocs/op
BenchmarkStructMapCheck_iterations2-4           10000000               227 ns/op               0 B/op          0 allocs/op
BenchmarkShortSliceCheck_iterations10-4         30000000                53.9 ns/op             0 B/op          0 allocs/op
BenchmarkBoolMapCheck_iterations10-4             5000000               278 ns/op               0 B/op          0 allocs/op
BenchmarkStructMapCheck_iterations10-4           5000000               256 ns/op               0 B/op          0 allocs/op
BenchmarkShortSliceCheck_iterations100-4         3000000               566 ns/op               0 B/op          0 allocs/op
BenchmarkBoolMapCheck_iterations100-4            1000000              1392 ns/op               0 B/op          0 allocs/op
BenchmarkStructMapCheck_iterations100-4          1000000              1279 ns/op               0 B/op          0 allocs/op
BenchmarkShortSliceCheckString_iterations5-4    30000000                53.3 ns/op             0 B/op          0 allocs/op
BenchmarkStructMapCheckString_iterations5-4      5000000               259 ns/op               0 B/op          0 allocs/op
```
