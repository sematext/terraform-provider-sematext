package sematext

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/sematext/sematext-api-client-go/stcloud"
)

// CommonMonitorCreate is a common creation handler used by most resources.
func CommonMonitorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}, appType string) diag.Diagnostics {

	var diags diag.Diagnostics
	var err error
	var appsResponse stcloud.AppsResponse
	var tokenResponse stcloud.TokenResponse
	var createTokenDto stcloud.CreateTokenDto
	var appTokenNames []string
	var tokenAccumulator map[string]string

	client := meta.(*stcloud.APIClient) // TODO - get client from context instead of meta?

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

	name, namePresent := d.GetOk("name")
	if namePresent {
		createAppInfo.Name = name.(string)
	} else {
		return diag.FromErr(errors.New("error : missing name field"))
	}

	if discountCode, discountCodePresent := d.GetOk("discount_code"); discountCodePresent {
		createAppInfo.DiscountCode = discountCode.(string)
	}

	if initialPlanID, initialPlanIDPresent := d.GetOk("billing_plan_id"); initialPlanIDPresent {
		if _, found := stcloud.LookupPlanID2Apptypes[initialPlanID.(int)]; !found {
			return diag.FromErr(fmt.Errorf("error : %v is invalid billing_plan_id for %v", initialPlanID, appType))
		}
		createAppInfo.InitialPlanId = int64(initialPlanID.(int))
	} else {
		return diag.FromErr(errors.New("error : missing billing_plan_id"))
	}

	if awsRegion, awsRegionPresent := d.GetOk("aws_region"); awsRegionPresent {
		if stRegion, found := stcloud.AWSRegion2STRegion[awsRegion.(string)]; found {
			createAppInfo.MetaData.AwsRegion = stRegion
		}
	}
	if awsAccessKey, awsAccessKeyPresent := d.GetOk("aws_access_key"); awsAccessKeyPresent {
		createAppInfo.MetaData.AwsCloudWatchAccessKey = awsAccessKey.(string)
	}
	if awsSecretKey, awsSecretKeyPresent := d.GetOk("aws_secret_key"); awsSecretKeyPresent {
		createAppInfo.MetaData.AwsCloudWatchSecretKey = awsSecretKey.(string)
	}
	if awsFetchFrequency, awsFetchFrequencyPresent := d.GetOk("aws_fetch_frequency"); awsFetchFrequencyPresent {
		createAppInfo.MetaData.AwsFetchFrequency = awsFetchFrequency.(string)
	}

	if appType == "Logsene" || appType == "mobile-logs" {
		appsResponse, _, err = client.LogsAppApi.CreateLogseneApplication(ctx, *createAppInfo)
	} else {
		appsResponse, _, err = client.MonitoringAppApi.CreateSpmApplication1(ctx, *createAppInfo)
	}

	if err != nil {
		var body map[string]interface{}
		json.Unmarshal([]byte(err.(stcloud.GenericSwaggerError).Body()), &body)
		return diag.FromErr(errors.New(body["message"].(string)))
	}

	var app *stcloud.App
	app, err = extractFirstApp(&appsResponse)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(app.Id, 10))

	appTokenNames = extractAppTokenNames(d.Get("apptoken"))
	tokenAccumulator = map[string]string{}
	for _, tokenName := range appTokenNames {

		createTokenDto = stcloud.CreateTokenDto{}
		createTokenDto.Name = tokenName
		createTokenDto.Readable = true
		createTokenDto.Writeable = true
		tokenResponse, _, err = client.TokensApiControllerApi.CreateAppToken1(ctx, createTokenDto, app.Id) // TODO handle Model_Error better
		if err != nil {
			return diag.FromErr(err)
		}
		tokenAccumulator[tokenName] = tokenResponse.Data.Token.Token
	}

	d.Set("sc_apptoken_entries", tokenAccumulator)

	return diags

}

