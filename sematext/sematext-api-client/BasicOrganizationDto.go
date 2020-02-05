package sematext

type BasicOrganizationDto struct {
	authMethods []BasicAuthMethodDto
	name        string
	status      string // TODO Enum:[ ACTIVE, IN_REGISTRATION, DISABLED, EXPIRED, INVITED, DEMO ]
	uuid        string
}
