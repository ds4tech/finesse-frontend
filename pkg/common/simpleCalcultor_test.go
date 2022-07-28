package common

import (
	"testing"
)

func TestSum(t *testing.T) {
	result := Sum(int64(3), int64(4))
	expecting := int64(7)
	if result != expecting {
		t.Errorf("expecting %d, got %d", expecting, result)
	}
}
func TestSqrt(t *testing.T) {
	result := Sqrt(9)
	expecting := float64(3)
	if result != expecting {
		t.Errorf("expecting %v, got %v", expecting, result)
	}
	result = Sqrt(4)
	expecting = 2
	if result != expecting {
		t.Errorf("expecting %v, got %v", expecting, result)
	}
}
func TestFactorial(t *testing.T) {
	result := Factorial(4)
	expecting := uint64(24)
	if result != expecting {
		t.Errorf("expecting %d, got %d", expecting, result)
	}
}
func TestLog(t *testing.T) {
	result := Log(1)
	expecting := float64(0)
	if result != expecting {
		t.Errorf("expecting %v, got %v", expecting, result)
	}
}
