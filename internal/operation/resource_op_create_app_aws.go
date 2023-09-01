package operation

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/sematext/sematext-api-client-go/stcloud"

	"github.com/sematext/terraform-provider-sematext/internal/common"
	"github.com/sematext/terraform-provider-sematext/internal/util"
)

func ResourceOpCreateAWS(client *stcloud.APIClient, ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse, appType string) {

	if resp.Diagnostics.HasError() {
		return
	}

	var err error
	var appsResponse stcloud.AppsResponse
	var tokenResponse stcloud.TokenResponse
	var createTokenDto stcloud.CreateTokenDto
	var appTokenNames []string
	var tokenAccumulator map[string]string
	var resourceAppModelAWS common.AppResourceModelAWS

	var httpResponse *http.Response
	var body map[string]interface{}

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &resourceAppModelAWS)...)

	if resp.Diagnostics.HasError() {
		return
	}

	createAppInfo := &stcloud.CreateAppInfo{}
	createAppInfo.Name = resourceAppModelAWS.Name
	createAppInfo.InitialPlanId = resourceAppModelAWS.BillingPlanId //@TODO - validate in stcloud.LookupPlanID2Apptypes[initialPlanID

	switch appType {

	case "AWS EBS":
		createAppInfo.MetaData = &stcloud.AppMetadata{}
		createAppInfo.MetaData.SubTypes = []string{"AWS EBS"}
		createAppInfo.AppType = "aws"
		createAppInfo.MetaData.AwsRegion = resourceAppModelAWS.AwsRegion.String()
		createAppInfo.MetaData.AwsCloudWatchAccessKey = resourceAppModelAWS.AwsAccessKey.String()
		createAppInfo.MetaData.AwsCloudWatchSecretKey = resourceAppModelAWS.AwsSecretKey.String()
		createAppInfo.MetaData.AwsFetchFrequency = resourceAppModelAWS.AwsFetchFrequency.String()

	case "AWS ELB":
		createAppInfo.MetaData = &stcloud.AppMetadata{}
		createAppInfo.MetaData.SubTypes = []string{"AWS ELB"}
		createAppInfo.AppType = "aws"
		createAppInfo.MetaData.AwsRegion = resourceAppModelAWS.AwsRegion.String()
		createAppInfo.MetaData.AwsCloudWatchAccessKey = resourceAppModelAWS.AwsAccessKey.String()
		createAppInfo.MetaData.AwsCloudWatchSecretKey = resourceAppModelAWS.AwsSecretKey.String()
		createAppInfo.MetaData.AwsFetchFrequency = resourceAppModelAWS.AwsFetchFrequency.String()

	case "AWS EC2":
		createAppInfo.MetaData = &stcloud.AppMetadata{}
		createAppInfo.MetaData.SubTypes = []string{"AWS EC2"}
		createAppInfo.AppType = "aws"
		createAppInfo.MetaData.AwsRegion = resourceAppModelAWS.AwsRegion.String()
		createAppInfo.MetaData.AwsCloudWatchAccessKey = resourceAppModelAWS.AwsAccessKey.String()
		createAppInfo.MetaData.AwsCloudWatchSecretKey = resourceAppModelAWS.AwsSecretKey.String()
		createAppInfo.MetaData.AwsFetchFrequency = resourceAppModelAWS.AwsFetchFrequency.String()

	}

	if appType == "Logsene" || appType == "mobile-logs" {
		appsResponse, httpResponse, err = client.LogsAppApi.CreateLogseneApplication(ctx, *createAppInfo)
	} else {
		appsResponse, httpResponse, err = client.MonitoringAppApi.CreateSpmApplication1(ctx, *createAppInfo)
	}

	if err != nil {

		json.Unmarshal([]byte(err.(stcloud.GenericSwaggerError).Body()), &body)
		resp.Diagnostics.AddError(
			"Unable to Create Resource",
			"An unexpected error occurred while attempting to create the resource. "+
				"Please retry the operation or report this issue to the provider developers.\n\n"+
				"Error: "+body["message"].(string),
		)

		return
	}

	// Return error if the HTTP status code is not 200 OK
	if httpResponse.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError(
			"Unable to Create Resource",
			"An unexpected error occurred while attempting to create the resource. "+
				"Please retry the operation or report this issue to the provider developers.\n\n"+
				"HTTP Status: "+httpResponse.Status,
		)
		return
	}

	var app *stcloud.App
	app, err = util.ExtractFirstApp(&appsResponse)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Resource",
			"An unexpected error occurred while attempting to parse the created resource. "+
				"Please retry the operation or report this issue to the provider developers.",
		)
		return
	}

	resourceAppModelAWS.Id = strconv.FormatInt(app.Id, 10)

	appTokenNames = app.Tokens

	tokenAccumulator = map[string]string{}
	for _, tokenName := range appTokenNames {

		createTokenDto = stcloud.CreateTokenDto{}
		createTokenDto.Name = tokenName
		createTokenDto.Readable = true
		createTokenDto.Writeable = true
		tokenResponse, httpResponse, err = client.TokensApiControllerApi.CreateAppToken1(ctx, createTokenDto, app.Id) // TODO handle Model_Error better

		// Return error if the HTTP status code is not 200 OK
		if httpResponse.StatusCode != http.StatusOK {
			resp.Diagnostics.AddError(
				"Unable to retrieve app-token.",
				"An unexpected error occurred while attempting to retrieve the app-token."+
					"Please retry the operation or report this issue to the provider developers.\n\n"+
					"HTTP Status: "+httpResponse.Status,
			)
			return
		}

		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Create Resource",
				"An unexpected error occurred while attempting to create the resource. "+
					"Please retry the operation or report this issue to the provider developers.\n\n"+
					"Error: "+body["message"].(string),
			)
			return
		}

		tokenAccumulator[tokenName] = tokenResponse.Data.Token.Token
	}

	resourceAppModelAWS.ScAppTokenEntries = tokenAccumulator

	tflog.Trace(ctx, "created a resource")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &resourceAppModelAWS)...)

}