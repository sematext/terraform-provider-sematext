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


type AppNodejsResource struct {
	client *stcloud.APIClient
}


// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppNodejsResource{}
var _ resource.ResourceWithImportState = &AppNodejsResource{}


func NewAppNodejsResource() resource.Resource {

	return &AppNodejsResource{}

}


func (r *AppNodejsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {

	resp.TypeName = "sematext_app_nodejs"	

}


func (r *AppNodejsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

	appType := "Node.js"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		resp.Schema = common.ResourceSchemaAppAWS("Node.js")
	} else{
		resp.Schema = common.ResourceSchemaAppDefault("Node.js")
	}

}


func (r *AppNodejsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	appType := "Node.js"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpConfigureAppAWS(r.client, ctx, req, resp, "Nodejs")
	} else {
		common.ResourceOpConfigureAppDefault(r.client, ctx, req, resp, "Nodejs")
	}

}


func (r *AppNodejsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	appType := "Node.js"
		
	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpCreateAppAWS(r.client, ctx, req, resp, "Nodejs")
	} else{
		common.ResourceOpCreateAppDefault(r.client, ctx, req, resp, "Nodejs")
	}

}


func (r *AppNodejsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	appType := "Node.js"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpReadAppAWS(r.client, ctx, req, resp, "Nodejs")
	} else {
		common.ResourceOpReadAppDefault(r.client, ctx, req, resp, "Nodejs")
	}

}


func (r *AppNodejsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	appType := "Node.js"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpUpdateAppAWS(r.client, ctx, req, resp, "Nodejs")
	} else {
		common.ResourceOpUpdateAppDefault(r.client, ctx, req, resp, "Nodejs")
	}

}


func (r *AppNodejsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	appType := "Node.js"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpDeleteAppAWS(r.client, ctx, req, resp, "Nodejs")
	} else { 
		common.ResourceOpDeleteAppDefault(r.client, ctx, req, resp, "Nodejs")
	}
	

}


func (r *AppNodejsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	appType := "Node.js"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
  		common.ResourceOpImportAppAWS(r.client, ctx, req, resp, "Nodejs")
	} else {
		common.ResourceOpImportAppDefault(r.client, ctx, req, resp, "Nodejs")
	}

}

