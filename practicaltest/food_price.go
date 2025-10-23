package practicaltest

var FoodPrice = make(map[string]int)

func getFoodPrice() {
	for _, v := range konsumsiData {
		FoodPrice[v.Name] = v.MaxPrice
	}
}
