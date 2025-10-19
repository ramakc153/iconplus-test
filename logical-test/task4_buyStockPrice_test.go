package logicaltest

import "testing"

type StockPriceStruct struct {
	Name  string
	Input []int
	Want  int
}

func TestBuyStockPrice(t *testing.T) {
	testTable := []StockPriceStruct{
		{"stock1", []int{10, 9, 6, 5, 15}, 5},
		{"stock2", []int{7, 8, 3, 10, 8}, 3},
		{"stock3", []int{5, 12, 11, 12, 10}, 5},
		{"stock4", []int{7, 18, 27, 10, 29}, 10},
		{"stock5", []int{20, 17, 15, 14, 10}, 0},
	}

	for _, tt := range testTable {
		result := BuyStockPrice(tt.Input)
		t.Run(tt.Name, func(t *testing.T) {
			if result != tt.Want {
				t.Errorf("expected %v, found %v", tt.Want, result)

			}
		})
	}
}
