package sematext

type ServiceIntegration struct {
	AppTypeId           int64  `json:"appTypeId"`
	AppTypeName         string `json:"appTypeName"`
	DisplayName         string `json:"displayName"`
	Enabled             bool   `json:"enabled"`
	ExternalProductId   int64  `json:"externalProductId"`
	ExternalProductName string `json:"externalProductName"`
	Id                  int64  `json:"id"`
	IntegrationType     string `json:"integrationType"`
	Ordinal             int32  `json:"ordinal"`
	ParentIntegrationId int64  `json:"parentIntegrationId"`
	SematextService     string `json:"sematextService"`
	Visible             bool   `json:"visible"`
}
