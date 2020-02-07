package sematext

type SavedQuery struct {
	AlertRule         AlertRule       `json:"alertRule"`
	AllowModification bool            `json:"allowModification"` // TODO example: false readOnly: true
	ApplicationId     string          `json:"applicationId"`
	ApplicationName   string          `json:"applicationName"`  // TODO readOnly: true
	ApplicationToken  string          `json:"applicationToken"` // TODO readOnly: true
	CreatorEmail      string          `json:"creatorEmail"`     // TODO readOnly: true
	Id                string          `json:"id"`               // TODO readOnly: true
	LabelColor        string          `json:"labelColor"`
	LogseneAlertType  map[int]string  `json:"logseneAlertType"` // TODO Check datatype
	OwnerEmail        string          `json:"ownerEmail"`       // TODO readOnly: true
	QueryName         string          `json:"queryName"`
	QueryString       string          `json:"queryString"`
	UserPermissions   UserPermissions `json:"userPermissions"`
}
