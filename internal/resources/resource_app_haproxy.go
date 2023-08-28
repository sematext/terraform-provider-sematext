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


type AppHaproxyResource struct {
	client *stcloud.APIClient
}


// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppHaproxyResource{}
var _ resource.ResourceWithImportState = &AppHaproxyResource{}


func NewAppHaproxyResource() resource.Resource {

	return &AppHaproxyResource{}

}


func (r *AppHaproxyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {

	resp.TypeName = "sematext_app_haproxy"	

}


func (r *AppHaproxyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

	appType := "HAProxy"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		resp.Schema = common.ResourceSchemaAppAWS("HAProxy")
	} else{
		resp.Schema = common.ResourceSchemaAppDefault("HAProxy")
	}

}


func (r *AppHaproxyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	appType := "HAProxy"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpConfigureAppAWS(r.client, ctx, req, resp, "Haproxy")
	} else {
		common.ResourceOpConfigureAppDefault(r.client, ctx, req, resp, "Haproxy")
	}

}


func (r *AppHaproxyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	appType := "HAProxy"
		
	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpCreateAppAWS(r.client, ctx, req, resp, "Haproxy")
	} else{
		common.ResourceOpCreateAppDefault(r.client, ctx, req, resp, "Haproxy")
	}

}


func (r *AppHaproxyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	appType := "HAProxy"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpReadAppAWS(r.client, ctx, req, resp, "Haproxy")
	} else {
		common.ResourceOpReadAppDefault(r.client, ctx, req, resp, "Haproxy")
	}

}


func (r *AppHaproxyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	appType := "HAProxy"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpUpdateAppAWS(r.client, ctx, req, resp, "Haproxy")
	} else {
		common.ResourceOpUpdateAppDefault(r.client, ctx, req, resp, "Haproxy")
	}

}


func (r *AppHaproxyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	appType := "HAProxy"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpDeleteAppAWS(r.client, ctx, req, resp, "Haproxy")
	} else { 
		common.ResourceOpDeleteAppDefault(r.client, ctx, req, resp, "Haproxy")
	}
	

}


func (r *AppHaproxyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	appType := "HAProxy"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
  		common.ResourceOpImportAppAWS(r.client, ctx, req, resp, "Haproxy")
	} else {
		common.ResourceOpImportAppDefault(r.client, ctx, req, resp, "Haproxy")
	}

}

