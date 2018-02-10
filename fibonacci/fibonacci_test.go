package fibonacci

import (
	"reflect"
	"testing"
)

func TestFibonacci(t *testing.T) {
	fibs := Fibonacci(2)
	expected := {1, 2}

	if !reflect.DeepEqual(fibs, {1, 2}){
		t.Error("Expected [1, 2], got ", fibs)
	}
}
