# golang 历史上的今天
## 简介
获取历史上当天发生的大事
## 使用
```
go get github.com/wms3001/historyToday
```
## 实例
1. 按月份获取
```go
historyToday := &HistoryTody{}
historyToday.Type = "month"
historyToday.Month = "09"
resp := historyToday.GetHistoryToday()
```
2. 按天获取
```go
historyToday := &HistoryTody{}
historyToday.Type = "day"
historyToday.Month = "09"
historyToday.Day = "28"
resp := historyToday.GetHistoryDay()
```
