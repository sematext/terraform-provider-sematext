package sematext

type Invitation struct {
	app           sematext.App
	apps          []sematext.App
	id            int64  // TODO readOnly: true
	inviteDate    string // TODO readOnly: true, ISO 8063 or long as ($date-time)?
	inviteStatus  string // TODO readOnly: true Enum:[ PENDING, ACCEPTED, CANCELLED, DECLINED ]
	inviteeEmail  string // example: guest@sematext.com
	inviteeRole   string // TODO example: DEMO Enum:[ SUPER_USER, OWNER, ADMIN, USER, DEMO, ANONYMOUS ]
	inviteeStatus string // TODO example: ACTIVE readOnly: true Enum:[ INACTIVE, ACTIVE ]
	inviterEmail  string // TODO readOnly: true
	uuid          string // TODO readOnly: true
}
