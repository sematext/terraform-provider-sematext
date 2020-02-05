package sematext

type SavedQuery struct {
	alertRule         AlertRule
	allowModification bool // TODO example: false readOnly: true
	applicationId     string
	applicationName   string // TODO readOnly: true
	applicationToken  string // TODO readOnly: true
	creatorEmail      string // TODO readOnly: true
	id                string // TODO readOnly: true
	labelColor        string
	logseneAlertType  map[int]string // TODO Check datatype
	ownerEmail        string         // TODO readOnly: true
	queryName         string
	queryString       string
	userPermissions   UserPermissions
}
