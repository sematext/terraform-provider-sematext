package storm

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_App.go.template
	Then run generate/generate.sh
*/

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/sematext/sematext-api-client-go/stcloud"
	"github.com/sematext/terraform-provider-sematext/internal/operation"
	"github.com/sematext/terraform-provider-sematext/internal/common"
)


type Resource struct {
	client *stcloud.APIClient
}


// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &Resource{}
var _ resource.ResourceWithImportState = &Resource{}


func NewResource() resource.Resource {

	return &Resource{}

}


func (r *Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {

	resp.TypeName = "sematext_app_storm"	

}


func (r *Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

	appType := "Storm"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		resp.Schema = common.ResourceSchemaAWS("Storm")
	} else{
		resp.Schema = common.ResourceSchemaDefault("Storm")
	}

}


func (r *Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	appType := "Storm"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		operation.ResourceOpConfigureAWS(r.client, ctx, req, resp, "Storm")
	} else {
		operation.ResourceOpConfigureDefault(r.client, ctx, req, resp, "Storm")
	}

}


func (r *Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	appType := "Storm"
		
	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		operation.ResourceOpCreateAWS(r.client, ctx, req, resp, "Storm")
	} else{
		operation.ResourceOpCreateDefault(r.client, ctx, req, resp, "Storm")
	}

}


func (r *Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	appType := "Storm"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		operation.ResourceOpReadAWS(r.client, ctx, req, resp, "Storm")
	} else {
		operation.ResourceOpReadDefault(r.client, ctx, req, resp, "Storm")
	}

}


func (r *Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	appType := "Storm"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		operation.ResourceOpUpdateAWS(r.client, ctx, req, resp, "Storm")
	} else {
		operation.ResourceOpUpdateDefault(r.client, ctx, req, resp, "Storm")
	}

}


func (r *Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	appType := "Storm"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		operation.ResourceOpDeleteAWS(r.client, ctx, req, resp, "Storm")
	} else { 
		operation.ResourceOpDeleteDefault(r.client, ctx, req, resp, "Storm")
	}
	

}


func (r *Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	appType := "Storm"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
  		operation.ResourceOpImportAWS(r.client, ctx, req, resp, "Storm")
	} else {
		operation.ResourceOpImportDefault(r.client, ctx, req, resp, "Storm")
	}

}

