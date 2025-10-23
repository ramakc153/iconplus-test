package practicaltest

type SummaryAPI struct {
	Id               int        `json:"id"`
	OfficeName       string     `json:"officeName"`
	RoomName         string     `json:"roomName"`
	UsagePercentage  float32    `json:"usagePercentage"`
	ConsumptionPrice int        `json:"consumptionPrice"`
	ConsumptionList  []FoodList `json:"consumptionList"`
}

func ConstructSummaryApi() []SummaryAPI {
	var id = 0
	var SummaryAPIData []SummaryAPI
	for room, v := range roomMap {
		var data SummaryAPI
		data.Id = id
		data.OfficeName = v.UnitInduk
		data.RoomName = room
		data.UsagePercentage = v.UsagePercentage
		data.ConsumptionPrice = v.ConsumptionPrice
		data.ConsumptionList = v.Food
		SummaryAPIData = append(SummaryAPIData, data)
		id++
	}
	return SummaryAPIData
}
