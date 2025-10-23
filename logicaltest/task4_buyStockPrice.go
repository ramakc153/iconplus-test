package logicaltest

func BuyStockPrice(arr []int) int {
	var maxProfit int
	var currentIndex int
	for i := 0; i < len(arr)-1; i++ {
		currentProfit := arr[i+1] - arr[i]
		if currentProfit > maxProfit {
			maxProfit = currentProfit
			currentIndex = i
		}
	}
	if maxProfit == 0 {
		return 0 // return 0 means no profit
	} else {
		return arr[currentIndex]
	}
}
