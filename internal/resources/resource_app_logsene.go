package resources

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_App.go.template
	Then run generate/generate.sh
*/

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/sematext/sematext-api-client-go/stcloud"

	"github.com/sematext/terraform-provider-sematext/internal/common"
)


type AppLogseneResource struct {
	client *stcloud.APIClient
}


// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppLogseneResource{}
var _ resource.ResourceWithImportState = &AppLogseneResource{}


func NewAppLogseneResource() resource.Resource {

	return &AppLogseneResource{}

}


func (r *AppLogseneResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {

	resp.TypeName = "sematext_app_logsene"	

}


func (r *AppLogseneResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

	appType := "Logsene"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		resp.Schema = common.ResourceSchemaAppAWS("Logsene")
	} else{
		resp.Schema = common.ResourceSchemaAppDefault("Logsene")
	}

}


func (r *AppLogseneResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	appType := "Logsene"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpConfigureAppAWS(r.client, ctx, req, resp, "Logsene")
	} else {
		common.ResourceOpConfigureAppDefault(r.client, ctx, req, resp, "Logsene")
	}

}


func (r *AppLogseneResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	appType := "Logsene"
		
	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpCreateAppAWS(r.client, ctx, req, resp, "Logsene")
	} else{
		common.ResourceOpCreateAppDefault(r.client, ctx, req, resp, "Logsene")
	}

}


func (r *AppLogseneResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	appType := "Logsene"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpReadAppAWS(r.client, ctx, req, resp, "Logsene")
	} else {
		common.ResourceOpReadAppDefault(r.client, ctx, req, resp, "Logsene")
	}

}


func (r *AppLogseneResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	appType := "Logsene"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpUpdateAppAWS(r.client, ctx, req, resp, "Logsene")
	} else {
		common.ResourceOpUpdateAppDefault(r.client, ctx, req, resp, "Logsene")
	}

}


func (r *AppLogseneResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	appType := "Logsene"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpDeleteAppAWS(r.client, ctx, req, resp, "Logsene")
	} else { 
		common.ResourceOpDeleteAppDefault(r.client, ctx, req, resp, "Logsene")
	}
	

}


func (r *AppLogseneResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	appType := "Logsene"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
  		common.ResourceOpImportAppAWS(r.client, ctx, req, resp, "Logsene")
	} else {
		common.ResourceOpImportAppDefault(r.client, ctx, req, resp, "Logsene")
	}

}

