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


type AppStormResource struct {
	client *stcloud.APIClient
}


// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppStormResource{}
var _ resource.ResourceWithImportState = &AppStormResource{}


func NewAppStormResource() resource.Resource {

	return &AppStormResource{}

}


func (r *AppStormResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {

	resp.TypeName = "sematext_app_storm"	

}


func (r *AppStormResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

	appType := "Storm"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		resp.Schema = common.ResourceSchemaAppAWS("Storm")
	} else{
		resp.Schema = common.ResourceSchemaAppDefault("Storm")
	}

}


func (r *AppStormResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	appType := "Storm"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpConfigureAppAWS(r.client, ctx, req, resp, "Storm")
	} else {
		common.ResourceOpConfigureAppDefault(r.client, ctx, req, resp, "Storm")
	}

}


func (r *AppStormResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	appType := "Storm"
		
	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpCreateAppAWS(r.client, ctx, req, resp, "Storm")
	} else{
		common.ResourceOpCreateAppDefault(r.client, ctx, req, resp, "Storm")
	}

}


func (r *AppStormResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	appType := "Storm"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpReadAppAWS(r.client, ctx, req, resp, "Storm")
	} else {
		common.ResourceOpReadAppDefault(r.client, ctx, req, resp, "Storm")
	}

}


func (r *AppStormResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	appType := "Storm"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpUpdateAppAWS(r.client, ctx, req, resp, "Storm")
	} else {
		common.ResourceOpUpdateAppDefault(r.client, ctx, req, resp, "Storm")
	}

}


func (r *AppStormResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	appType := "Storm"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpDeleteAppAWS(r.client, ctx, req, resp, "Storm")
	} else { 
		common.ResourceOpDeleteAppDefault(r.client, ctx, req, resp, "Storm")
	}
	

}


func (r *AppStormResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	appType := "Storm"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
  		common.ResourceOpImportAppAWS(r.client, ctx, req, resp, "Storm")
	} else {
		common.ResourceOpImportAppDefault(r.client, ctx, req, resp, "Storm")
	}

}

