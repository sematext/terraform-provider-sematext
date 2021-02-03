package sematext

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mitchellh/mapstructure"
	"github.com/sematext/sematext-api-client-go/stcloud"
)

// CommonMonitorCreate is a common creation handler used by most resources.
func CommonMonitorCreate(d *schema.ResourceData, meta interface{}, appType string) error {

	var err error
	var genericAPIResponse stcloud.GenericAPIResponse
	var appTokenEntries *[]stcloud.AppTokenEntry

	client := meta.(*stcloud.APIClient)

	createAppInfo := &stcloud.CreateAppInfo{}

	switch appType {

	case "AWS EBS":
		createAppInfo.MetaData = &stcloud.AppMetadata{}
		createAppInfo.MetaData.SubTypes = []string{"AWS EBS"}
		createAppInfo.AppType = "aws"
	case "AWS ELB":
		createAppInfo.MetaData = &stcloud.AppMetadata{}
		createAppInfo.MetaData.SubTypes = []string{"AWS ELB"}
		createAppInfo.AppType = "aws"
	case "AWS EC2":
		createAppInfo.MetaData = &stcloud.AppMetadata{}
		createAppInfo.MetaData.SubTypes = []string{"AWS EC2"}
		createAppInfo.AppType = "aws"
	default:
		createAppInfo.AppType = appType
	}

	name, namePresent := d.GetOkExists("name")
	if namePresent {
		createAppInfo.Name = name.(string)
	} else {
		return errors.New("Missing name field")
	}

	if discountCode, discountCodePresent := d.GetOkExists("discount_code"); discountCodePresent {
		createAppInfo.DiscountCode = discountCode.(string)
	}

	if initialPlanID, initialPlanIDPresent := d.GetOkExists("billing_plan_id"); initialPlanIDPresent {
		if _, found := stcloud.LookupPlanID2Apptypes[initialPlanID.(int)]; !found {
			return fmt.Errorf("%v is invalid billing_plan_id for %v", initialPlanID, appType)
		}
		createAppInfo.InitialPlanID = int64(initialPlanID.(int))
	} else {
		return errors.New("Missing billing_plan_id")
	}

	if awsRegion, awsRegionPresent := d.GetOkExists("aws_region"); awsRegionPresent {
		if stRegion, found := stcloud.AWSRegion2STRegion[awsRegion.(string)]; found {
			createAppInfo.MetaData.AwsRegion = stRegion
		}
	}
	if awsAccessKey, awsAccessKeyPresent := d.GetOkExists("aws_access_key"); awsAccessKeyPresent {
		createAppInfo.MetaData.AwsCloudWatchAccessKey = awsAccessKey.(string)
	}
	if awsSecretKey, awsSecretKeyPresent := d.GetOkExists("aws_secret_key"); awsSecretKeyPresent {
		createAppInfo.MetaData.AwsCloudWatchSecretKey = awsSecretKey.(string)
	}
	if awsFetchFrequency, awsFetchFrequencyPresent := d.GetOkExists("aws_fetch_frequency"); awsFetchFrequencyPresent {
		createAppInfo.MetaData.AwsFetchFrequency = awsFetchFrequency.(string)
	}

	if appType == "Logsene" || appType == "mobile-logs" {
		genericAPIResponse, _, err = client.LogsAppAPI.CreateLogseneApplication(context.Background(), *createAppInfo)
	} else {
		genericAPIResponse, _, err = client.MonitoringAppAPI.CreateSpmApplication1(context.Background(), *createAppInfo)
	}

	//spew.Dump(createAppInfo)
	//spew.Dump(genericAPIResponse)

	if err != nil {
		return err
	}

	var app *stcloud.App
	app, err = genericAPIResponse.ExtractApp()
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(app.ID, 10))

	appTokenName, appTokenNamePresent := d.GetOkExists("apptoken.name")
	if appTokenNamePresent {
		createTokenDto := &stcloud.CreateTokenDto{}
		createTokenDto.Name = appTokenName.(string)
		createTokenDto.Readable = true
		createTokenDto.Writeable = true

		// app created , create and record apptoken to state
		genericAPIResponse, _, err = client.TokensAPIControllerAPI.CreateAppToken(context.Background(), *createTokenDto, app.ID)
		if err != nil {
			return err
		}
		appTokenEntries, err = ExtractAppTokens(genericAPIResponse)
		if err != nil {
			return err
		}
		d.Set("apptoken.id", (*appTokenEntries)[0].ID)

	} else {
		return errors.New("Missing apptoken.name field in resource")
	}

	return nil

}

