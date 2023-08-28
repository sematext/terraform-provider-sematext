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


type AppApacheResource struct {
	client *stcloud.APIClient
}


// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppApacheResource{}
var _ resource.ResourceWithImportState = &AppApacheResource{}


func NewAppApacheResource() resource.Resource {

	return &AppApacheResource{}

}


func (r *AppApacheResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {

	resp.TypeName = "sematext_app_apache"	

}


func (r *AppApacheResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

	appType := "Apache"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		resp.Schema = common.ResourceSchemaAppAWS("Apache")
	} else{
		resp.Schema = common.ResourceSchemaAppDefault("Apache")
	}

}


func (r *AppApacheResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	appType := "Apache"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpConfigureAppAWS(r.client, ctx, req, resp, "Apache")
	} else {
		common.ResourceOpConfigureAppDefault(r.client, ctx, req, resp, "Apache")
	}

}


func (r *AppApacheResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	appType := "Apache"
		
	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpCreateAppAWS(r.client, ctx, req, resp, "Apache")
	} else{
		common.ResourceOpCreateAppDefault(r.client, ctx, req, resp, "Apache")
	}

}


func (r *AppApacheResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	appType := "Apache"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpReadAppAWS(r.client, ctx, req, resp, "Apache")
	} else {
		common.ResourceOpReadAppDefault(r.client, ctx, req, resp, "Apache")
	}

}


func (r *AppApacheResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	appType := "Apache"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpUpdateAppAWS(r.client, ctx, req, resp, "Apache")
	} else {
		common.ResourceOpUpdateAppDefault(r.client, ctx, req, resp, "Apache")
	}

}


func (r *AppApacheResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	appType := "Apache"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpDeleteAppAWS(r.client, ctx, req, resp, "Apache")
	} else { 
		common.ResourceOpDeleteAppDefault(r.client, ctx, req, resp, "Apache")
	}
	

}


func (r *AppApacheResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	appType := "Apache"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
  		common.ResourceOpImportAppAWS(r.client, ctx, req, resp, "Apache")
	} else {
		common.ResourceOpImportAppDefault(r.client, ctx, req, resp, "Apache")
	}

}

