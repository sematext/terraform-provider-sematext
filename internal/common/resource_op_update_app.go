package sematext

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/sematext/sematext-api-client-go/stcloud"
)

// ResourceOperationUpdateApp is a common update handler used by most resources.
func ResourceOpUpdateApp(resourceApp ResourceApp, ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	var id int64
	var err error
	var appTokenNames []string
	var tokensResponse stcloud.TokensResponse
	var tokenEntries *[]stcloud.TokenDto
	var tokenEntry stcloud.TokenDto
	var tokenAccumulator map[string]string
	var createTokenDto stcloud.CreateTokenDto
	var tokenResponse stcloud.TokenResponse
	var resourceAppModel ResourceAppModel
	var httpResponse *http.Response
	var body map[string]interface{}

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &resourceAppModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if id, err = strconv.ParseInt(resourceAppModel.Id, 10, 64); err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update Resource",
			"An unexpected error occurred while attempting to update the resource. "+
				"Please retry the operation or report this issue to the provider developers.\n\n"+
				"Error: "+err.Error(),
		)
		return
	}

	updateAppInfo := &stcloud.UpdateAppInfo{}
	updateAppInfo.Name = resourceAppModel.Name

	_, httpResponse, err = resourceApp.client.AppsApi.UpdateUsingPUT2(context.Background(), *updateAppInfo, id)

	if err != nil {
		json.Unmarshal([]byte(err.(stcloud.GenericSwaggerError).Body()), &body)
		resp.Diagnostics.AddError(
			"Unable to Update Resource",
			"An unexpected error occurred while attempting to create the resource. "+
				"Please retry the operation or report this issue to the provider developers.\n\n"+
				"Error: "+body["message"].(string),
		)
		return
	}

	// Return error if the HTTP status code is not 200 OK
	if httpResponse.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError(
			"Unable to Update Resource",
			"An unexpected error occurred while attempting to create the resource. "+
				"Please retry the operation or report this issue to the provider developers.\n\n"+
				"HTTP Status: "+httpResponse.Status,
		)
		return
	}

	billingInfo := &stcloud.BillingInfo{}
	billingInfo.PlanId = resourceAppModel.BillingPlanId

	_, httpResponse, err = resourceApp.client.BillingApi.UpdatePlanUsingPUT(context.Background(), *billingInfo, id)

	if err != nil {
		json.Unmarshal([]byte(err.(stcloud.GenericSwaggerError).Body()), &body)
		resp.Diagnostics.AddError(
			"Unable to Update Resource Billing Code",
			"An unexpected error occurred while attempting to create the resource. "+
				"Please retry the operation or report this issue to the provider developers.\n\n"+
				"Error: "+body["message"].(string),
		)
		return
	}

	// Return error if the HTTP status code is not 200 OK
	if httpResponse.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError(
			"Unable to Update Resource Billing Code",
			"An unexpected error occurred while attempting to create the resource. "+
				"Please retry the operation or report this issue to the provider developers.\n\n"+
				"HTTP Status: "+httpResponse.Status,
		)
		return
	}

	cloudWatchSettings := &stcloud.CloudWatchSettings{}
	cloudWatchSettings.AccessKey = resourceAppModel.AwsAccessKey.String()
	cloudWatchSettings.SecretKey = resourceAppModel.AwsSecretKey.String()
	cloudWatchSettings.FetchFrequency = resourceAppModel.AwsFetchFrequency.String()
	cloudWatchSettings.Region = resourceAppModel.AwsRegion.String()

	_, httpResponse, err = resourceApp.client.AwsSettingsControllerApi.UpdateUsingPUT1(context.Background(), *cloudWatchSettings, id)

	if err != nil {
		json.Unmarshal([]byte(err.(stcloud.GenericSwaggerError).Body()), &body)
		resp.Diagnostics.AddError(
			"Unable to Update Resource Billing Code",
			"An unexpected error occurred while attempting to create the resource. "+
				"Please retry the operation or report this issue to the provider developers.\n\n"+
				"Error: "+body["message"].(string),
		)
		return
	}

	// Return error if the HTTP status code is not 200 OK
	if httpResponse.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError(
			"Unable to Update Resource Billing Code",
			"An unexpected error occurred while attempting to create the resource. "+
				"Please retry the operation or report this issue to the provider developers.\n\n"+
				"HTTP Status: "+httpResponse.Status,
		)
		return
	}

	// get the list of apptoken names that are supposed to be here
	appTokenNames = extractAppTokenNames(resourceAppModel.AppToken)

	// pull tokens for this app from SC.
	_, _, err = resourceApp.client.TokensApiControllerApi.GetAppTokens(ctx, id)

	if err != nil {
		json.Unmarshal([]byte(err.(stcloud.GenericSwaggerError).Body()), &body)
		resp.Diagnostics.AddError(
			"Unable to Update Resource Billing Code",
			"An unexpected error occurred while attempting to create the resource. "+
				"Please retry the operation or report this issue to the provider developers.\n\n"+
				"Error: "+body["message"].(string),
		)
		return
	}

	// Return error if the HTTP status code is not 200 OK
	if httpResponse.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError(
			"Unable to Update Resource Billing Code",
			"An unexpected error occurred while attempting to create the resource. "+
				"Please retry the operation or report this issue to the provider developers.\n\n"+
				"HTTP Status: "+httpResponse.Status,
		)
		return
	}

	tokenEntries, err = extractAppTokens(tokensResponse)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update - token error.",
			"An unexpected error occurred while attempting to create the resource. "+
				"Please retry the operation or report this issue to the provider developers.\n\n"+
				"Error: "+err.Error(),
		)
		return
	}

	//overwrite UUIDs from SC
	tokenAccumulator = map[string]string{}
	for _, tokenEntry = range *tokenEntries {
		if contains(appTokenNames, tokenEntry.Name) {
			tokenAccumulator[tokenEntry.Name] = tokenEntry.Token
			if !tokenEntry.Writeable {
				resp.Diagnostics.AddError(
					"Warning : App-token is not writable on Sematext Cloud..",
					"An apptoken named '"+tokenEntry.Name+"' is not writable in your cloud account - aborting lest resources get blocked from sending records.",
				)
				return
			}
		}
	}

	//create tokens for any appTokenNames that are missing from SC
	for _, tokenName := range appTokenNames {
		if _, found := tokenAccumulator[tokenName]; !found {

			resp.Diagnostics.AddError(
				"Warning : App-token will be created on Sematext Cloud.",
				"An app-token named '"+tokenEntry.Name+"' will be created in your cloud account",
			)

			createTokenDto = stcloud.CreateTokenDto{}
			createTokenDto.Name = tokenName
			createTokenDto.Readable = true
			createTokenDto.Writeable = true

			tokenResponse, httpResponse, err = resourceApp.client.TokensApiControllerApi.CreateAppToken1(ctx, createTokenDto, id)

			if err != nil {
				json.Unmarshal([]byte(err.(stcloud.GenericSwaggerError).Body()), &body)
				resp.Diagnostics.AddError(
					"Unable to create app-token.",
					"An unexpected error occurred while attempting to create an app-token while updating the resource. "+
						"Please retry the operation or report this issue to the provider developers.\n\n"+
						"Error: "+body["message"].(string),
				)
				return
			}

			// Return error if the HTTP status code is not 200 OK
			if httpResponse.StatusCode != http.StatusOK {
				resp.Diagnostics.AddError(
					"Unable to create app-token.",
					"An unexpected error occurred while attempting to create an app-token while updating the resource. "+
						"Please retry the operation or report this issue to the provider developers.\n\n"+
						"HTTP Status: "+httpResponse.Status,
				)
				return
			}

			tokenAccumulator[tokenName] = tokenResponse.Data.Token.Token

		}
	}

	resourceAppModel.ScAppTokenEntries = tokenAccumulator

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &resourceAppModel)...)

}
