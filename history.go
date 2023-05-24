package historyToday

type DayHistoryEvent struct {
	Year      string `json:"year"`
	Date      string `json:"date"`
	Title     string `json:"title"`
	Festival  string `json:"festival"`
	Link      string `json:"link"`
	Type      string `json:"type"`
	Desc      string `json:"desc"`
	Cover     bool   `json:"cover"`
	Recommend bool   `json:"recommend"`
}

type DayHistory struct {
	Day              string            `json:"day"`
	DayHistoryEvents []DayHistoryEvent `json:"dayHistoryEvents"`
}

type MonthHistory struct {
	Month        string       `json:"month"`
	DayHistories []DayHistory `json:"dayHistories"`
}
