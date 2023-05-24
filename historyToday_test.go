package historyToday

import (
	"encoding/json"
	"log"
	"testing"
)

func TestHistoryTody_GetHistoryTodayByMoth(t *testing.T) {
	historyToday := &HistoryTody{}
	historyToday.Type = "month"
	historyToday.Month = "09"
	resp := historyToday.GetHistoryToday()
	re, _ := json.Marshal(resp)
	log.Println(string(re))
}

func TestHistoryTody_GetHistoryTodayByDay(t *testing.T) {
	historyToday := &HistoryTody{}
	historyToday.Type = "day"
	historyToday.Month = "09"
	historyToday.Day = "28"
	resp := historyToday.GetHistoryDay()
	re, _ := json.Marshal(resp)
	log.Println(string(re))

}
