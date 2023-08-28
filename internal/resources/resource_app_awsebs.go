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


type AppAwsebsResource struct {
	client *stcloud.APIClient
}


// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppAwsebsResource{}
var _ resource.ResourceWithImportState = &AppAwsebsResource{}


func NewAppAwsebsResource() resource.Resource {

	return &AppAwsebsResource{}

}


func (r *AppAwsebsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {

	resp.TypeName = "sematext_app_awsebs"	

}


func (r *AppAwsebsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

	appType := "AWS EBS"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		resp.Schema = common.ResourceSchemaAppAWS("AWS EBS")
	} else{
		resp.Schema = common.ResourceSchemaAppDefault("AWS EBS")
	}

}


func (r *AppAwsebsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	appType := "AWS EBS"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpConfigureAppAWS(r.client, ctx, req, resp, "Awsebs")
	} else {
		common.ResourceOpConfigureAppDefault(r.client, ctx, req, resp, "Awsebs")
	}

}


func (r *AppAwsebsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	appType := "AWS EBS"
		
	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpCreateAppAWS(r.client, ctx, req, resp, "Awsebs")
	} else{
		common.ResourceOpCreateAppDefault(r.client, ctx, req, resp, "Awsebs")
	}

}


func (r *AppAwsebsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	appType := "AWS EBS"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpReadAppAWS(r.client, ctx, req, resp, "Awsebs")
	} else {
		common.ResourceOpReadAppDefault(r.client, ctx, req, resp, "Awsebs")
	}

}


func (r *AppAwsebsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	appType := "AWS EBS"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpUpdateAppAWS(r.client, ctx, req, resp, "Awsebs")
	} else {
		common.ResourceOpUpdateAppDefault(r.client, ctx, req, resp, "Awsebs")
	}

}


func (r *AppAwsebsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	appType := "AWS EBS"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpDeleteAppAWS(r.client, ctx, req, resp, "Awsebs")
	} else { 
		common.ResourceOpDeleteAppDefault(r.client, ctx, req, resp, "Awsebs")
	}
	

}


func (r *AppAwsebsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	appType := "AWS EBS"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
  		common.ResourceOpImportAppAWS(r.client, ctx, req, resp, "Awsebs")
	} else {
		common.ResourceOpImportAppDefault(r.client, ctx, req, resp, "Awsebs")
	}

}

