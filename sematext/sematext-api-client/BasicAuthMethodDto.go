package sematext

type BasicAuthMethodDto struct {
	authType string // TODO Enum:[ SEMATEXT_ACCOUNT, LDAP ]
	uuid     string
}
