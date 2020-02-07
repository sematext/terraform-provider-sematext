package sematext

type BasicAuthMethodDto struct {
	AuthType string `json:"authType"` // TODO Enum:[ SEMATEXT_ACCOUNT, LDAP ]
	Uuid     string `json:"uuid"`
}
