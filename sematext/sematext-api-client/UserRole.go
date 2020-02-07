package sematext

type UserRole struct {
	Role       string `json:"role"`       // TODO Enum: [ SUPER_USER, OWNER, ADMIN, USER, DEMO, ANONYMOUS ]
	RoleStatus string `json:"roleStatus"` // TODO Enum: [ INACTIVE, ACTIVE ]
	UserEmail  string `json:"userEmail"`
}
