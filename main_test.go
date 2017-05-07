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

func BenchmarkConstant1GoRecFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConstConcWrap(GoRecFib, 1)
	}
}

func BenchmarkConstant1CSimplRecFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConstConcWrap(CSimplRecFib, 1)
	}
}

func BenchmarkConstant1CStdintRecFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConstConcWrap(CStdintRecFib, 1)
	}
}

func BenchmarkConstant2GoRecFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConstConcWrap(GoRecFib, 2)
	}
}

func BenchmarkConstant2CSimplRecFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConstConcWrap(CSimplRecFib, 2)
	}
}

func BenchmarkConstant2CStdintRecFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConstConcWrap(CStdintRecFib, 2)
	}
}

func BenchmarkConstant5GoRecFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConstConcWrap(GoRecFib, 5)
	}
}

func BenchmarkConstant5CSimplRecFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConstConcWrap(CSimplRecFib, 5)
	}
}

func BenchmarkConstant5CStdintRecFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConstConcWrap(CStdintRecFib, 5)
	}
}

func BenchmarkConstant7GoRecFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConstConcWrap(GoRecFib, 7)
	}
}

func BenchmarkConstant7CSimplRecFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConstConcWrap(CSimplRecFib, 7)
	}
}

func BenchmarkConstant7CStdintRecFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConstConcWrap(CStdintRecFib, 7)
	}
}

func BenchmarkConstant10GoRecFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConstConcWrap(GoRecFib, 10)
	}
}

func BenchmarkConstant10CSimplRecFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConstConcWrap(CSimplRecFib, 10)
	}
}

func BenchmarkConstant10CStdintRecFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConstConcWrap(CStdintRecFib, 10)
	}
}

func BenchmarkConstant20GoRecFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConstConcWrap(GoRecFib, 20)
	}
}

func BenchmarkConstant20CSimplRecFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConstConcWrap(CSimplRecFib, 20)
	}
}

func BenchmarkConstant20CStdintRecFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConstConcWrap(CStdintRecFib, 20)
	}
}

func BenchmarkConstant30GoRecFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConstConcWrap(GoRecFib, 30)
	}
}

func BenchmarkConstant30CSimplRecFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConstConcWrap(CSimplRecFib, 30)
	}
}

func BenchmarkConstant30CStdintRecFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConstConcWrap(CStdintRecFib, 30)
	}
}

func BenchmarkConstant40GoRecFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConstConcWrap(GoRecFib, 40)
	}
}

func BenchmarkConstant40CSimplRecFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConstConcWrap(CSimplRecFib, 40)
	}
}

func BenchmarkConstant40CStdintRecFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConstConcWrap(CStdintRecFib, 40)
	}
}
