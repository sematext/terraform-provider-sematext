package sematext

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// App TODO Doc Comment
type App struct {
	AjaxThreshold         int64                `json:"ajaxThreshold"`
	AppType               string               `json:"appType"`
	AppTypeID             int64                `json:"appTypeId"`
	CreatorEmail          string               `json:"creatorEmail"`
	CreditCardExpiry      string               `json:"creditCardExpiry"`
	CreditCardNumber      string               `json:"creditCardNumber"`
	Description           string               `json:"description"`
	DisplayStatus         string               `json:"displayStatus"`
	FirstDataSavedDate    int64                `json:"firstDataSavedDate"`
	ID                    int64                `json:"id"`
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

// Create TODO Doc Comment
func (app App) Create(d *schema.ResourceData, meta interface{}) error {
	application, err := buildSematextApplication(d)
	if err != nil {
		return fmt.Errorf("Failed to parse resource configuration: %s", err.Error())
	}
	application, err = meta.(*sematext.Client).CreateApplication(application)
	if err != nil {
		return fmt.Errorf("Failed to create application using API: %s", err.Error())
	}
	return resourceSematextApplicationRead(d, meta)
}

// Read TODO Doc Comment
func (app App) Read(d *schema.ResourceData, meta interface{}) error {
	//client := meta.(*sematext.Config.Client)
	id := d.get("Id")
	application, err := client.getApplication(id)
	if err != nil {
		panic(err)
	}
	d.set("id", application.id)
	d.set("token", application.token)
	d.set("appType", application.appType)
	d.set("description", application.description)
	d.set("integration", application.integration)               // TODO - ServiceIntegration sub-schema in seperate file?
	d.set("name", application.name)                             // TODO - ServiceIntegration sub-schema in seperate file?
	d.set("owningOrganization", application.owningOrganization) // TODO - ServiceIntegration sub-schema in seperate file?
	d.set("prepaidAccount", application.prepaidAccount)
	d.set("status", application.status)
	d.set("userRoles", application.userRoles) // TODO - ServiceIntegration sub-schema in seperate file?

	return nil

}

// Update TODO Doc Comment
func (app App) Update(d *schema.ResourceData, meta interface{}) error {
	application, err := buildSematextApplication(d)
	if err != nil {
		return fmt.Errorf("Failed to parse resource configuration: %s", err.Error())
	}
	application, err = meta.(*sematext.Client).UpdateApplication(application)
	if err != nil {
		return fmt.Errorf("Failed to create application using API: %s", err.Error())
	}
	return resourceSematextApplicationRead(d, meta)
}

// Delete TODO Doc Comment
func (app App) Delete(d *schema.ResourceData, meta interface{}) error { // TODO check this is protected
	id := d.Id()
	if err := meta.(*sematext.Client).DeleteApplication(id); err != nil {
		return err
	}
	return nil
}

// Exists TODO Doc Comment
func (app App) Exists(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
	id := d.Id()
	if _, err := meta.(*sematext.Config).Client.GetApplication(id); err != nil {
		if strings.Contains(err.Error(), "404 Not Found") {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// Import TODO Doc Comment
func (app App) Import(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	if err := resourceSematextAppRead(d, meta); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}
