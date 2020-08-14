goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/benchmarks/gc-blocksize
BenchmarkBlockSizeAllocateAndGC/blockSize_1-8         	   57853	    112647 ns/op	      16 B/op	       1 allocs/op
BenchmarkBlockSizeAllocateAndGC/blockSize_1024-8      	   52795	    114308 ns/op	    1024 B/op	       1 allocs/op
BenchmarkBlockSizeAllocateAndGC/blockSize_32768-8     	   49084	    123642 ns/op	   32768 B/op	       1 allocs/op
BenchmarkBlockSizeAllocateAndGC/blockSize_10485760-8  	    2535	   2559597 ns/op	10485775 B/op	       1 allocs/op
BenchmarkBlockSizeAllocateAndGC/blockSize_20971520-8  	    1230	   4608001 ns/op	20971541 B/op	       1 allocs/op
BenchmarkBlockSizeAllocateAndGC/blockSize_41943040-8  	    1155	   4461621 ns/op	41943060 B/op	       1 allocs/op
BenchmarkBlockSizeAllocate/blockSize_1-8              	477077799	        12.6 ns/op	       1 B/op	       1 allocs/op
BenchmarkBlockSizeAllocate/blockSize_1024-8           	34407218	       163 ns/op	    1024 B/op	       1 allocs/op
BenchmarkBlockSizeAllocate/blockSize_32768-8          	 1675790	      3620 ns/op	   32768 B/op	       1 allocs/op
BenchmarkBlockSizeAllocate/blockSize_10485760-8       	    2732	   2135901 ns/op	10485767 B/op	       1 allocs/op
BenchmarkBlockSizeAllocate/blockSize_20971520-8       	    1294	   4614378 ns/op	20971530 B/op	       1 allocs/op
BenchmarkBlockSizeAllocate/blockSize_41943040-8       	    1384	   4254765 ns/op	41943050 B/op	       1 allocs/op
BenchmarkBlockSizeGC/blockSize_1-8                    	   50448	    117935 ns/op	       3 B/op	       0 allocs/op
BenchmarkBlockSizeGC/blockSize_1024-8                 	   51127	    118572 ns/op	       3 B/op	       0 allocs/op
BenchmarkBlockSizeGC/blockSize_32768-8                	   50878	    118907 ns/op	       3 B/op	       0 allocs/op
BenchmarkBlockSizeGC/blockSize_10485760-8             	   33643	    179520 ns/op	       0 B/op	       0 allocs/op
BenchmarkBlockSizeGC/blockSize_20971520-8             	   31728	    203932 ns/op	       0 B/op	       0 allocs/op
BenchmarkBlockSizeGC/blockSize_41943040-8             	   34375	    174879 ns/op	       1 B/op	       0 allocs/op
PASS
ok  	github.com/xaionaro-go/benchmarks/gc-blocksize	626.747s
