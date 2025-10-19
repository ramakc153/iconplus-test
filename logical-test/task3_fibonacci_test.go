package logicaltest

import (
	"slices"
	"testing"
)

type FibonacciStruct struct {
	Name  string
	Input int
	Want  []int
}

func TestFibonacci(t *testing.T) {
	testTable := []FibonacciStruct{
		{Name: "fibonacci 3", Input: 3, Want: []int{0, 1, 1, 2, 3}},
		{Name: "fibonacci 4", Input: 4, Want: []int{0, 1, 1, 2, 3, 5}},
		{Name: "fibonacci 5", Input: 5, Want: []int{0, 1, 1, 2, 3, 5, 8}},
		{Name: "fibonacci 6", Input: 6, Want: []int{0, 1, 1, 2, 3, 5, 8, 13}},
	}

	for _, tt := range testTable {
		t.Run(tt.Name, func(t *testing.T) {
			result := Fibonacci(tt.Input)
			if !slices.Equal(result, tt.Want) {
				t.Errorf("expected %v, found %v", tt.Want, result)
			}
		})
	}

}
