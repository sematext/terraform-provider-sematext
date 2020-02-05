package sematext

type NotificationIntegration struct {
	applicability   string // TODO Enum: [ NONE, USE_ALWAYS ]
	createDate      string // TODO ($date-time) ISO8063 or long?
	createdByOwner  bool
	creatorId       int64
	id              int64
	integrationType string // TODO Enum:[ PAGER_DUTY, NAGIOS, WEB_HOOKS, WEB_HOOKS_TEMPLATE, HIP_CHAT, EMAIL_LIST, TEMPORARY_EMAIL_LIST ]
	name            string
	params          map[int]string //TODO Check key appropriate type?
	state           string         // TODO Enum: [ ACTIVE, DISABLED, DELETED ]
	userId          int64
}
