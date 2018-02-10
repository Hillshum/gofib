## Gofib

A Fibonacci generator API written in Go

### Things that still need to be added

* Potentially using math.Big to handle arbitrarly large numbers (currently an error is thrown if the sequence starts wrapping)
* More thorough testing of the route handling code (particularly mocking the `Fibonacci` function to ensure it is getting called)