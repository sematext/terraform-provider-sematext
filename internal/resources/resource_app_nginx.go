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


type AppNginxResource struct {
	client *stcloud.APIClient
}


// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppNginxResource{}
var _ resource.ResourceWithImportState = &AppNginxResource{}


func NewAppNginxResource() resource.Resource {

	return &AppNginxResource{}

}


func (r *AppNginxResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {

	resp.TypeName = "sematext_app_nginx"	

}


func (r *AppNginxResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

	appType := "Nginx"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		resp.Schema = common.ResourceSchemaAppAWS("Nginx")
	} else{
		resp.Schema = common.ResourceSchemaAppDefault("Nginx")
	}

}


func (r *AppNginxResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	appType := "Nginx"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpConfigureAppAWS(r.client, ctx, req, resp, "Nginx")
	} else {
		common.ResourceOpConfigureAppDefault(r.client, ctx, req, resp, "Nginx")
	}

}


func (r *AppNginxResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	appType := "Nginx"
		
	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpCreateAppAWS(r.client, ctx, req, resp, "Nginx")
	} else{
		common.ResourceOpCreateAppDefault(r.client, ctx, req, resp, "Nginx")
	}

}


func (r *AppNginxResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	appType := "Nginx"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpReadAppAWS(r.client, ctx, req, resp, "Nginx")
	} else {
		common.ResourceOpReadAppDefault(r.client, ctx, req, resp, "Nginx")
	}

}


func (r *AppNginxResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	appType := "Nginx"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpUpdateAppAWS(r.client, ctx, req, resp, "Nginx")
	} else {
		common.ResourceOpUpdateAppDefault(r.client, ctx, req, resp, "Nginx")
	}

}


func (r *AppNginxResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	appType := "Nginx"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpDeleteAppAWS(r.client, ctx, req, resp, "Nginx")
	} else { 
		common.ResourceOpDeleteAppDefault(r.client, ctx, req, resp, "Nginx")
	}
	

}


func (r *AppNginxResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	appType := "Nginx"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
  		common.ResourceOpImportAppAWS(r.client, ctx, req, resp, "Nginx")
	} else {
		common.ResourceOpImportAppDefault(r.client, ctx, req, resp, "Nginx")
	}

}

