package historyToday

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"github.com/wms3001/goCommon"
	"github.com/wms3001/goTool"
	"github.com/wms3001/http"
)

var url string = "https://baike.baidu.com"

type HistoryTody struct {
	Type  string `json:"type"`
	Month string `json:"date"`
	Day   string `json:"day"`
}

func (historyTody *HistoryTody) GetHistoryToday() *goCommon.Resp {
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
	var monthEvent = MonthEvent{}
	goTool := &goTool.GoTool{}
	for k, v := range mp {
		monthEvent.Month = k
		sv := goTool.Strval(v)
		mp1 := make(map[string]interface{})
		json.Unmarshal([]byte(sv), &mp1)
		//fmt.Println(mp1)
		for kk, vv := range mp1 {
			var dayEvent DayEvent
			dayEvent.Day = kk
			svv := goTool.Strval(vv)
			var ar2 []interface{}
			json.Unmarshal([]byte(svv), &ar2)
			for _, vvv := range ar2 {
				var event Event
				mapstructure.Decode(vvv, &event)
				event.Year = event.Year + kk
				dayEvent.Events = append(dayEvent.Events, event)
			}
			monthEvent.DayEvent = append(monthEvent.DayEvent, dayEvent)
		}
	}
	r, _ := json.Marshal(monthEvent)
	resp.Data = string(r)
	return resp
}

func (historyTody *HistoryTody) GetHistoryDay() *goCommon.Resp {
	var monthEvent MonthEvent
	res := &goCommon.Resp{}
	res.Code = 1
	res.Message = "success"
	resp := historyTody.GetHistoryToday()
	json.Unmarshal([]byte(resp.Data), &monthEvent)
	for _, v := range monthEvent.DayEvent {
		if v.Day == historyTody.Month+historyTody.Day {
			bv, _ := json.Marshal(v)
			res.Data = string(bv)
			break
		}
	}
	return res
}
