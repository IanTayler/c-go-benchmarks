// Copyright ® 2017 Ian G. Tayler <ian.g.tayler@gmail.com>
// Distribute according to the LICENSE.
package main

/*
#cgo CFLAGS: -I./clib -O3
#include "basics.c"
*/
import "C"
import "fmt"

const MAXTHREADS = 8

/*****************************
 * RECURSIVE FIBONACCI TESTS *
 *****************************/

/* Wrap the C functions to return to a Go channel */
func CSimplRecFib(channel chan uint32, n int) {
	res := uint32(C.SimplRecFib(C.int(n)))
	channel <- res
}

func CStdintRecFib(channel chan uint32, n int) {
	res := uint32(C.StdintRecFib(C.int(n)))
	channel <- res
}

/* The base Go recursive fibonacci */
func GoBaseRecFib(n int) uint32 {
	if n < 2 {
		return 1
	} else {
		return GoBaseRecFib(n-1) + GoBaseRecFib(n-2)
	}
}

/* Wrap the base function to connect it to a channel */
func GoRecFib(channel chan uint32, n int) {
	res := GoBaseRecFib(n)
	channel <- res
}

/* Our type of function, which communicates through a channel and takes an int as input */
type concurrFunc func(chan uint32, int)

func ConcWrap(procFunct concurrFunc) {
	channel := make(chan uint32, MAXTHREADS)
	for i := 0; i < MAXTHREADS; i++ {
		go procFunct(channel, i*5)
	}
	for i := 0; i < MAXTHREADS; i++ {
		<-channel
	}
}

func main() {
	fmt.Println("Running as main")
}
