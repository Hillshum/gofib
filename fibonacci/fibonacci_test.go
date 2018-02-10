package fibonacci

import (
	"testing"
)

func compareFibs(a []uint, b []uint) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestFibonacci13(t *testing.T) {
	fibs := Fibonacci(13)
	expected := []uint{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144}

	if !compareFibs(fibs, expected) {
		t.Errorf("Expected %v, got %v", expected, fibs)
	}
}

func TestFibonacci0(t *testing.T) {
	fibs := Fibonacci(0)
	expected := []uint{0}

	if !compareFibs(fibs, expected) {
		t.Errorf("Expected %v, got %v", expected, fibs)
	}
}
