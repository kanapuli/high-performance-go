package factorialtest

import "testing"

//STARTFAC OMIT
func factorial(n int) int {
	switch n {
	case 0:
		return 1
	default:
		return n * factorial(n-1)
	}
}

//ENDFAC OMIT

//STARTBEN OMIT
func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factorial(2000)
	}
}

//ENDBEN OMIT
