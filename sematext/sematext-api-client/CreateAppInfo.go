package sematext

// CreateAppInfo TODO Doc Comment
type CreateAppInfo struct {
	AppType       string      `json:"appType"` //TODO Enum? example: aws
	DiscountCode  string      `json:"discountCode"`
	InitialPlanID int64       `json:"initialPlanId"`
	MetaData      AppMetadata `json:"metaData"`
	Name          string      `json:"name"` //TODO Check for enum?	example: new-aws-app-1
}
