package operation

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/sematext/sematext-api-client-go/stcloud"

	"github.com/sematext/terraform-provider-sematext/internal/common"
	"github.com/sematext/terraform-provider-sematext/internal/util"
)

func ResourceOpReadDefault(client *stcloud.APIClient, ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse, appType string) {

	var appResponse stcloud.AppResponse
	var id int64
	var err error
	var app *stcloud.App
	var appTokenNames []string
	var tokensResponse stcloud.TokensResponse
	var tokenEntries *[]stcloud.TokenDto
	var tokenEntry stcloud.TokenDto
	var tokenAccumulator map[string]string
	var appResourceModelDefault common.AppResourceModelDefault
	var httpResponse *http.Response
	var body map[string]interface{}

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &appResourceModelDefault)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if id, err = strconv.ParseInt(appResourceModelDefault.Id, 10, 64); err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete Resource",
			"An unexpected error occurred while attempting to convert the id. "+
				"Please retry the operation or report this issue to the provider developers.\n\n"+
				"Error: "+err.Error(),
		)
		return
	}

	appResponse, httpResponse, err = client.AppsApi.GetUsingGET(ctx, id)

	if err != nil {

		json.Unmarshal([]byte(err.(stcloud.GenericSwaggerError).Body()), &body)
		resp.Diagnostics.AddError(
			"Error trying to Read Resource",
			"An unexpected error occurred while attempting to read the resource. "+
				"Please retry the operation or report this issue to the provider developers.\n\n"+
				"Error: "+body["message"].(string),
		)

		return
	}

	// Return error if the HTTP status code is not 200 OK
	if httpResponse.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError(
			"Error trying to Read Resource",
			"An unexpected error occurred while attempting to read the resource. "+
				"Please retry the operation or report this issue to the provider developers.\n\n"+
				"HTTP Status: "+httpResponse.Status,
		)
		return
	}

	app, err = util.ExtractApp(&appResponse)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error trying to Read Resource",
			"An unexpected error occurred while attempting to read the resource (bad response). "+
				"Please retry the operation or report this issue to the provider developers.\n\n"+
				"Error: "+err.Error(),
		)
		return
	}

	appResourceModelDefault.Name = app.Name
	appResourceModelDefault.BillingPlanId = app.Plan.Id

	// get the list of apptoken names that are supposed to be here
	appTokenNames = util.ExtractAppTokenNames(appResourceModelDefault.AppToken)

	// pull tokens for this app from SC.
	if tokensResponse, httpResponse, err = client.TokensApiControllerApi.GetAppTokens(ctx, id); err != nil {
		json.Unmarshal([]byte(err.(stcloud.GenericSwaggerError).Body()), &body)
		resp.Diagnostics.AddError(
			"Error trying to Read Resource",
			"An unexpected error occurred while attempting to read the resource. "+
				"Please retry the operation or report this issue to the provider developers.\n\n"+
				"Error: "+body["message"].(string),
		)

		return
	}

	// Return error if the HTTP status code is not 200 OK
	if httpResponse.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError(
			"Error trying to Read Resource",
			"An unexpected error occurred while attempting to read the resource. "+
				"Please retry the operation or report this issue to the provider developers.\n\n"+
				"HTTP Status: "+httpResponse.Status,
		)
		return
	}

	tokenEntries, err = util.ExtractAppTokens(tokensResponse)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error trying to Read Resource",
			"An unexpected error occurred while attempting to read the resource (bad response). "+
				"Please retry the operation or report this issue to the provider developers.\n\n"+
				"Error: "+err.Error(),
		)
		return
	}

	// filter list of tokenEntries, ignore any not present in the Terraform HCL script
	//   if any are not writable output a warning
	//   if any are missing output a warning
	tokenAccumulator = map[string]string{}
	for _, tokenEntry = range *tokenEntries {
		if util.Contains(appTokenNames, tokenEntry.Name) {
			tokenAccumulator[tokenEntry.Name] = tokenEntry.Token
			if !tokenEntry.Writeable {
				resp.Diagnostics.AddError(
					"Warning: AppToken on Sematext Cloud is not set writable",
					"An apptoken named '"+tokenEntry.Name+"' exists but is not writable - update will fail unless made writable.",
				)
			}
		}
	}

	for _, tokenName := range appTokenNames {
		if _, found := tokenAccumulator[tokenName]; !found {

			resp.Diagnostics.AddError(
				"Warning: AppToken on Sematext Cloud does not exist yet.",
				"An apptoken named '"+tokenEntry.Name+"' is not yet present in your cloud account - terraform update will create this.",
			)

		}
	}

	appResourceModelDefault.ScAppTokenEntries = tokenAccumulator

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &appResourceModelDefault)...)

}