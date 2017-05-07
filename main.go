/*

Run 'go test -bench=.' to get some benchmark results or visit
http://github.com/IanTayler/c-go-benchmarks.git for more information.

Copyright ® 2017 Ian G. Tayler <ian.g.tayler@gmail.com>

Distribute according to the LICENSE.
*/
package main

/*
#cgo CFLAGS: -I./clib -O3
#include "basics.c"
*/
import "C"
import "fmt"

/*MAXTHREADS defines how many threads we're going to spawn for each function. */
const MAXTHREADS = 8

/*****************************
 * RECURSIVE FIBONACCI TESTS *
 *****************************/

/*CSimplRecFib wraps the C SimplRecFib to feed a Go channel. */
func CSimplRecFib(channel chan uint32, n int) {
	res := uint32(C.SimplRecFib(C.int(n)))
	channel <- res
}

/*CStdintRecFib wraps the C StdintRecFib to feed a Go channel. */
func CStdintRecFib(channel chan uint32, n int) {
	res := uint32(C.StdintRecFib(C.int(n)))
	channel <- res
}

/*GoBaseRecFib is the Go function we use as a base recursive fibonacci function. */
func GoBaseRecFib(n int) uint32 {
	if n < 2 {
		return 1
	}
	return GoBaseRecFib(n-1) + GoBaseRecFib(n-2)
}

/*GoRecFib wraps the base function (GoBaseRecFib) to connect it to a channel. */
func GoRecFib(channel chan uint32, n int) {
	res := GoBaseRecFib(n)
	channel <- res
}

/* Our type of function, which communicates through a channel and takes an int as input. */
type concurrFunc func(chan uint32, int)

/*ConcWrap does the work of wrapping a normal function that feeds a
Go channel and run it concurrently several times. */
func ConcWrap(procFunct concurrFunc) {
	channel := make(chan uint32, MAXTHREADS)
	for i := 0; i < MAXTHREADS; i++ {
		go procFunct(channel, i*5)
	}
	for i := 0; i < MAXTHREADS; i++ {
		<-channel
	}
}

/*ConstConcWrap wraps a concurrFunc but, unlike ConcWrap, each run
of the function takes the same input, which is the second argument of
ConstConcWrap. */
func ConstConcWrap(procFunct concurrFunc, inputn int) {
	channel := make(chan uint32, MAXTHREADS)
	for i := 0; i < MAXTHREADS; i++ {
		go procFunct(channel, inputn)
	}
	for i := 0; i < MAXTHREADS; i++ {
		<-channel
	}
}

func main() {
	fmt.Println("Running as main. You won't get much from this.\nTry running 'go test -bench=.' instead.")
}
