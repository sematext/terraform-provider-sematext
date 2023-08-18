package sematext

import (
	"context"
	"net/http"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/sematext/sematext-api-client-go/stcloud"
)

// ResourceOperationUpdateApp is a common update handler used by most resources.
func ResourceOpUpdateApp(ctx context.Context, d *schema.ResourceData, meta interface{}, appType string) diag.Diagnostics {

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
