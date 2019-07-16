goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/benchmarks/shuffle-vs-perm
BenchmarkShuffleShort-8    	20000000	       102 ns/op	       0 B/op	       0 allocs/op
BenchmarkShuffleMedium-8   	  200000	      7660 ns/op	       0 B/op	       0 allocs/op
BenchmarkShuffleLong-8     	    1000	   2041611 ns/op	       0 B/op	       0 allocs/op
BenchmarkPermShort-8       	10000000	       152 ns/op	      32 B/op	       1 allocs/op
BenchmarkPermMedium-8      	  200000	      9947 ns/op	    2048 B/op	       1 allocs/op
BenchmarkPermLong-8        	     500	   2765165 ns/op	  524290 B/op	       1 allocs/op
PASS
ok  	github.com/xaionaro-go/benchmarks/shuffle-vs-perm	11.401s
