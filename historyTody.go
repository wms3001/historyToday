package historyToday

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"github.com/wms3001/goTool"
	"github.com/wms3001/http"
)

var url string = "https://baike.baidu.com"

type HistoryTody struct {
	Type  string `json:"type"`
	Month string `json:"date"`
	Day   string `json:"day"`
}

func (historyTody *HistoryTody) GetHistoryToday() MonthHistory {
	m := map[string]string{
		"Content-Type": "application/json",
	}
	a := map[string]string{
		"Accept": "application/json",
		//"Accept": "text/plain",
	}
	var wHttp = http.Http{
		url + "/cms/home/eventsOnHistory/" + historyTody.Month + ".json",
		"",
		"",
		m,
		a,
		"",
		10,
		"",
		[]byte{},
	}
	resp := wHttp.Get()
	mp := make(map[string]interface{})
	json.Unmarshal([]byte(resp.Data), &mp)

	var history = MonthHistory{}
	goTool := &goTool.GoTool{}
	for k, v := range mp {
		history.Month = k
		sv := goTool.Strval(v)
		mp1 := make(map[string]interface{})
		json.Unmarshal([]byte(sv), &mp1)
		var dayhistories = []DayHistory{}
		for kk, vv := range mp1 {
			var dayHistory = DayHistory{}
			dayHistory.Day = kk
			svv := goTool.Strval(vv)
			var ar2 []interface{}
			json.Unmarshal([]byte(svv), &ar2)
			var dayHistoryEvents = []DayHistoryEvent{}
			for _, vvv := range ar2 {
				var dayHistoryEvent DayHistoryEvent
				mapstructure.Decode(vvv, &dayHistoryEvent)
				dayHistoryEvent.Date = dayHistoryEvent.Year + "." + kk[0:2] + "." + kk[2:]
				dayHistoryEvents = append(dayHistoryEvents, dayHistoryEvent)
			}
			dayHistory.DayHistoryEvents = dayHistoryEvents
			dayhistories = append(dayhistories, dayHistory)
		}
		history.DayHistories = dayhistories
	}
	return history
}

func (historyTody *HistoryTody) GetHistoryDay() DayHistory {
	var dayHistory DayHistory
	if len(historyTody.Day) < 2 {
		historyTody.Day = "0" + historyTody.Day
	}
	resp := historyTody.GetHistoryToday()
	for _, v := range resp.DayHistories {
		if v.Day == historyTody.Month+historyTody.Day {
			dayHistory = v
		}
	}
	return dayHistory
}
