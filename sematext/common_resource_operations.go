package sematext

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/sematext/sematext-api-client/stcloud"
)

// CommonMonitorCreate is a common creation handler used by most resources.
func CommonMonitorCreate(d *schema.ResourceData, meta interface{}, appType string) error {

	var err error
	var genericAPIResponse stcloud.GenericAPIResponse

	client := meta.(*stcloud.APIClient)

	// TODO - Check pre-existence of same name/apptype in API (consider as a status update) - post MVP.

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

	switch appType {

	case "Logsene":
		genericAPIResponse, _, err = client.LogsAppAPI.CreateLogseneApplication(context.Background(), *createAppInfo)

	default:
		genericAPIResponse, _, err = client.MonitoringAppAPI.CreateSpmApplication1(context.Background(), *createAppInfo)
	}

	if err != nil {
		return err
	}
	var app *stcloud.App
	app, err = genericAPIResponse.ExtractApp()
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(app.ID, 10))
	d.Set("token", app.Token) //TODO confirm this becomes available to other resources and decide if it should be a resource paramater in .tf script

	return nil

}

// CommonMonitorRead is a common read handler used by most resources.
func CommonMonitorRead(d *schema.ResourceData, meta interface{}, appType string) error {

	client := meta.(*stcloud.APIClient)
	var genericAPIResponse stcloud.GenericAPIResponse
	var id int64
	var err error
	var app *stcloud.App

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

	return nil
}

// CommonMonitorUpdate is a common update handler used by most resources.
func CommonMonitorUpdate(d *schema.ResourceData, meta interface{}, apptype string) error {

	var id int64
	var err error

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

	return nil

}

// CommonMonitorDelete is a common retire handler used by most resources.
func CommonMonitorDelete(d *schema.ResourceData, meta interface{}, apptype string) error {

	var id int64
	var err error

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

	exists = app.Status != "ARCHIVED" // TODO Reconfirm archive vs deletion once some examples available post MVP

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
