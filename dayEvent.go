package historyToday

type Event struct {
	Year      string `json:"year"`
	Title     string `json:"title"`
	Festival  string `json:"festival"`
	Link      string `json:"link"`
	Type      string `json:"type"`
	Desc      string `json:"desc"`
	Cover     bool   `json:"cover"`
	Recommend bool   `json:"recommend"`
}

type DayEvent struct {
	Day    string  `json:"day"`
	Events []Event `json:"events"`
}
