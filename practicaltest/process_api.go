package practicaltest

import (
	"fmt"
)

type FoodList struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type ProcessingData struct {
	UnitInduk        string
	UsagePercentage  float32
	ConsumptionPrice int
	Food             []FoodList
}

var roomMap = make(map[string]*ProcessingData)

func getUniqueRoom() {
	for _, v := range bookingData {
		roomMap[v.RoomName] = &ProcessingData{}
	}

}

var TotalUse = 0

func Response() {
	getUniqueRoom()
	getFoodPrice()

	for _, v := range bookingData {
		roomMap[v.RoomName].UnitInduk = v.OfficeName
		roomMap[v.RoomName].ConsumptionPrice += 0
		roomMap[v.RoomName].UsagePercentage += float32(parsetime(v.StartTime, v.EndTime))
		TotalUse += parsetime(v.StartTime, v.EndTime)
		for _, food := range v.ListConsumption {
			var isExist = false
			for i, existedFood := range roomMap[v.RoomName].Food {
				if existedFood.Name == food.Name {
					isExist = true
					roomMap[v.RoomName].Food[i].Count += v.Participants
					fmt.Printf("on %s room: %s already exists\n", v.RoomName, food.Name)
				}
			}
			if !isExist {
				roomMap[v.RoomName].Food = append(roomMap[v.RoomName].Food, FoodList{Name: food.Name, Count: v.Participants})
				isExist = false
				fmt.Printf("on %s room: %s created\n", v.RoomName, food.Name)
			}
			fmt.Printf("current food list: %s on %v\n", v.RoomName, roomMap[v.RoomName].Food)
			roomMap[v.RoomName].ConsumptionPrice += FoodPrice[food.Name] * v.Participants
		}
	}

	for key, _ := range roomMap {
		roomMap[key].UsagePercentage = roomMap[key].UsagePercentage / float32(TotalUse)
	}
	for key, room := range roomMap {
		fmt.Println("this is ", key, ": ", room)

	}

}