// CommonMonitorRead is a common read handler used by most resources.
func CommonMonitorRead(ctx context.Context, d *schema.ResourceData, meta interface{}, appType string) diag.Diagnostics {

	var diags diag.Diagnostics
	var appResponse stcloud.AppResponse
	var id int64
	var err error
	var app *stcloud.App
	var appTokenNames []string
	var tokensResponse stcloud.TokensResponse
	var tokenEntries *[]stcloud.TokenDto
	var tokenEntry stcloud.TokenDto
	var tokenAccumulator map[string]string
	var httpResponse *http.Response

	client := meta.(*stcloud.APIClient)

	if id, err = strconv.ParseInt(d.Id(), 10, 64); err != nil {
		return diag.FromErr(err)
	}

	appResponse, httpResponse, err = client.AppsApi.GetUsingGET(ctx, id)
	if err != nil {
		return diag.FromErr(err)
	}

	//existance check
	if httpResponse.StatusCode == 404 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "resource not found during Read",
			Detail:   "resource '" + d.Get("name").(string) + "' is not present on Sematext Cloud during Read",
		})
		d.SetId("")
		return diags
	}

	app, err = extractApp(&appResponse)
	if err != nil {
		return diag.FromErr(err)
	}

	switch appType {

	case "AWS EBS", "AWS EC2", "AWS ELB":
		d.Set("name", app.Name)
		d.Set("billing_plan_id", app.Plan.Id)

	default:
		d.Set("name", app.Name)
		d.Set("billing_plan_id", app.Plan.Id)
	}

	// get the list of apptoken names that are supposed to be here
	appTokenNames = extractAppTokenNames(d.Get("apptoken"))

	// pull tokens for this app from SC.
	if tokensResponse, _, err = client.TokensApiControllerApi.GetAppTokens(ctx, id); err != nil {
		return diag.FromErr(err)
	}
	tokenEntries, err = extractAppTokens(tokensResponse)
	if err != nil {
		return diag.FromErr(err)
	}

	// filter list of tokenEntries, ignore any not present in the Terraform HCL script
	//   if any are not writable output a warning
	//   if any are missing output a warning
	tokenAccumulator = map[string]string{}
	for _, tokenEntry = range *tokenEntries {
		if contains(appTokenNames, tokenEntry.Name) {
			tokenAccumulator[tokenEntry.Name] = tokenEntry.Token
			if !tokenEntry.Writeable {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Warning,
					Summary:  "apptoken not writable",
					Detail:   "an apptoken named '" + tokenEntry.Name + "' exists but is not writable - update will fail unless made writable",
				})
			}
		}
	}

	for _, tokenName := range appTokenNames {
		if _, found := tokenAccumulator[tokenName]; !found {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  "apptoken not present at sematext cloud",
				Detail:   "an apptoken named '" + tokenEntry.Name + "' is not yet present in your cloud account - terraform update will create this",
			})
		}
	}

	d.Set("sc_apptoken_entries", tokenAccumulator)

	return diags
}

