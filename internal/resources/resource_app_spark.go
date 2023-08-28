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


type AppSparkResource struct {
	client *stcloud.APIClient
}


// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppSparkResource{}
var _ resource.ResourceWithImportState = &AppSparkResource{}


func NewAppSparkResource() resource.Resource {

	return &AppSparkResource{}

}


func (r *AppSparkResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {

	resp.TypeName = "sematext_app_spark"	

}


func (r *AppSparkResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

	appType := "Spark"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		resp.Schema = common.ResourceSchemaAppAWS("Spark")
	} else{
		resp.Schema = common.ResourceSchemaAppDefault("Spark")
	}

}


func (r *AppSparkResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	appType := "Spark"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpConfigureAppAWS(r.client, ctx, req, resp, "Spark")
	} else {
		common.ResourceOpConfigureAppDefault(r.client, ctx, req, resp, "Spark")
	}

}


func (r *AppSparkResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	appType := "Spark"
		
	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpCreateAppAWS(r.client, ctx, req, resp, "Spark")
	} else{
		common.ResourceOpCreateAppDefault(r.client, ctx, req, resp, "Spark")
	}

}


func (r *AppSparkResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	appType := "Spark"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpReadAppAWS(r.client, ctx, req, resp, "Spark")
	} else {
		common.ResourceOpReadAppDefault(r.client, ctx, req, resp, "Spark")
	}

}


func (r *AppSparkResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	appType := "Spark"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpUpdateAppAWS(r.client, ctx, req, resp, "Spark")
	} else {
		common.ResourceOpUpdateAppDefault(r.client, ctx, req, resp, "Spark")
	}

}


func (r *AppSparkResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	appType := "Spark"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpDeleteAppAWS(r.client, ctx, req, resp, "Spark")
	} else { 
		common.ResourceOpDeleteAppDefault(r.client, ctx, req, resp, "Spark")
	}
	

}


func (r *AppSparkResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	appType := "Spark"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
  		common.ResourceOpImportAppAWS(r.client, ctx, req, resp, "Spark")
	} else {
		common.ResourceOpImportAppDefault(r.client, ctx, req, resp, "Spark")
	}

}

