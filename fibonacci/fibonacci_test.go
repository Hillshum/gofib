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
	fibs, err := Fibonacci(13)
	expected := []uint{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144}

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !compareFibs(fibs, expected) {
		t.Errorf("Expected %v, got %v", expected, fibs)
	}
}

func TestFibonacciZero(t *testing.T) {
	fibs, err := Fibonacci(0)
	expected := []uint{0}

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !compareFibs(fibs, expected) {
		t.Errorf("Expected %v, got %v", expected, fibs)
	}
}

func TestFibonacciOverflow(t *testing.T) {
	_, err := Fibonacci(94)

	if err == nil {
		t.Error("Expected error, got none")
	}

}

func TestFibonacciNegative(t *testing.T) {
	_, err := Fibonacci(-1)

	if err == nil {
		t.Error("Expected error, got none")
	}

}
