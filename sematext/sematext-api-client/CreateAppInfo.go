package sematext

import {
	"json"
}

// CreateAppInfo TODO Doc Comment
type CreateAppInfo struct {
	AppType       string      `json:"appType"`
	DiscountCode  string      `json:"discountCode,omitempty"`
	InitialPlanID int64       `json:"initialPlanId"`
	MetaData      AppMetadata `json:"metaData,omitempty"`
	Name          string      `json:"name"`
}

// Create TODO Doc comment
func (createAppInfo *CreateAppInfo) Create(client *APIClient) interface{}, error {

	path := "/spm-reports/api/v3/apps"

	genericAPIResponse, err := client.PostJSON(path, createAppInfo)
	if err!= nil {
		return nil, err
	}

	var extract interface{}
	extract, err := json.Marshall(genericAPIResponse.Data.apps)
	var apps :=[]App
	err := json.UnMarshall(extract, &apps)
	if err!=nil{
		return nill, err
	}

	return apps[0], nil
}