// CommonMonitorUpdate is a common update handler used by most resources.
func CommonMonitorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}, appType string) diag.Diagnostics {

	var diags diag.Diagnostics
	var id int64
	var err error
	var appTokenNames []string
	var tokensResponse stcloud.TokensResponse
	var tokenEntries *[]stcloud.TokenDto
	var tokenEntry stcloud.TokenDto
	var tokenAccumulator map[string]string
	var createTokenDto stcloud.CreateTokenDto
	var tokenResponse stcloud.TokenResponse
	var httpResponse *http.Response

	client := meta.(*stcloud.APIClient)
	if id, err = strconv.ParseInt(d.Id(), 10, 64); err != nil {
		return diag.FromErr(err)
	}

	//existance check
	if httpResponse.StatusCode == 404 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "resource not found during Update",
			Detail:   "resource '" + d.Get("name").(string) + "' is not present on Sematext Cloud during Update",
		})
		return diags
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
		_, _, err = client.AppsApi.UpdateUsingPUT2(context.Background(), *updateAppInfo, id)
		if err != nil {
			return diag.FromErr(err)
		}

	}

	billingInfo := &stcloud.BillingInfo{}
	billingInfoChanged := false

	if d.HasChange("billing_plan_id") {
		_, newBillingPlanId := d.GetChange("billing_plan_id")
		billingInfo.PlanId = newBillingPlanId.(int64)
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
		_, _, err = client.AppsApi.UpdateUsingPUT2(context.Background(), *updateAppInfo, id)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if billingInfoChanged {
		_, _, err = client.BillingApi.UpdatePlanUsingPUT(context.Background(), *billingInfo, id)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if cloudWatchSettingsChanged {
		_, _, err = client.AwsSettingsControllerApi.UpdateUsingPUT1(context.Background(), *cloudWatchSettings, id)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	// get the list of apptoken names that are supposed to be here
	//appTokenNames = extractAppTokenNames(d.Get("apptoken"))

	// pull tokens for this app from SC.
	if tokensResponse, _, err = client.TokensApiControllerApi.GetAppTokens(ctx, id); err != nil {
		return diag.FromErr(err)
	}
	tokenEntries, err = extractAppTokens(tokensResponse)
	if err != nil {
		return diag.FromErr(err)
	}

	//overwrite UUIDs from SC
	tokenAccumulator = map[string]string{}
	for _, tokenEntry = range *tokenEntries {
		if contains(appTokenNames, tokenEntry.Name) {
			tokenAccumulator[tokenEntry.Name] = tokenEntry.Token
			if !tokenEntry.Writeable {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "apptoken not writable at sematext cloud",
					Detail:   "an apptoken named '" + tokenEntry.Name + "' is not writable in your cloud account - aborting lest resources blocked from sending records",
				})
				return diags
			}
		}
	}

	//create tokens for any appTokenNames that are missing from SC
	for _, tokenName := range appTokenNames {
		if _, found := tokenAccumulator[tokenName]; !found {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  "creating apptoken",
				Detail:   "an apptoken named '" + tokenEntry.Name + "' will be created in your cloud account",
			})

			createTokenDto = stcloud.CreateTokenDto{}
			createTokenDto.Name = tokenName
			createTokenDto.Readable = true
			createTokenDto.Writeable = true
			tokenResponse, _, err = client.TokensApiControllerApi.CreateAppToken1(ctx, createTokenDto, id) // TODO handle Model_Error better
			if err != nil {
				return diag.FromErr(err)
			}
			tokenAccumulator[tokenName] = tokenResponse.Data.Token.Token

		}
	}

	d.Set("sc_apptoken_entries", tokenAccumulator)

	return diags

}

// CommonMonitorDelete is a common retire handler used by most resources.
func CommonMonitorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}, appType string) diag.Diagnostics {

	var diags diag.Diagnostics
	var id int64
	var err error
	var httpResponse *http.Response

	client := meta.(*stcloud.APIClient)
	if id, err = strconv.ParseInt(d.Id(), 10, 64); err != nil {
		return diag.FromErr(err)
	}

	updateAppInfo := &stcloud.UpdateAppInfo{}
	updateAppInfo.Status = "DISABLED"
	_, _, err = client.AppsApi.UpdateUsingPUT2(context.Background(), *updateAppInfo, id)
	if err != nil {
		return diag.FromErr(err)
	}

	time.Sleep(2 * time.Second)

	_, httpResponse, err = client.AppsApi.DeleteUsingDELETE(context.Background(), id)
	if err != nil {
		return diag.FromErr(err)
	}

	if httpResponse.StatusCode != 200 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "error deleting app",
			Detail:   "error deleting an app named '" + d.Get("name").(string) + "'",
		})

		return diags
	}

	return diags
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
