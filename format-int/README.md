```
BenchmarkSprintfValue-4                 20000000                90.6 ns/op             3 B/op          1 allocs/op
BenchmarkSprintfInt-4                   20000000                90.8 ns/op             3 B/op          1 allocs/op
BenchmarkFormatInt-4                    50000000                30.6 ns/op             3 B/op          1 allocs/op
BenchmarkFormatFloat-4                  10000000               229 ns/op              35 B/op          2 allocs/op
BenchmarkFormatFloatConcat-4             5000000               337 ns/op              35 B/op          2 allocs/op
BenchmarkSprintfFloatConcat-4            5000000               304 ns/op              16 B/op          1 allocs/op
```
