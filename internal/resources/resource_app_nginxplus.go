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


type AppNginxplusResource struct {
	client *stcloud.APIClient
}


// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppNginxplusResource{}
var _ resource.ResourceWithImportState = &AppNginxplusResource{}


func NewAppNginxplusResource() resource.Resource {

	return &AppNginxplusResource{}

}


func (r *AppNginxplusResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {

	resp.TypeName = "sematext_app_nginxplus"	

}


func (r *AppNginxplusResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

	appType := "Nginx-Plus"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		resp.Schema = common.ResourceSchemaAppAWS("Nginx-Plus")
	} else{
		resp.Schema = common.ResourceSchemaAppDefault("Nginx-Plus")
	}

}


func (r *AppNginxplusResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	appType := "Nginx-Plus"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpConfigureAppAWS(r.client, ctx, req, resp, "Nginxplus")
	} else {
		common.ResourceOpConfigureAppDefault(r.client, ctx, req, resp, "Nginxplus")
	}

}


func (r *AppNginxplusResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	appType := "Nginx-Plus"
		
	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpCreateAppAWS(r.client, ctx, req, resp, "Nginxplus")
	} else{
		common.ResourceOpCreateAppDefault(r.client, ctx, req, resp, "Nginxplus")
	}

}


func (r *AppNginxplusResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	appType := "Nginx-Plus"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpReadAppAWS(r.client, ctx, req, resp, "Nginxplus")
	} else {
		common.ResourceOpReadAppDefault(r.client, ctx, req, resp, "Nginxplus")
	}

}


func (r *AppNginxplusResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	appType := "Nginx-Plus"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpUpdateAppAWS(r.client, ctx, req, resp, "Nginxplus")
	} else {
		common.ResourceOpUpdateAppDefault(r.client, ctx, req, resp, "Nginxplus")
	}

}


func (r *AppNginxplusResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	appType := "Nginx-Plus"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpDeleteAppAWS(r.client, ctx, req, resp, "Nginxplus")
	} else { 
		common.ResourceOpDeleteAppDefault(r.client, ctx, req, resp, "Nginxplus")
	}
	

}


func (r *AppNginxplusResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	appType := "Nginx-Plus"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
  		common.ResourceOpImportAppAWS(r.client, ctx, req, resp, "Nginxplus")
	} else {
		common.ResourceOpImportAppDefault(r.client, ctx, req, resp, "Nginxplus")
	}

}

