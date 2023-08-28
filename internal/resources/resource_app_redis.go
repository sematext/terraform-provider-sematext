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


type AppRedisResource struct {
	client *stcloud.APIClient
}


// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppRedisResource{}
var _ resource.ResourceWithImportState = &AppRedisResource{}


func NewAppRedisResource() resource.Resource {

	return &AppRedisResource{}

}


func (r *AppRedisResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {

	resp.TypeName = "sematext_app_redis"	

}


func (r *AppRedisResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

	appType := "Redis"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		resp.Schema = common.ResourceSchemaAppAWS("Redis")
	} else{
		resp.Schema = common.ResourceSchemaAppDefault("Redis")
	}

}


func (r *AppRedisResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	appType := "Redis"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpConfigureAppAWS(r.client, ctx, req, resp, "Redis")
	} else {
		common.ResourceOpConfigureAppDefault(r.client, ctx, req, resp, "Redis")
	}

}


func (r *AppRedisResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	appType := "Redis"
		
	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpCreateAppAWS(r.client, ctx, req, resp, "Redis")
	} else{
		common.ResourceOpCreateAppDefault(r.client, ctx, req, resp, "Redis")
	}

}


func (r *AppRedisResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	appType := "Redis"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpReadAppAWS(r.client, ctx, req, resp, "Redis")
	} else {
		common.ResourceOpReadAppDefault(r.client, ctx, req, resp, "Redis")
	}

}


func (r *AppRedisResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	appType := "Redis"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpUpdateAppAWS(r.client, ctx, req, resp, "Redis")
	} else {
		common.ResourceOpUpdateAppDefault(r.client, ctx, req, resp, "Redis")
	}

}


func (r *AppRedisResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	appType := "Redis"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpDeleteAppAWS(r.client, ctx, req, resp, "Redis")
	} else { 
		common.ResourceOpDeleteAppDefault(r.client, ctx, req, resp, "Redis")
	}
	

}


func (r *AppRedisResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	appType := "Redis"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
  		common.ResourceOpImportAppAWS(r.client, ctx, req, resp, "Redis")
	} else {
		common.ResourceOpImportAppDefault(r.client, ctx, req, resp, "Redis")
	}

}

