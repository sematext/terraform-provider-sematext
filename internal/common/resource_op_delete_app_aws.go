package common

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/sematext/sematext-api-client-go/stcloud"
)

func ResourceOpDeleteAWS(client *stcloud.APIClient, ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse, appType string) {

	var id int64
	var err error
	var appResourceModelAWS AppResourceModelAWS
	var httpResponse *http.Response
	var body map[string]interface{}

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &appResourceModelAWS)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if id, err = strconv.ParseInt(appResourceModelAWS.Id, 10, 64); err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete Resource",
			"An unexpected error occurred while attempting to convert the id. "+
				"Please retry the operation or report this issue to the provider developers.\n\n"+
				"Error: "+err.Error(),
		)
		return
	}

	updateAppInfo := &stcloud.UpdateAppInfo{}
	updateAppInfo.Status = "DISABLED"
	_, httpResponse, err = client.AppsApi.UpdateUsingPUT2(context.Background(), *updateAppInfo, id)
	if err != nil {
		json.Unmarshal([]byte(err.(stcloud.GenericSwaggerError).Body()), &body)
		resp.Diagnostics.AddError(
			"Unable to Delete Resource",
			"An unexpected error occurred while attempting to disable a resource prior to deletion. "+
				"Please retry the operation or report this issue to the provider developers.\n\n"+
				"Error: "+body["message"].(string),
		)
		return
	}

	// Return error if the HTTP status code is not 200 OK
	if httpResponse.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError(
			"Unable to Delete Resource",
			"An unexpected error occurred while attempting to disable a resource prior to deletion. "+
				"Please retry the operation or report this issue to the provider developers.\n\n"+
				"HTTP Status: "+httpResponse.Status,
		)
		return
	}

	time.Sleep(2 * time.Second)

	_, httpResponse, err = client.AppsApi.DeleteUsingDELETE(context.Background(), id)
	if err != nil {
		json.Unmarshal([]byte(err.(stcloud.GenericSwaggerError).Body()), &body)
		resp.Diagnostics.AddError(
			"Unable to Delete Resource",
			"An unexpected error occurred while attempting to delete a resource. "+
				"Please retry the operation or report this issue to the provider developers.\n\n"+
				"Error: "+body["message"].(string),
		)
		return
	}

	// Return error if the HTTP status code is not 200 OK
	if httpResponse.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError(
			"Unable to Delete Resource",
			"An unexpected error occurred while attempting to delete a resource. "+
				"Please retry the operation or report this issue to the provider developers.\n\n"+
				"HTTP Status: "+httpResponse.Status,
		)
		return
	}

	// @TODO check how deletion gets back into state.

}
