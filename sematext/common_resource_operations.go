package sematext

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/sematext/sematext-api-client/api"
)

// CommonMonitorCreate TODO Doc Comment
func CommonMonitorCreate(d *schema.ResourceData, meta interface{}, appType string) error {

	fmt.Println("---------------------------------------")
	fmt.Println("CommonMonitorCreate Called")
	fmt.Println("---------------------------------------")

	var err error
	client := meta.(*api.Client)

	// TODO - Check pre-existence of same name/apptype in API (consider as a status update) - post MVP.

	createAppInfo := &api.CreateAppInfo{}

	switch appType {

	case "AWS EBS":
		createAppInfo.MetaData = &api.AppMetadata{}
		createAppInfo.MetaData.SubTypes = []string{"AWS EBS"}
		createAppInfo.AppType = "aws"
	case "AWS ELB":
		createAppInfo.MetaData = &api.AppMetadata{}
		createAppInfo.MetaData.SubTypes = []string{"AWS ELB"}
		createAppInfo.AppType = "aws"
	case "AWS EC2":
		createAppInfo.MetaData = &api.AppMetadata{}
		createAppInfo.MetaData.SubTypes = []string{"AWS EC2"}
		createAppInfo.AppType = "aws"
	default:
		createAppInfo.AppType = appType
	}

	fmt.Println("------------------")
	fmt.Println("AppType")
	fmt.Println(createAppInfo.AppType)
	fmt.Println("------------------")

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
		if _, found := api.LookupPlanID2Apptypes[initialPlanID.(int)]; !found {
			return fmt.Errorf("%v is invalid billing_plan_id for %v", initialPlanID, appType)
		}
		createAppInfo.InitialPlanID = initialPlanID.(int)
	} else {
		return errors.New("Missing billing_plan_id")
	}

	if awsRegion, awsRegionPresent := d.GetOkExists("aws_region"); awsRegionPresent {
		if stRegion, found := api.AWSRegion2STRegion[awsRegion.(string)]; found {
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

	app, err := createAppInfo.Persist(client)
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(app.ID))
	d.Set("token", app.Token) //TODO confirm this becomes available to other resources and decide if it should be a resource paramater in .tf script

	return nil

}

// CommonMonitorRead TODO Doc Comment
func CommonMonitorRead(d *schema.ResourceData, meta interface{}, apptype string) error {

	fmt.Println("---------------------------------------")
	fmt.Println("CommonMonitorRead Called")
	fmt.Println("---------------------------------------")

	client := meta.(*api.Client)
	var app *api.App
	var id int
	var err error
	if id, err = strconv.Atoi(d.Id()); err != nil {
		return err
	}

	if app, err = app.Load(id, client); err != nil {
		return err
	}

	d.Set("name", app.Name)
	d.Set("billing_plan_id", app.Plan.ID)

	return nil
}

// CommonMonitorUpdate TODO Doc Comment
func CommonMonitorUpdate(d *schema.ResourceData, meta interface{}, apptype string) error {

	fmt.Println("---------------------------------------")
	fmt.Println("CommonMonitorUpdate Called")
	fmt.Println("---------------------------------------")

	// TODO - accomodate new endpoint for metadata update

	var id int
	var err error
	var valid bool

	client := meta.(*api.Client)
	if id, err = strconv.Atoi(d.Id()); err != nil {
		return err
	}

	updateAppInfo := &api.UpdateAppInfo{}
	appInfoChanged := false

	if d.HasChange("name") {
		_, newName := d.GetChange("name")
		updateAppInfo.Name = newName.(string) // TODO name vs resource name/label in terraform
		appInfoChanged = true
	}

	if appInfoChanged {
		updateAppInfo.Status = "ACTIVE" // TODO Consider if is this correct or not?
		updateAppInfo.Persist(id, client)
	}

	billingInfo := &api.BillingInfo{}
	billingInfoChanged := false

	if d.HasChange("billing_plan_id") {
		_, newBillingPlanID := d.GetChange("billing_plan_id")
		billingInfo.PlanID = newBillingPlanID.(int)
		billingInfoChanged = true
	}

	cloudWatchSettings := &api.CloudWatchSettings{}
	cloudWatchSettingsChanged := false

	if d.HasChange("aws_access_key") {
		_, newAwsAccessKey := d.GetChange("aws_access_key")
		cloudWatchSettings.AwsCloudWatchAccessKey = newAwsAccessKey.(string)
		cloudWatchSettingsChanged = true
	}
	if d.HasChange("aws_secret_key") {
		_, newAwsSecretKey := d.GetChange("aws_secret_key")
		cloudWatchSettings.AwsCloudWatchSecretKey = newAwsSecretKey.(string)
		cloudWatchSettingsChanged = true
	}
	if d.HasChange("aws_fetch_frequency") {
		_, newAwsFetchFrequency := d.GetChange("aws_fetch_frequency")
		cloudWatchSettings.AwsFetchFrequency = newAwsFetchFrequency.(string)
		cloudWatchSettingsChanged = true
	}
	if d.HasChange("aws_region") {
		_, newAwsRegion := d.GetChange("aws_region")
		cloudWatchSettings.AwsRegion = newAwsRegion.(string)
		cloudWatchSettingsChanged = true
	}

	// update only if both are valid

	if appInfoChanged {
		valid, err = updateAppInfo.IsValid()
		if !valid {
			return err
		}
	}
	if billingInfoChanged {
		valid, err = billingInfo.IsValid()
		if !valid {
			return err
		}
	}
	if appInfoChanged {
		_, err = updateAppInfo.Persist(id, client)
		if err != nil {
			return err
		}
	}
	if billingInfoChanged {
		err = billingInfo.Persist(id, client)
		if err != nil {
			return err
		}
	}

	if cloudWatchSettingsChanged {
		_, err = cloudWatchSettings.Persist(id, client)
		if err != nil {
			return err
		}
	}

	return nil

}

// CommonMonitorDelete TODO Doc Comment
func CommonMonitorDelete(d *schema.ResourceData, meta interface{}, apptype string) error {

	fmt.Println("---------------------------------------")
	fmt.Println("CommonMonitorDelete Called")
	fmt.Println("---------------------------------------")

	var id int
	var err error
	client := meta.(*api.Client)
	if id, err = strconv.Atoi(d.Id()); err != nil {
		return err
	}
	updateAppInfo := &api.UpdateAppInfo{}
	updateAppInfo.Status = "DISABLED"
	if _, err = updateAppInfo.Persist(id, client); err != nil {
		return err
	}

	return nil
}

// CommonMonitorExists TODO Doc Comment
func CommonMonitorExists(d *schema.ResourceData, meta interface{}, apptype string) (exists bool, err error) {

	fmt.Println("---------------------------------------")
	fmt.Println("CommonMonitorExists Called")
	fmt.Println("---------------------------------------")

	// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
	var id int
	var app api.App

	client := meta.(*api.Client)
	if id, err = strconv.Atoi(d.Id()); err != nil {
		return false, err
	}
	exists, err = app.Retired(id, client)
	return exists, err
}

// CommonMonitorImport TODO Doc Comment
func CommonMonitorImport(d *schema.ResourceData, meta interface{}, apptype string) ([]*schema.ResourceData, error) {

	// TODO Decide if Resource Import necessary post-MVP

	fmt.Println("---------------------------------------")
	fmt.Println("CommonMonitorImport Called")
	fmt.Println("---------------------------------------")

	return nil, nil

}
