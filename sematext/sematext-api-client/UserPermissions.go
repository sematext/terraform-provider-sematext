package sematext

type UserPermissions struct {
	CanDelete bool `json:"canDelete"`
	CanEdit   bool `json:"canEdit"`
	CanView   bool `json:"canView"`
}
