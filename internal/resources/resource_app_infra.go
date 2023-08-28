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


type AppInfraResource struct {
	client *stcloud.APIClient
}


// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppInfraResource{}
var _ resource.ResourceWithImportState = &AppInfraResource{}


func NewAppInfraResource() resource.Resource {

	return &AppInfraResource{}

}


func (r *AppInfraResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {

	resp.TypeName = "sematext_app_infra"	

}


func (r *AppInfraResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

	appType := "Infra"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		resp.Schema = common.ResourceSchemaAppAWS("Infra")
	} else{
		resp.Schema = common.ResourceSchemaAppDefault("Infra")
	}

}


func (r *AppInfraResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	appType := "Infra"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpConfigureAppAWS(r.client, ctx, req, resp, "Infra")
	} else {
		common.ResourceOpConfigureAppDefault(r.client, ctx, req, resp, "Infra")
	}

}


func (r *AppInfraResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	appType := "Infra"
		
	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpCreateAppAWS(r.client, ctx, req, resp, "Infra")
	} else{
		common.ResourceOpCreateAppDefault(r.client, ctx, req, resp, "Infra")
	}

}


func (r *AppInfraResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	appType := "Infra"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpReadAppAWS(r.client, ctx, req, resp, "Infra")
	} else {
		common.ResourceOpReadAppDefault(r.client, ctx, req, resp, "Infra")
	}

}


func (r *AppInfraResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	appType := "Infra"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpUpdateAppAWS(r.client, ctx, req, resp, "Infra")
	} else {
		common.ResourceOpUpdateAppDefault(r.client, ctx, req, resp, "Infra")
	}

}


func (r *AppInfraResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	appType := "Infra"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpDeleteAppAWS(r.client, ctx, req, resp, "Infra")
	} else { 
		common.ResourceOpDeleteAppDefault(r.client, ctx, req, resp, "Infra")
	}
	

}


func (r *AppInfraResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	appType := "Infra"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
  		common.ResourceOpImportAppAWS(r.client, ctx, req, resp, "Infra")
	} else {
		common.ResourceOpImportAppDefault(r.client, ctx, req, resp, "Infra")
	}

}

