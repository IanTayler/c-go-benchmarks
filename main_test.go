// Copyright ® 2017 Ian G. Tayler <ian.g.tayler@gmail.com>
// Distribute according to the LICENSE.
package main

import "testing"

func BenchmarkGoRecFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcWrap(GoRecFib)
	}
}

func BenchmarkCSimplRecFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcWrap(CSimplRecFib)
	}
}

func BenchmarkCStdintRecFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcWrap(CStdintRecFib)
	}
}
