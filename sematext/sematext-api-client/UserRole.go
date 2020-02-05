package sematext

type UserRole struct {
	role       string // TODO Enum: [ SUPER_USER, OWNER, ADMIN, USER, DEMO, ANONYMOUS ]
	roleStatus string // TODO Enum: [ INACTIVE, ACTIVE ]
	userEmail  string
}
