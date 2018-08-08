package factorialtest

import "testing"

func factorial(n int) int {
	switch n {
	case 0:
		return 1
	default:
		return n * factorial(n-1)
	}
}

func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factorial(20)
	}
}
