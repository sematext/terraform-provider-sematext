package sematext

type ReportInfo struct {
	Addresses  []string `json:"addresses"`
	AppId      int64    `json:"appId"`
	EndDate    string   `json:"endDate"` // TODO ($date-time)
	Filters    string   `json:"filters"`
	ReportName string   `json:"reportName"`
	StartDate  string   `json:"startDate"` // TODO ($date-time)
	Subject    string   `json:"subject"`
	Text       string   `json:"text"`
}