// CommonMonitorRead is a common read handler used by most resources.
func CommonMonitorRead(d *schema.ResourceData, meta interface{}, appType string) error {

	client := meta.(*stcloud.APIClient)
	var genericAPIResponse stcloud.GenericAPIResponse
	var id int64
	var err error
	var app *stcloud.App
	var appTokenEntries *[]stcloud.AppTokenEntry

	if id, err = strconv.ParseInt(d.Id(), 10, 64); err != nil {
		return err
	}

	genericAPIResponse, _, err = client.AppsAPI.GetUsingGET(context.Background(), id)
	if err != nil {
		return err
	}

	app, err = genericAPIResponse.ExtractApp()
	if err != nil {
		return err
	}

	switch appType {

	case "AWS EBS", "AWS EC2", "AWS ELB":
		d.Set("name", app.Name)
		d.Set("billing_plan_id", app.Plan.ID)

	default:
		d.Set("name", app.Name)
		d.Set("billing_plan_id", app.Plan.ID)

	}

	appTokenName, appTokenNamePresent := d.GetOkExists("apptoken.name")
	if !appTokenNamePresent {
		return errors.New("Missing apptoken.name in resource")
	}

	// pull tokens for this app.
	genericAPIResponse, _, err = client.TokensAPIControllerAPI.GetAppTokens1(context.Background(), id)
	if err != nil {
		return err
	}

	appTokenEntries, err = ExtractAppTokens(genericAPIResponse)
	if err != nil {
		return err
	}

	// Note if the id changes then downstream reconfig of dependant resources should occur.
	for _, appTokenEntry := range *appTokenEntries {
		if appTokenEntry.Name == appTokenName {
			if d.Get("apptoken.id") != appTokenEntry.ID {
				fmt.Println("Note  - resource " + d.Get("name").(string) + " has changed had apptoken.id reset for " + appTokenName.(string) + " - dependant resources should reconfigure APM clients")
				d.Set("apptoken.id", appTokenEntry.ID)
			}
			return nil
		}
	}

	return errors.New(d.Get("apptoken.name").(string) + " expected in Sematext Cloud but cannot be found")
}

// CommonMonitorUpdate is a common update handler used by most resources.
func CommonMonitorUpdate(d *schema.ResourceData, meta interface{}, apptype string) error {

	var id int64
	var err error
	var genericAPIResponse stcloud.GenericAPIResponse
	var appTokenEntries *[]stcloud.AppTokenEntry

	client := meta.(*stcloud.APIClient)
	if id, err = strconv.ParseInt(d.Id(), 10, 64); err != nil {
		return err
	}

	updateAppInfo := &stcloud.UpdateAppInfo{}
	appInfoChanged := false

	if d.HasChange("name") {
		_, newName := d.GetChange("name")
		updateAppInfo.Name = newName.(string)
		appInfoChanged = true
	}

	if appInfoChanged {
		updateAppInfo.Status = "ACTIVE"
		_, _, err = client.AppsAPI.UpdateUsingPUT1(context.Background(), *updateAppInfo, id)
		if err != nil {
			return err
		}

	}

	billingInfo := &stcloud.BillingInfo{}
	billingInfoChanged := false

	if d.HasChange("billing_plan_id") {
		_, newBillingPlanID := d.GetChange("billing_plan_id")
		billingInfo.PlanID = newBillingPlanID.(int64)
		billingInfoChanged = true
	}

	cloudWatchSettings := &stcloud.CloudWatchSettings{}
	cloudWatchSettingsChanged := false

	if d.HasChange("aws_access_key") {
		_, newAwsAccessKey := d.GetChange("aws_access_key")
		cloudWatchSettings.AccessKey = newAwsAccessKey.(string)
		cloudWatchSettingsChanged = true
	}
	if d.HasChange("aws_secret_key") {
		_, newAwsSecretKey := d.GetChange("aws_secret_key")
		cloudWatchSettings.SecretKey = newAwsSecretKey.(string)
		cloudWatchSettingsChanged = true
	}
	if d.HasChange("aws_fetch_frequency") {
		_, newAwsFetchFrequency := d.GetChange("aws_fetch_frequency")
		cloudWatchSettings.FetchFrequency = newAwsFetchFrequency.(string)
		cloudWatchSettingsChanged = true
	}
	if d.HasChange("aws_region") {
		_, newAwsRegion := d.GetChange("aws_region")
		cloudWatchSettings.Region = newAwsRegion.(string)
		cloudWatchSettingsChanged = true
	}

	if appInfoChanged {
		_, _, err = client.AppsAPI.UpdateUsingPUT1(context.Background(), *updateAppInfo, id)
		if err != nil {
			return err
		}
	}

	if billingInfoChanged {
		_, _, err = client.BillingAPI.UpdatePlanUsingPUT(context.Background(), id, *billingInfo)
		if err != nil {
			return err
		}
	}

	if cloudWatchSettingsChanged {
		_, _, err = client.AwsSettingsControllerAPI.UpdateUsingPUT(context.Background(), id, *cloudWatchSettings)
		if err != nil {
			return err
		}

	}

	if d.HasChange("apptoken.name") {
		_, newTokenName := d.GetChange("apptoken.name")

		// pull tokens for this app.
		genericAPIResponse, _, err = client.TokensAPIControllerAPI.GetAppTokens1(context.Background(), id)
		if err != nil {
			return err
		}

		appTokenEntries, err = ExtractAppTokens(genericAPIResponse)
		if err != nil {
			return err
		}

		// Note if the id changes then downstream reconfig of dependant resources should occur.
		for _, appTokenEntry := range *appTokenEntries {
			if appTokenEntry.Name == newTokenName {
				if d.Get("apptoken.id") != appTokenEntry.ID {
					fmt.Println("Note  - resource " + d.Get("name").(string) + " has changed had apptoken.id reset for " + newTokenName.(string) + " - dependant resources should reconfigure APM clients")
					d.Set("apptoken.id", appTokenEntry.ID)
				}
				return nil
			}
		}

		// check if apptoken.name exists, create and record apptoken.id to state if not

		if d.Get("apptoken.create_missing").(bool) { // @TODO What if this is missing?
			createTokenDto := &stcloud.CreateTokenDto{}
			createTokenDto.Name = newTokenName.(string)
			createTokenDto.Readable = true
			createTokenDto.Writeable = true
			genericAPIResponse, _, err = client.TokensAPIControllerAPI.CreateAppToken(context.Background(), *createTokenDto, id)
			if err != nil {
				return err
			}
			appTokenEntries, err = ExtractAppTokens(genericAPIResponse)
			if err != nil {
				return err
			}
			d.Set("token.id", (*appTokenEntries)[0].ID)

		} else {
			return errors.New("apptoken.name has changed for resource " + d.Get("name").(string) + " apptoken.create_missing is turned off")
		}

	}

	return nil

}

