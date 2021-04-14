```
goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/benchmarks/panic-recover
BenchmarkPanicRecover-8   	24628460	       276 ns/op	       0 B/op	       0 allocs/op
BenchmarkErrorFunc-8      	1000000000	         1.09 ns/op
PASS
ok  	github.com/xaionaro-go/benchmarks/panic-recover	8.243s
```

```
go test ./... -bench=PanicRecover -benchtime=5s -cpuprofile /tmp/cpu.profile
```
Resulted into:
```
xaionaro@void:~/go/src/github.com/xaionaro-go/benchmarks/panic-recover$ go tool pprof /tmp/cpu.profile
File: panic-recover.test
Type: cpu
Time: Apr 14, 2021 at 8:14pm (IST)
Duration: 7.22s, Total samples = 7.08s (98.10%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 5480ms, 77.40% of 7080ms total
Dropped 12 nodes (cum <= 35.40ms)
Showing top 10 nodes out of 32
      flat  flat%   sum%        cum   cum%
    2060ms 29.10% 29.10%     6240ms 88.14%  runtime.gopanic
     600ms  8.47% 37.57%     3020ms 42.66%  runtime.gentraceback
     600ms  8.47% 46.05%      630ms  8.90%  runtime.readvarintUnsafe
     490ms  6.92% 52.97%      800ms 11.30%  runtime.pcvalue
     470ms  6.64% 59.60%      530ms  7.49%  runtime.findfunc
     410ms  5.79% 65.40%      810ms 11.44%  runtime.addOneOpenDeferFrame.func1.1
     290ms  4.10% 69.49%      290ms  4.10%  runtime.step
     250ms  3.53% 73.02%     6610ms 93.36%  github.com/xaionaro-go/benchmarks/panic-recover.BenchmarkPanicRecover
     160ms  2.26% 75.28%      160ms  2.26%  runtime.gogo
     150ms  2.12% 77.40%      150ms  2.12%  runtime.gorecover
```

Here's the `runtime.gopanic` function implementation: https://github.com/golang/go/blob/585b52261c1b4e26b029616581ee0e891ad49183/src/runtime/panic.go#L951
