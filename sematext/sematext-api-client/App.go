package sematext

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// App TODO Doc Comment
type App struct {
	AjaxThreshold         int64                `json:"ajaxThreshold",omitempty`
	AppType               string               `json:"appType"`
	AppTypeID             int64                `json:"appTypeId"`
	CreatorEmail          string               `json:"creatorEmail"`
	CreditCardExpiry      string               `json:"creditCardExpiry"`
	CreditCardNumber      string               `json:"creditCardNumber"`
	Description           string               `json:"description"`
	DisplayStatus         string               `json:"displayStatus"`
	FirstDataSavedDate    int64                `json:"firstDataSavedDate"`
	ID                    string               `json:"id"`
	Integration           ServiceIntegration   `json:"integration"`
	LastDataReceivedDate  int64                `json:"lastDataReceivedDate"`
	LastDataSavedDate     int64                `json:"lastDataSavedDate"`
	LoggedInUserAppRole   string               `json:"loggedInUserAppRole"`
	MonthlyInvoiceAccount bool                 `json:"monthlyInvoiceAccount"`
	Name                  string               `json:"name"`
	OwnerEmail            string               `json:"ownerEmail"`
	OwningOrganization    BasicOrganizationDto `json:"owningOrganization"`
	PageLoadThreshold     int64                `json:"pageLoadThreshold"`
	PaymentMethodID       int64                `json:"paymentMethodId"`
	Plan                  Plan                 `json:"plan"`
	PrepaidAccount        bool                 `json:"prepaidAccount"`
	Status                string               `json:"status"`
	Token                 string               `json:"token"`
	TrialEndDate          int64                `json:"trialEndDate"`
	URLGroupLimit         int32                `json:"urlGroupLimit"`
	UserRoles             []UserRole           `json:"userRoles"`
}

// Pull TODO Doc comment
func (app *App) Retrieve(id string, client *APIClient) (*App, error) { // TODO Rework once API has a route for GET by id

	path := "/spm-reports/api/v3/apps"
	genericAPIResponse, err := client.GetJSON(path, nil)
	if err != nil {
		return nil, err
	}

	app, err = genericAPIResponse.extractAppById(id)
	if err != nil {
		return nil, err
	}

	return app, nil

}

// Exists TODO Doc comment
func (app *App) Exists(id string, client *APIClient) (*bool, error) { // TODO Rework once API has a route for GET by id

	path := "/spm-reports/api/v3/apps"
	genericAPIResponse, err := client.GetJSON(path, nil)
	if err != nil {
		return nil, err
	}

	var result bool
	app, err = genericAPIResponse.extractAppById(id)
	if app != nil {
		result = true
	} else {
		result = false
	}

	return &result, nil

}

// Update TODO Doc comment
func (app *App) Update(id string, client *APIClient, dto Dto) (*App, error) {

	path := fmt.Sprintf("/spm-reports/api/v3/apps/%s", id)
	genericAPIResponse, err := client.PutJSON(path, dto)
	if err != nil {
		return nil, err
	}

	app, err = genericAPIResponse.extractAppById(id)
	if err != nil {
		return nil, err
	}

	return app, nil
}

// Delete TODO Doc comment
func (app *App) Delete(id string, client *APIClient) error {

	path := fmt.Sprintf("/spm-reports/api/v3/apps/%s", id)
	err := client.Delete(path, nil)
	if err != nil {
		return err
	}
	return nil
}
