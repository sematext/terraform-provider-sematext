package sematext

type ServiceIntegration struct {
	appTypeId           int64
	appTypeName         string
	displayName         string
	enabled             bool
	externalProductId   int64
	externalProductName string
	id                  int64
	integrationType     string
	ordinal             int32
	parentIntegrationId int64
	sematextService     string
	visible             bool
}
