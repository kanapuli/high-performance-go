package fib

import (
	"testing"
)

func Fib(n int) int {
	switch n {
	case 0:
		return 1
	case 1:
		return 1
		//	case 2:
		//		return 1
	default:
		return Fib(n-1) + Fib(n-2)
	}
}

func BenchmarkFib(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		Fib(20)
	}
}
