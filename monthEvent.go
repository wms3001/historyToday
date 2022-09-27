package historyToday

type MonthEvent struct {
	Month    string     `json:"month"`
	DayEvent []DayEvent `json:"dayEvent"`
}
