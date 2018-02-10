/*
Package fibonacci provides the Fibonacci sequence
An iterative Fibbonacci sequence generator
Much of this is from https://www.dotnetperls.com/fibonacci-go
*/
package fibonacci

import (
	"errors"
)

// Fibonacci Returns a slice with all Fibonacci numbers up to n
func Fibonacci(n int) ([]uint, error) {

	fibs := []uint{0}

	var (
		a uint
		b uint = 1
	)

	if n == 0 {
		return fibs, nil
	}
	// Iterate until desired position in sequence.
	for i := 0; i < n-1; i++ {
		// Use temporary variable to swap values.
		temp := a
		a = b
		b = temp + a
		if b < a {
			return nil, errors.New("Input to Fibonacci too high. Overflow resulted")
		}
		fibs = append(fibs, a)
	}
	return fibs, nil
}
