package sematext

type AlertRuleScheduleWeekdayDto struct {
	Day       string                          `json:"day"`
	Index     int32                           `json:"index"`
	Intervals []AlertRuleScheduleTimeRangeDto `json:"intervals"`
	label     string                          `json:"end"`
	Type      string                          `json:"type"`
}
