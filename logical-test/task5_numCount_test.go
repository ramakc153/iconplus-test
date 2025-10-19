package logicaltest

import "testing"

type NumCountStruct struct {
	Name  string
	Input []string
	Want  int
}

func TestNumCount(t *testing.T) {
	testTable := []NumCountStruct{
		{"this should be 4", []string{"2", "h", "6", "u", "y", "t", "7", "j", "y", "h", "8"}, 4},
		{"this should be 5", []string{"b", "7", "h", "6", "h", "k", "i", "5", "g", "7", "8"}, 5},
		{"this should be 7", []string{"7", "b", "8", "5", "6", "9", "n", "f", "y", "6", "9"}, 7},
		{"this should be 7", []string{"u", "h", "b", "n", "7", "6", "5", "1", "g", "7", "9"}, 6},
	}

	for _, tt := range testTable {
		result := NumCount(tt.Input)
		t.Run(tt.Name, func(t *testing.T) {
			if result != tt.Want {
				t.Errorf("expected %v, found %v", tt.Want, result)
			}
		})
	}
}
