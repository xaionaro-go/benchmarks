```
goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/benchmarks/preallocated-map
BenchmarkMap/itemsCount0-8   						 2779698	      4.34 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap/itemsCount1-8   						  847792	      13.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap/itemsCount8-8   						   91653	       132 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap/itemsCount256-8 						     622	     17033 ns/op	   12613 B/op	      28 allocs/op
BenchmarkMap/itemsCount65536-8  				       	       3	   4448147 ns/op	 3069530 B/op	    2249 allocs/op
BenchmarkMap/itemsCount1048576-8				       	       1	 128568155 ns/op	49898200 B/op	   38472 allocs/op
BenchmarkPreallocatedMap/itemsCount8/preallocated0-8         		   53624	       198 ns/op	       0 B/op	       0 allocs/op
BenchmarkPreallocatedMap/itemsCount8/preallocated8-8         		   59730	       200 ns/op	       0 B/op	       0 allocs/op
BenchmarkPreallocatedMap/itemsCount8/preallocated12-8        		   45272	       277 ns/op	     160 B/op	       1 allocs/op
BenchmarkPreallocatedMap/itemsCount8/preallocated16-8        		   39642	       303 ns/op	     320 B/op	       1 allocs/op
BenchmarkPreallocatedMap/itemsCount8/preallocated24-8        		   41491	       308 ns/op	     320 B/op	       1 allocs/op
BenchmarkPreallocatedMap/itemsCount8/preallocated64-8        		   27554	       455 ns/op	    1440 B/op	       2 allocs/op
BenchmarkPreallocatedMap/itemsCount256/preallocated0-8       		     514	     23303 ns/op	   12611 B/op	      28 allocs/op
BenchmarkPreallocatedMap/itemsCount256/preallocated256-8     		    1196	      9955 ns/op	    6215 B/op	       3 allocs/op
BenchmarkPreallocatedMap/itemsCount256/preallocated384-8     		    1245	      9920 ns/op	    6217 B/op	       4 allocs/op
BenchmarkPreallocatedMap/itemsCount256/preallocated512-8     		    1868	      6654 ns/op	   10913 B/op	       2 allocs/op
BenchmarkPreallocatedMap/itemsCount256/preallocated768-8     		    1911	      6386 ns/op	   10913 B/op	       2 allocs/op
BenchmarkPreallocatedMap/itemsCount256/preallocated2048-8    		    1567	      7927 ns/op	   49184 B/op	       2 allocs/op
BenchmarkPreallocatedMap/itemsCount65536/preallocated0-8     		       3	   4414268 ns/op	 3078192 B/op	    2282 allocs/op
BenchmarkPreallocatedMap/itemsCount65536/preallocated65536-8 		       5	   2242873 ns/op	 1400888 B/op	      13 allocs/op
BenchmarkPreallocatedMap/itemsCount65536/preallocated98304-8 		       5	   2216141 ns/op	 1400888 B/op	      13 allocs/op
BenchmarkPreallocatedMap/itemsCount65536/preallocated131072-8         	       6	   2413569 ns/op	 2785474 B/op	       7 allocs/op
BenchmarkPreallocatedMap/itemsCount65536/preallocated196608-8         	       6	   2141580 ns/op	 2785517 B/op	       7 allocs/op
BenchmarkPreallocatedMap/itemsCount65536/preallocated524288-8         	       3	   4532747 ns/op	11141152 B/op	       2 allocs/op
BenchmarkPreallocatedMap/itemsCount1048576/preallocated0-8            	       1	 137364433 ns/op	49888664 B/op	   38358 allocs/op
BenchmarkPreallocatedMap/itemsCount1048576/preallocated1048576-8      	       1	  91558949 ns/op	22496568 B/op	      21 allocs/op
BenchmarkPreallocatedMap/itemsCount1048576/preallocated1572864-8      	       1	  93014944 ns/op	22496664 B/op	      22 allocs/op
BenchmarkPreallocatedMap/itemsCount1048576/preallocated2097152-8      	       1	  87000212 ns/op	44566584 B/op	      11 allocs/op
BenchmarkPreallocatedMap/itemsCount1048576/preallocated3145728-8      	       1	  83927656 ns/op	44566584 B/op	      11 allocs/op
BenchmarkPreallocatedMap/itemsCount1048576/preallocated8388608-8      	       1	  97125452 ns/op       178257952 B/op	       2 allocs/op
PASS
ok  	github.com/xaionaro-go/benchmarks/preallocated-map	1.123s
```
