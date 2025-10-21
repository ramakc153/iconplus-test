package practicaltest

import (
	"fmt"
)

type FoodList struct {
	Name  string
	Count int
}

type ProcessingData struct {
	UnitInduk        string
	UsagePercentage  float32
	ConsumptionPrice int
	Food             []FoodList
}

type ResponseData struct {
	Id               int
	UnitInduk        string
	RoomName         string
	UsagePercentage  float32
	ConsumptionPrice int
	Food             FoodList
}

var roomMap = make(map[string]*ProcessingData)

func getUniqueRoom() {
	for _, v := range bookingData {
		roomMap[v.RoomName] = &ProcessingData{}
	}

}

func Response() {
	// var Summary []ResponseData
	getUniqueRoom()
	getFoodPrice()

	for _, v := range bookingData {
		roomMap[v.RoomName].UnitInduk = v.OfficeName
		roomMap[v.RoomName].ConsumptionPrice += 0
		roomMap[v.RoomName].UsagePercentage += 1
		for _, food := range v.ListConsumption {
			var exist = false
			for i, existedFood := range roomMap[v.RoomName].Food {
				if existedFood.Name == food.Name {
					exist = true
					roomMap[v.RoomName].Food[i].Count += v.Participants
					fmt.Printf("on %s room: %s already exists\n", v.RoomName, food.Name)
				}
			}
			if !exist {
				roomMap[v.RoomName].Food = append(roomMap[v.RoomName].Food, FoodList{Name: food.Name, Count: v.Participants})
				exist = false
				fmt.Printf("on %s room: %s created\n", v.RoomName, food.Name)
			}
			fmt.Printf("current food list: %s on %v\n", v.RoomName, roomMap[v.RoomName].Food)
			roomMap[v.RoomName].ConsumptionPrice += FoodPrice[food.Name] * v.Participants
		}
		// // tempPrice := 0
		// // innermap := roomMap[v.RoomName].(map[string]any)
		// innermap["unitInduk"] = v.OfficeName
		// innermap["usagePercentage"] = 0
		// // roomMap[v.RoomName].(map[string]any)["unitInduk"] = v.OfficeName
		// // roomMap[v.RoomName].(map[string]any)["usagePercentage"] = 0
		// consumptionPrice, ok := innermap["consumptionPrice"].(int)
		// if !ok {
		// 	log.Println("initiate innermap[consumptionPrice]")
		// 	consumptionPrice = 0
		// }
		// // roomMap[v.RoomName].(map[string]any)["ConsumptionPrice"] = 0
		// consumptionList, ok := innermap["consumptionList"].([]map[string]any)
		// if !ok{
		// 	log.Println("initiate innermap[consumptionList] with type " )
		// 	consumptionList = make([]map[string]any, 0)
		// }
		// for _, food := range v.ListConsumption {
		// 	// get food count
		// 	currentFoodName, ok := consumptionList["name"]
		// 	if !ok {
		// 		currentFoodName = food.Name
		// 	}
		// 	// consumptionList = append(consumptionList, FoodList{Name: food.Name, Count: })
		// 	currentCount += v.Participants
		// 	innermap[food.Name] = currentCount

		// 	// get total price for each room
		// 	consumptionPrice += FoodPrice[food.Name] * v.Participants
		// 	innermap["consumptionPrice"] = consumptionPrice

		// }
	}

	for key, _ := range roomMap {
		roomMap[key].UsagePercentage = roomMap[key].UsagePercentage / float32(len(bookingData))
	}
	for key, room := range roomMap {
		fmt.Println("this is ", key, ": ", room)

	}
	// for id, v := range bookingData{
	// 	var room ResponseData
	// 	var totalPrice int
	// 	room.Id = id
	// 	room.UnitInduk = v.OfficeName
	// 	room.RoomName = v.RoomName

	// }

}
