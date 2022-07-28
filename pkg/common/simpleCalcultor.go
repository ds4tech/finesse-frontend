package common

import (
	"math"
	"math/big"
)

func Sum(x int64, y int64) int64 {
	return x + y
}

func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func IsPrime(n int) (result bool) {
	for i := 2; i <= n; i++ {
		if big.NewInt(int64(i)).ProbablyPrime(20) {
			result = true
		} else {
			result = false
		}
	}
	return result
}

func Log(n float64) (result float64) {
	result = math.Log(n)
	return result
}