// CommonMonitorDelete is a common retire handler used by most resources.
func CommonMonitorDelete(d *schema.ResourceData, meta interface{}, apptype string) error {

	var id int64
	var err error
	var response *http.Response

	client := meta.(*stcloud.APIClient)
	if id, err = strconv.ParseInt(d.Id(), 10, 64); err != nil {
		return err
	}

	updateAppInfo := &stcloud.UpdateAppInfo{}
	updateAppInfo.Status = "DISABLED"
	_, _, err = client.AppsAPI.UpdateUsingPUT1(context.Background(), *updateAppInfo, id)
	if err != nil {
		return err
	}

	time.Sleep(2 * time.Second)

	_, response, err = client.AppsAPI.DeleteUsingDELETE1(context.Background(), id)
	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return errors.New("ERROR : problem prevented app from being deleted")
	}

	return nil
}

// CommonMonitorExists is a common existance check handler used by most resources.
func CommonMonitorExists(d *schema.ResourceData, meta interface{}, apptype string) (exists bool, err error) {

	var id int64
	var app *stcloud.App
	var genericAPIResponse stcloud.GenericAPIResponse

	client := meta.(*stcloud.APIClient)

	if id, err = strconv.ParseInt(d.Id(), 10, 64); err != nil {
		return false, err
	}

	if genericAPIResponse, _, err = client.AppsAPI.GetUsingGET(context.Background(), id); err != nil {
		return false, nil
	}

	app, err = genericAPIResponse.ExtractApp()
	if err != nil {
		return false, err
	}

	exists = !(app.Status == "ARCHIVED" || app.Status == "DELETED")

	return exists, nil
}

/*

Placeholder - not implemented

// CommonMonitorImport  is a common import handler used by most resources.
func CommonMonitorImport(d *schema.ResourceData, meta interface{}, apptype string) ([]*schema.ResourceData, error) {

	// TODO Decide if Resource Import necessary post-MVP

	fmt.Println("---------------------------------------")
	fmt.Println("CommonMonitorImport Called")
	fmt.Println("---------------------------------------")

	return nil, nil

}

*/

// ExtractAppTokens pulls token or tokens field out of GenericAPIResponse
func ExtractAppTokens(genericAPIResponse stcloud.GenericAPIResponse) (*[]stcloud.AppTokenEntry, error) {

	var dataField map[string]interface{}
	var appTokenField []interface{}
	var appTokenEntries []stcloud.AppTokenEntry
	var appTokenEntry stcloud.AppTokenEntry
	var exists bool

	dataField = (*genericAPIResponse.Data).(map[string]interface{})
	if appTokenField, exists = dataField["tokens"].([]interface{}); exists {
		mapstructure.Decode(appTokenField, appTokenEntries)
		return &appTokenEntries, nil
	}
	if appTokenField, exists = dataField["token"].([]interface{}); exists {
		mapstructure.Decode(appTokenField, appTokenEntry)
		appTokenEntries[0] = appTokenEntry
		return &appTokenEntries, nil
	}

	return nil, fmt.Errorf("Unexpected missing token(s) field in API response")

}
