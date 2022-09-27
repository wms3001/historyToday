package historyToday

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestHistoryTody_GetHistoryTodayByMoth(t *testing.T) {
	historyToday := &HistoryTody{}
	historyToday.Type = "month"
	historyToday.Month = "09"
	resp := historyToday.GetHistoryToday()
	fmt.Println(resp)
	var monthEvent MonthEvent
	json.Unmarshal([]byte(resp.Data), &monthEvent)
}

func TestHistoryTody_GetHistoryTodayByDay(t *testing.T) {
	historyToday := &HistoryTody{}
	historyToday.Type = "day"
	historyToday.Month = "09"
	historyToday.Day = "28"
	resp := historyToday.GetHistoryDay()
	fmt.Println(resp)

}
