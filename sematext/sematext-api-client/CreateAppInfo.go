package sematext

type CreateAppInfo struct {
	appType       string //TODO Enum? example: aws
	discountCode  string
	initialPlanId int64
	metaData      AppMetadata
	name          string //TODO Check for enum?	example: new-aws-app-1
}
