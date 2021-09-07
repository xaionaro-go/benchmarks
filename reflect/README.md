```
$ go test -bench=. ./...
goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/benchmarks/reflect
BenchmarkReflectTypeOf-8   	1000000000	         0.488 ns/op
PASS
ok  	github.com/xaionaro-go/benchmarks/reflect	0.539s
```

```
xaionaro@void:~/go/src/github.com/xaionaro-go/benchmarks/reflect$ go test -gcflags=-l -bench=. ./...
goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/benchmarks/reflect
BenchmarkReflectTypeOf-8   	826415682	         1.47 ns/op
PASS
ok  	github.com/xaionaro-go/benchmarks/reflect	1.362s
```
