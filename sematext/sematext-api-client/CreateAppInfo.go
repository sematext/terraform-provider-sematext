package sematext


// CreateAppInfo TODO Doc Comment
type CreateAppInfo struct {
	AppType       string      `json:"appType"`
	DiscountCode  string      `json:"discountCode,omitempty"`
	InitialPlanID int64       `json:"initialPlanId"`
	MetaData      AppMetadata `json:"metaData,omitempty"`
	Name          string      `json:"name"`
}

// Create TODO Doc comment
func (createAppInfo *CreateAppInfo) Create(path string, client *APIClient) interface{}, error {

	genericAPIResponse, err := client.PostJSON(path, createAppInfo)
	if err!= nil {
		return nil, err
	}



	// TODO - convert genericAPIResponse.Data.apps => []Apps



}
