# c-go-benchmarks

## Description

This repository contains some benchmarks I ran to study the possibility of using Go's concurrency to parallely run procedural C code. Most of the benchmarks out there ([example](https://benchmarksgame.alioth.debian.org/u64q/compare.php?lang=go&lang2=gcc)) show that C code can _sometimes_ be several times faster than Go code. That said, writing safe concurrent code in Go is much easier, so if we could mix the two it would be a great success for concurrent programs that need to go as fast as possible.

The question is: can we find any real improvements in example cases? It's perfectly possible that that isn't the case, for several different reasons (for example, there may be some optimizations that the go compiler manages to do on its concurrent code but that can't be done when some of the code is written in C). This is what these benchmarks are made to test.

## Tests

### First test: Recursive fibonacci

As a first test, we wrote three simple recursive functions that calculate the n-th fibonacci number. One is written in Go, and called `GoRecFib`. The other two are written in C (wrapped in Go to spit to channel), called `CSimplRecFib` and `CStdintRecFib`. The difference is that the first one uses `int` as a type while the second one uses the `stdint.h` type `uint32_t`.

`MAXTHREADS` is a constant that sets in how many different threads are we going to run each function. It also sets the amount of times the basic recursive functions are going to be run.

Here are the results:
```
###### FIRST RUN     ######
###### MAXTHREADS: 8 ######
BenchmarkGoRecFib-2        	      10	 126382177 ns/op
BenchmarkCSimplRecFib-2    	      30	  46614938 ns/op
BenchmarkCStdintRecFib-2   	      30	  46631506 ns/op
PASS
ok  	github.com/IanTayler/c-go-benchmarks	4.335s

###### SECOND RUN    ######
###### MAXTHREADS: 8 ######
BenchmarkGoRecFib-2        	      10	 126112046 ns/op
BenchmarkCSimplRecFib-2    	      30	  47690399 ns/op
BenchmarkCStdintRecFib-2   	      30	  48016443 ns/op
PASS
ok  	github.com/IanTayler/c-go-benchmarks	4.418s

###### THIRD RUN     ######
###### MAXTHREADS: 4 ######
BenchmarkGoRecFib-2        	  100000	     14952 ns/op
BenchmarkCSimplRecFib-2    	  200000	      9948 ns/op
BenchmarkCStdintRecFib-2   	  200000	      9558 ns/op
PASS
ok  	github.com/IanTayler/c-go-benchmarks	5.759s

###### FOURTH RUN    ######
###### MAXTHREADS: 2 ######
BenchmarkGoRecFib-2        	 1000000	      1778 ns/op
BenchmarkCSimplRecFib-2    	 1000000	      2197 ns/op
BenchmarkCStdintRecFib-2   	 1000000	      2184 ns/op
PASS
ok  	github.com/IanTayler/c-go-benchmarks	6.237s
```

As you can see, in this case there was an actual improvement in speed with the C-based code. The difference is larger the larger the amount of threads is, and Go even performs better when we only run the functions twice. If I had to advance an explanation of that fact, I'd guess it is because a lower MAXTHREADS means we're actually running the functions with a smaller input (because the input we give to the function is in the range `(0 .. MAXTHREADS*5)`), and then the difference linear increase in speed that C gives us gets overshadowed by the constant overhead of having to convert between Go types and C types and, possibly, the overhead of calling a C function from Go.

To test this, let's run the test with `MAXTHREADS = 2` but passing arguments in the range `(30 .. 30+MAXTHREADS)`.
```
BenchmarkGoRecFib-2        	     100	  19034983 ns/op
BenchmarkCSimplRecFib-2    	     200	   6956307 ns/op
BenchmarkCStdintRecFib-2   	     200	   6974458 ns/op
PASS
ok  	github.com/IanTayler/c-go-benchmarks	6.180s
```
As expected, using C functions as a base is much faster even with `MAXTHREADS = 2`, as long as the input number is large enough (i.e. as long as we actually spend a large percentage of the time running the base functions, and not the wrapper Go code).

## Preliminary conclusions

We will run many more tests. We have to, before we can advance any real conclusions. As of now, it looks as if this type of C/Go interaction can actually be profitable. The only test we have run so far has shown a great increase of speed when using procedural C functions as a base.

Another observation is that using `stdint.h` types instead of the standard C types doesn't alter the performance significantly.

## Author

Copyright ® 2017 Ian G. Tayler <ian.g.tayler@gmail.com>
