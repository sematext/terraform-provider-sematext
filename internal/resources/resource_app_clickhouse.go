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


type AppClickhouseResource struct {
	client *stcloud.APIClient
}


// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppClickhouseResource{}
var _ resource.ResourceWithImportState = &AppClickhouseResource{}


func NewAppClickhouseResource() resource.Resource {

	return &AppClickhouseResource{}

}


func (r *AppClickhouseResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {

	resp.TypeName = "sematext_app_clickhouse"	

}


func (r *AppClickhouseResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

	appType := "ClickHouse"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		resp.Schema = common.ResourceSchemaAppAWS("ClickHouse")
	} else{
		resp.Schema = common.ResourceSchemaAppDefault("ClickHouse")
	}

}


func (r *AppClickhouseResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	appType := "ClickHouse"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpConfigureAppAWS(r.client, ctx, req, resp, "Clickhouse")
	} else {
		common.ResourceOpConfigureAppDefault(r.client, ctx, req, resp, "Clickhouse")
	}

}


func (r *AppClickhouseResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	appType := "ClickHouse"
		
	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpCreateAppAWS(r.client, ctx, req, resp, "Clickhouse")
	} else{
		common.ResourceOpCreateAppDefault(r.client, ctx, req, resp, "Clickhouse")
	}

}


func (r *AppClickhouseResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	appType := "ClickHouse"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpReadAppAWS(r.client, ctx, req, resp, "Clickhouse")
	} else {
		common.ResourceOpReadAppDefault(r.client, ctx, req, resp, "Clickhouse")
	}

}


func (r *AppClickhouseResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	appType := "ClickHouse"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpUpdateAppAWS(r.client, ctx, req, resp, "Clickhouse")
	} else {
		common.ResourceOpUpdateAppDefault(r.client, ctx, req, resp, "Clickhouse")
	}

}


func (r *AppClickhouseResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	appType := "ClickHouse"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpDeleteAppAWS(r.client, ctx, req, resp, "Clickhouse")
	} else { 
		common.ResourceOpDeleteAppDefault(r.client, ctx, req, resp, "Clickhouse")
	}
	

}


func (r *AppClickhouseResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	appType := "ClickHouse"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
  		common.ResourceOpImportAppAWS(r.client, ctx, req, resp, "Clickhouse")
	} else {
		common.ResourceOpImportAppDefault(r.client, ctx, req, resp, "Clickhouse")
	}

}

