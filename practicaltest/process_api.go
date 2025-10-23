package practicaltest

type FoodList struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type ProcessingData struct {
	Id               int
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
				}
			}
			if !isExist {
				roomMap[v.RoomName].Food = append(roomMap[v.RoomName].Food, FoodList{Name: food.Name, Count: v.Participants})
				isExist = false
			}
			roomMap[v.RoomName].ConsumptionPrice += FoodPrice[food.Name] * v.Participants
		}
	}

	for key, _ := range roomMap {
		roomMap[key].UsagePercentage = roomMap[key].UsagePercentage / float32(TotalUse)
	}

}
