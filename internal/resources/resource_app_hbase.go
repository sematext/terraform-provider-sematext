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


type AppHbaseResource struct {
	client *stcloud.APIClient
}


// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppHbaseResource{}
var _ resource.ResourceWithImportState = &AppHbaseResource{}


func NewAppHbaseResource() resource.Resource {

	return &AppHbaseResource{}

}


func (r *AppHbaseResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {

	resp.TypeName = "sematext_app_hbase"	

}


func (r *AppHbaseResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

	appType := "HBase"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		resp.Schema = common.ResourceSchemaAppAWS("HBase")
	} else{
		resp.Schema = common.ResourceSchemaAppDefault("HBase")
	}

}


func (r *AppHbaseResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	appType := "HBase"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpConfigureAppAWS(r.client, ctx, req, resp, "Hbase")
	} else {
		common.ResourceOpConfigureAppDefault(r.client, ctx, req, resp, "Hbase")
	}

}


func (r *AppHbaseResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	appType := "HBase"
		
	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpCreateAppAWS(r.client, ctx, req, resp, "Hbase")
	} else{
		common.ResourceOpCreateAppDefault(r.client, ctx, req, resp, "Hbase")
	}

}


func (r *AppHbaseResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	appType := "HBase"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpReadAppAWS(r.client, ctx, req, resp, "Hbase")
	} else {
		common.ResourceOpReadAppDefault(r.client, ctx, req, resp, "Hbase")
	}

}


func (r *AppHbaseResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	appType := "HBase"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpUpdateAppAWS(r.client, ctx, req, resp, "Hbase")
	} else {
		common.ResourceOpUpdateAppDefault(r.client, ctx, req, resp, "Hbase")
	}

}


func (r *AppHbaseResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	appType := "HBase"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpDeleteAppAWS(r.client, ctx, req, resp, "Hbase")
	} else { 
		common.ResourceOpDeleteAppDefault(r.client, ctx, req, resp, "Hbase")
	}
	

}


func (r *AppHbaseResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	appType := "HBase"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
  		common.ResourceOpImportAppAWS(r.client, ctx, req, resp, "Hbase")
	} else {
		common.ResourceOpImportAppDefault(r.client, ctx, req, resp, "Hbase")
	}

}

