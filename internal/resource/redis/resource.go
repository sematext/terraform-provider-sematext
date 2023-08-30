package redis

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

	resp.TypeName = "sematext_app_redis"	

}


func (r *Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

	appType := "Redis"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		resp.Schema = common.ResourceSchemaAWS("Redis")
	} else{
		resp.Schema = common.ResourceSchemaDefault("Redis")
	}

}


func (r *Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	appType := "Redis"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpConfigureAWS(r.client, ctx, req, resp, "Redis")
	} else {
		common.ResourceOpConfigureDefault(r.client, ctx, req, resp, "Redis")
	}

}


func (r *Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	appType := "Redis"
		
	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpCreateAWS(r.client, ctx, req, resp, "Redis")
	} else{
		common.ResourceOpCreateDefault(r.client, ctx, req, resp, "Redis")
	}

}


func (r *Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	appType := "Redis"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpReadAWS(r.client, ctx, req, resp, "Redis")
	} else {
		common.ResourceOpReadDefault(r.client, ctx, req, resp, "Redis")
	}

}


func (r *Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	appType := "Redis"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpUpdateAWS(r.client, ctx, req, resp, "Redis")
	} else {
		common.ResourceOpUpdateDefault(r.client, ctx, req, resp, "Redis")
	}

}


func (r *Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	appType := "Redis"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpDeleteAWS(r.client, ctx, req, resp, "Redis")
	} else { 
		common.ResourceOpDeleteDefault(r.client, ctx, req, resp, "Redis")
	}
	

}


func (r *Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	appType := "Redis"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
  		common.ResourceOpImportAWS(r.client, ctx, req, resp, "Redis")
	} else {
		common.ResourceOpImportDefault(r.client, ctx, req, resp, "Redis")
	}

}

