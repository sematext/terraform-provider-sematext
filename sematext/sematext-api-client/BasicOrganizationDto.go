package sematext

type BasicOrganizationDto struct {
	AuthMethods []BasicAuthMethodDto `json:"authMethods"`
	Name        string               `json:"name"`
	Status      string               `json:"status"` // TODO Enum:[ ACTIVE, IN_REGISTRATION, DISABLED, EXPIRED, INVITED, DEMO ]
	Uuid        string               `json:"uuid"`
}
