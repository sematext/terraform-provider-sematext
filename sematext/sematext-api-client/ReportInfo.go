package sematext

type ReportInfo struct {
	addresses  []string
	appId      int64
	endDate    string // TODO ($date-time)
	filters    string
	reportName string
	startDate  string // TODO ($date-time)
	subject    string
	text       string
}
