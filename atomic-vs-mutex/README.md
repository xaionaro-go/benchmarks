```
BenchmarkWriteAtomic_1goroutine-4        	200000000	         8.34 ns/op
BenchmarkWriteAtomic_2goroutine-4        	100000000	        51.5 ns/op
BenchmarkWriteAtomic_4goroutine-4        	20000000	       100 ns/op
BenchmarkWriteAtomic_16goroutine-4       	 3000000	       411 ns/op
BenchmarkWriteAtomic_64goroutine-4       	 1000000	      1666 ns/op
BenchmarkWriteAtomic_256goroutine-4      	  200000	      6621 ns/op
BenchmarkReadAtomic_1goroutine-4         	2000000000	         1.95 ns/op
BenchmarkReadAtomic_2goroutine-4         	1000000000	         2.03 ns/op
BenchmarkReadAtomic_4goroutine-4         	1000000000	         3.02 ns/op
BenchmarkReadAtomic_16goroutine-4        	100000000	        13.5 ns/op
BenchmarkReadAtomic_64goroutine-4        	30000000	        52.7 ns/op
BenchmarkReadAtomic_256goroutine-4       	10000000	       212 ns/op
BenchmarkWriteCGOAtomic_1goroutine-4     	20000000	        83.5 ns/op
BenchmarkWriteCGOAtomic_2goroutine-4     	20000000	       117 ns/op
BenchmarkWriteCGOAtomic_4goroutine-4     	10000000	       177 ns/op
BenchmarkWriteCGOAtomic_16goroutine-4    	 2000000	       742 ns/op
BenchmarkWriteCGOAtomic_64goroutine-4    	  500000	      2901 ns/op
BenchmarkWriteCGOAtomic_256goroutine-4   	  200000	     11428 ns/op
BenchmarkWriteIdeal_1goroutine-4         	1000000000	         2.27 ns/op
BenchmarkWriteIdeal_2goroutine-4         	100000000	        15.8 ns/op
BenchmarkWriteIdeal_4goroutine-4         	50000000	        22.6 ns/op
BenchmarkWriteIdeal_16goroutine-4        	20000000	        90.3 ns/op
BenchmarkWriteIdeal_64goroutine-4        	 5000000	       358 ns/op
BenchmarkWriteIdeal_256goroutine-4       	 1000000	      1441 ns/op
BenchmarkReadIdeal_1goroutine-4          	2000000000	         1.95 ns/op
BenchmarkReadIdeal_2goroutine-4          	1000000000	         2.00 ns/op
BenchmarkReadIdeal_4goroutine-4          	500000000	         3.36 ns/op
BenchmarkReadIdeal_16goroutine-4         	100000000	        13.6 ns/op
BenchmarkReadIdeal_64goroutine-4         	30000000	        52.8 ns/op
BenchmarkReadIdeal_256goroutine-4        	10000000	       211 ns/op
BenchmarkWriteMutex_1goroutine-4         	100000000	        18.4 ns/op
BenchmarkWriteMutex_2goroutine-4         	30000000	        56.9 ns/op
BenchmarkWriteMutex_4goroutine-4         	 5000000	       289 ns/op
BenchmarkWriteMutex_16goroutine-4        	 1000000	      2091 ns/op
BenchmarkWriteMutex_64goroutine-4        	  200000	      8679 ns/op
BenchmarkWriteMutex_256goroutine-4       	   50000	     22566 ns/op
BenchmarkReadMutex_1goroutine-4          	30000000	        55.1 ns/op
BenchmarkReadMutex_2goroutine-4          	10000000	       160 ns/op
BenchmarkReadMutex_4goroutine-4          	 3000000	       514 ns/op
BenchmarkReadMutex_16goroutine-4         	  500000	      3007 ns/op
BenchmarkReadMutex_64goroutine-4         	  100000	     12818 ns/op
BenchmarkReadMutex_256goroutine-4        	   50000	     39210 ns/op
```
