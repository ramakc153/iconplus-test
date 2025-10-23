package logicaltest

import (
	"fmt"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	for v := range 100 {
		n := v + 1
		result := FizzBuzz(n)
		var expected string
		switch {
		case n%5 == 0 && n%3 == 0:
			expected = "FizzBuzz"
		case n%3 == 0:
			expected = "Fizz"
		case n%5 == 0:
			expected = "Buzz"
		default:
			expected = fmt.Sprintf("%d", n)
		}

		if result != expected {
			t.Errorf("case : %d expected %s, found %s", n, expected, result)
		}
	}
}
