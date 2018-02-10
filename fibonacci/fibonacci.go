/*
Package fibonacci provides the Fibonacci sequence
An iterative Fibbonacci sequence generator
Much of this is from https://www.dotnetperls.com/fibonacci-go
*/
package fibonacci

// Fibonacci Returns a slice with all Fibonacci numbers up to n
func Fibonacci(n int) []uint {

	fibs := []uint{}

	var (
		a uint
		b uint = 1
	)
	// Iterate until desired position in sequence.
	for i := 0; i < n; i++ {
		// Use temporary variable to swap values.
		temp := a
		a = b
		b = temp + a
		fibs = append(fibs, a)
	}
	return fibs
}
