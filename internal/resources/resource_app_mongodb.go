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


type AppMongodbResource struct {
	client *stcloud.APIClient
}


// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppMongodbResource{}
var _ resource.ResourceWithImportState = &AppMongodbResource{}


func NewAppMongodbResource() resource.Resource {

	return &AppMongodbResource{}

}


func (r *AppMongodbResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {

	resp.TypeName = "sematext_app_mongodb"	

}


func (r *AppMongodbResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

	appType := "MongoDB"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		resp.Schema = common.ResourceSchemaAppAWS("MongoDB")
	} else{
		resp.Schema = common.ResourceSchemaAppDefault("MongoDB")
	}

}


func (r *AppMongodbResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	appType := "MongoDB"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpConfigureAppAWS(r.client, ctx, req, resp, "Mongodb")
	} else {
		common.ResourceOpConfigureAppDefault(r.client, ctx, req, resp, "Mongodb")
	}

}


func (r *AppMongodbResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	appType := "MongoDB"
		
	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpCreateAppAWS(r.client, ctx, req, resp, "Mongodb")
	} else{
		common.ResourceOpCreateAppDefault(r.client, ctx, req, resp, "Mongodb")
	}

}


func (r *AppMongodbResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	appType := "MongoDB"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpReadAppAWS(r.client, ctx, req, resp, "Mongodb")
	} else {
		common.ResourceOpReadAppDefault(r.client, ctx, req, resp, "Mongodb")
	}

}


func (r *AppMongodbResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	appType := "MongoDB"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpUpdateAppAWS(r.client, ctx, req, resp, "Mongodb")
	} else {
		common.ResourceOpUpdateAppDefault(r.client, ctx, req, resp, "Mongodb")
	}

}


func (r *AppMongodbResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	appType := "MongoDB"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpDeleteAppAWS(r.client, ctx, req, resp, "Mongodb")
	} else { 
		common.ResourceOpDeleteAppDefault(r.client, ctx, req, resp, "Mongodb")
	}
	

}


func (r *AppMongodbResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	appType := "MongoDB"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
  		common.ResourceOpImportAppAWS(r.client, ctx, req, resp, "Mongodb")
	} else {
		common.ResourceOpImportAppDefault(r.client, ctx, req, resp, "Mongodb")
	}

}

