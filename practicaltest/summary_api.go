package practicaltest

type SummaryAPI struct {
	OfficeName       string     `json:"officeName"`
	RoomName         string     `json:"roomName"`
	UsagePercentage  float32    `json:"usagePercentage"`
	ConsumptionPrice float32    `json:"consumptionPrice"`
	ConsumptionList  []FoodList `json:"consumptionList"`
}

func ConstructSummaryApi() []SummaryAPI {
	var SummaryAPIData []SummaryAPI
	for room, v := range roomMap {
		var data SummaryAPI
		data.OfficeName = v.UnitInduk
		data.RoomName = room
		data.UsagePercentage = v.UsagePercentage
		data.ConsumptionPrice = float32(v.ConsumptionPrice)
		data.ConsumptionList = v.Food
		SummaryAPIData = append(SummaryAPIData, data)
	}
	return SummaryAPIData
}
