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


type AppMobilelogsResource struct {
	client *stcloud.APIClient
}


// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppMobilelogsResource{}
var _ resource.ResourceWithImportState = &AppMobilelogsResource{}


func NewAppMobilelogsResource() resource.Resource {

	return &AppMobilelogsResource{}

}


func (r *AppMobilelogsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {

	resp.TypeName = "sematext_app_mobilelogs"	

}


func (r *AppMobilelogsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

	appType := "mobile-logs"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		resp.Schema = common.ResourceSchemaAppAWS("mobile-logs")
	} else{
		resp.Schema = common.ResourceSchemaAppDefault("mobile-logs")
	}

}


func (r *AppMobilelogsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	appType := "mobile-logs"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpConfigureAppAWS(r.client, ctx, req, resp, "Mobilelogs")
	} else {
		common.ResourceOpConfigureAppDefault(r.client, ctx, req, resp, "Mobilelogs")
	}

}


func (r *AppMobilelogsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	appType := "mobile-logs"
		
	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpCreateAppAWS(r.client, ctx, req, resp, "Mobilelogs")
	} else{
		common.ResourceOpCreateAppDefault(r.client, ctx, req, resp, "Mobilelogs")
	}

}


func (r *AppMobilelogsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	appType := "mobile-logs"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpReadAppAWS(r.client, ctx, req, resp, "Mobilelogs")
	} else {
		common.ResourceOpReadAppDefault(r.client, ctx, req, resp, "Mobilelogs")
	}

}


func (r *AppMobilelogsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	appType := "mobile-logs"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpUpdateAppAWS(r.client, ctx, req, resp, "Mobilelogs")
	} else {
		common.ResourceOpUpdateAppDefault(r.client, ctx, req, resp, "Mobilelogs")
	}

}


func (r *AppMobilelogsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	appType := "mobile-logs"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpDeleteAppAWS(r.client, ctx, req, resp, "Mobilelogs")
	} else { 
		common.ResourceOpDeleteAppDefault(r.client, ctx, req, resp, "Mobilelogs")
	}
	

}


func (r *AppMobilelogsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	appType := "mobile-logs"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
  		common.ResourceOpImportAppAWS(r.client, ctx, req, resp, "Mobilelogs")
	} else {
		common.ResourceOpImportAppDefault(r.client, ctx, req, resp, "Mobilelogs")
	}

}

