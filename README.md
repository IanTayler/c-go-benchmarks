# c-go-benchmarks

## Description

This repository contains some benchmarks I ran to study the possibility of using Go's concurrency to parallely run procedural C code. Most of the benchmarks out there ([example](https://benchmarksgame.alioth.debian.org/u64q/compare.php?lang=go&lang2=gcc)) show that C code can _sometimes_ be several times faster than Go code. That said, writing safe concurrent code in Go is much easier, so if we could mix the two it would be a great success for concurrent programs that need to go as fast as possible.

The question is: can we find any real improvements in example cases? It's perfectly possible that that isn't the case, for several different reasons (for example, there may be some optimizations that the go compiler manages to do on its concurrent code but that can't be done when some of the code is written in C). This is what these benchmarks are made to test.

## Tests

### First test: Recursive fibonacci

As a first test, we wrote three simple recursive functions that calculate the n-th fibonacci number. One is written in Go, and called `GoRecFib`. The other two are written in C (wrapped in Go to spit their return values to a Go channel), called `CSimplRecFib` and `CStdintRecFib`. The difference between them is that the first one uses `int` as the return type while the second one uses the `stdint.h` type `uint32_t`.

`MAXTHREADS` is a constant that defines how many different threads we are going to spawn to run our functions. It also sets the amount of times the basic recursive functions are going to be run.

Here are the results:
```
###### FIRST RUN     ######
###### MAXTHREADS: 8 ######
BenchmarkGoRecFib-2        	      10	 126382177 ns/op #Go
BenchmarkCSimplRecFib-2    	      30	  46614938 ns/op #C
BenchmarkCStdintRecFib-2   	      30	  46631506 ns/op #C stdint.h
PASS
ok  	github.com/IanTayler/c-go-benchmarks	4.335s

###### SECOND RUN    ######
###### MAXTHREADS: 8 ######
BenchmarkGoRecFib-2        	      10	 126112046 ns/op #Go
BenchmarkCSimplRecFib-2    	      30	  47690399 ns/op #C
BenchmarkCStdintRecFib-2   	      30	  48016443 ns/op #C stdint.h
PASS
ok  	github.com/IanTayler/c-go-benchmarks	4.418s

###### THIRD RUN     ######
###### MAXTHREADS: 4 ######
BenchmarkGoRecFib-2        	  100000	     14952 ns/op #Go
BenchmarkCSimplRecFib-2    	  200000	      9948 ns/op #C
BenchmarkCStdintRecFib-2   	  200000	      9558 ns/op #C stdint.h
PASS
ok  	github.com/IanTayler/c-go-benchmarks	5.759s

###### FOURTH RUN    ######
###### MAXTHREADS: 2 ######
BenchmarkGoRecFib-2        	 1000000	      1778 ns/op #Go
BenchmarkCSimplRecFib-2    	 1000000	      2197 ns/op #C
BenchmarkCStdintRecFib-2   	 1000000	      2184 ns/op #C stdint.h
PASS
ok  	github.com/IanTayler/c-go-benchmarks	6.237s
```

As you can see, in this case there was an actual improvement in speed with the C-based code. The difference is larger the larger the amount of threads is, and Go even performs better when we only run the functions twice. If I had to advance an explanation of that fact, I'd guess it is because a lower MAXTHREADS means we're actually running the functions with a smaller input (because the input we give to the function is in the range `(0 .. MAXTHREADS*5)`), and then the linear increase in speed that C gives us gets overshadowed by the constant overhead of having to convert between Go types and C types and, possibly, the overhead of calling a C function from Go.

To test this, let's run the test with `MAXTHREADS = 2` but passing arguments to the fibonacci functions in the range `(30 .. 30+MAXTHREADS)`.
```
BenchmarkGoRecFib-2        	     100	  19034983 ns/op #Go
BenchmarkCSimplRecFib-2    	     200	   6956307 ns/op #C
BenchmarkCStdintRecFib-2   	     200	   6974458 ns/op #C stdint.h
PASS
ok  	github.com/IanTayler/c-go-benchmarks	6.180s
```
As expected, using C functions as a base is much faster even with `MAXTHREADS = 2`, as long as the input number is large enough (i.e. as long as we actually spend a large percentage of the time running the base functions, and not the wrapper Go code).

### Second test: fibonacci with constant input
Following a suggestion by Dave Cheney (check out [his blog](https://dave.cheney.net/)), I added tests where we always take the same input for the fibonacci function that's being run concurrently. This way, we can see how 'heavy' a function needs to be so that the linear improvement of using C outweights the constant overhead of calling it through Go.

Here are the results.
```
###### FIB(1) ######
BenchmarkConstant1GoRecFib-2         	  200000	      6403 ns/op #Go
BenchmarkConstant1CSimplRecFib-2     	  200000	      8009 ns/op #C
BenchmarkConstant1CStdintRecFib-2    	  200000	      8175 ns/op #C stdint.h

###### FIB(2) ######
BenchmarkConstant2GoRecFib-2         	  300000	      6343 ns/op #Go
BenchmarkConstant2CSimplRecFib-2     	  200000	      8097 ns/op #C
BenchmarkConstant2CStdintRecFib-2    	  200000	      8112 ns/op #C stdint.h

###### FIB(5) ######
BenchmarkConstant5GoRecFib-2         	  200000	      6843 ns/op #Go
BenchmarkConstant5CSimplRecFib-2     	  200000	      8126 ns/op #C
BenchmarkConstant5CStdintRecFib-2    	  200000	      8249 ns/op #C stdint.h

###### FIB(7) ######
BenchmarkConstant7GoRecFib-2         	  200000	      8144 ns/op #Go
BenchmarkConstant7CSimplRecFib-2     	  200000	      8619 ns/op #C
BenchmarkConstant7CStdintRecFib-2    	  200000	      8612 ns/op #C stdint.h

###### FIB(10) ######
BenchmarkConstant10GoRecFib-2        	  100000	     14273 ns/op #Go
BenchmarkConstant10CSimplRecFib-2    	  200000	     11171 ns/op #C
BenchmarkConstant10CStdintRecFib-2   	  200000	     11223 ns/op #C stdint.h

###### FIB(20) ######
BenchmarkConstant20GoRecFib-2        	    5000	    335232 ns/op #Go
BenchmarkConstant20CSimplRecFib-2    	   10000	    155335 ns/op #C
BenchmarkConstant20CStdintRecFib-2   	   10000	    155101 ns/op #C stdint.h

###### FIB(30) ######
BenchmarkConstant30GoRecFib-2        	      30	  45240073 ns/op #Go
BenchmarkConstant30CSimplRecFib-2    	     100	  18227397 ns/op #C
BenchmarkConstant30CStdintRecFib-2   	     100	  17661102 ns/op #C stdint.h

###### FIB(40) ######
BenchmarkConstant40GoRecFib-2        	       1	7457899354 ns/op #Go
BenchmarkConstant40CSimplRecFib-2    	       1	3032271971 ns/op #C
BenchmarkConstant40CStdintRecFib-2   	       1	2887700742 ns/op #C stdint.h
PASS
ok  	github.com/IanTayler/c-go-benchmarks	50.608s
```
Remember a recursive fibonacci function runs in constant time for input <= 2. These results show that even relatively fast functions (rec-fib(10), for example, which is quite fast for _I-need-to-lower-this-code-to-C_ standards) can have a small speed-up using C.

The clear information this gives us is that if you're planning on calling a short, fast function hundreds of times that's **NOT** the function to lower to C level. Doing that would actually make your program slower. The cases where you *should* consider using cgo is when you have a very heavy function. In those cases we're seeing x2 and x3 speedups.

## Preliminary conclusions

We will run many more tests. We have to, before we can advance any real conclusions. As of now, it looks as if this type of C/Go interaction can actually be profitable. The only tests we have run so far have shown a great increase of speed when using procedural C functions as a base when we have _heavy_ functions that take a lot of time to complete inside the C code. The constant overhead caused by calling through cgo and having to convert between language-types means it's not worth it if what we're talking about is a fast function we're planning on calling several times.

So, as a rule of thumb: if you minimize the amount of time where you're passing the ball from C to Go and from Go to C, and actually run non-stop for a long time in C, it's possible to see an interesting increase in speed using procedural C code.

Another observation is that using `stdint.h` types instead of the standard C types doesn't alter the performance significantly.

## Author

Copyright ® 2017 Ian G. Tayler <ian.g.tayler@gmail.com>
