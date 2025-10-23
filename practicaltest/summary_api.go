package practicaltest

import "fmt"

type SummaryAPI struct {
	OfficeName       string     `json:"officeName"`
	RoomName         string     `json:"roomName"`
	UsagePercentage  float32    `json:"usagePercentage"`
	ConsumptionPrice float32    `json:"consumptionPrice"`
	ConsumptionList  []FoodList `json:"consumptionList"`
}

var is100 = 0.0

func ConstructSummaryApi() []SummaryAPI {
	var SummaryAPIData []SummaryAPI
	for room, v := range roomMap {
		var data SummaryAPI
		data.OfficeName = v.UnitInduk
		data.RoomName = room
		is100 += float64(v.UsagePercentage)
		data.UsagePercentage = v.UsagePercentage
		data.ConsumptionPrice = float32(v.ConsumptionPrice)
		data.ConsumptionList = v.Food
		SummaryAPIData = append(SummaryAPIData, data)
	}
	fmt.Printf("this is total percentage hours across all rooms: %f\n ", is100)
	return SummaryAPIData
}
