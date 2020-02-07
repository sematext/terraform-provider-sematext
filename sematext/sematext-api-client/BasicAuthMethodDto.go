package sematext

// BasicAuthMethodDto TODO Doc Comment
type BasicAuthMethodDto struct {
	AuthType string `json:"authType"` // TODO Enum:[ SEMATEXT_ACCOUNT, LDAP ]
	UUID     string `json:"uuid"`
}
