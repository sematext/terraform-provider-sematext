package sematext

type AlertRuleScheduleWeekdayDto struct {
	day       string
	index     int32
	intervals []sematext.AlertRuleScheduleTimeRangeDto
	label     string
	Type      string `json:"type"`
}
