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


type AppPostgresqlResource struct {
	client *stcloud.APIClient
}


// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppPostgresqlResource{}
var _ resource.ResourceWithImportState = &AppPostgresqlResource{}


func NewAppPostgresqlResource() resource.Resource {

	return &AppPostgresqlResource{}

}


func (r *AppPostgresqlResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {

	resp.TypeName = "sematext_app_postgresql"	

}


func (r *AppPostgresqlResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

	appType := "postgresql"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		resp.Schema = common.ResourceSchemaAppAWS("postgresql")
	} else{
		resp.Schema = common.ResourceSchemaAppDefault("postgresql")
	}

}


func (r *AppPostgresqlResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	appType := "postgresql"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpConfigureAppAWS(r.client, ctx, req, resp, "Postgresql")
	} else {
		common.ResourceOpConfigureAppDefault(r.client, ctx, req, resp, "Postgresql")
	}

}


func (r *AppPostgresqlResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	appType := "postgresql"
		
	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpCreateAppAWS(r.client, ctx, req, resp, "Postgresql")
	} else{
		common.ResourceOpCreateAppDefault(r.client, ctx, req, resp, "Postgresql")
	}

}


func (r *AppPostgresqlResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	appType := "postgresql"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpReadAppAWS(r.client, ctx, req, resp, "Postgresql")
	} else {
		common.ResourceOpReadAppDefault(r.client, ctx, req, resp, "Postgresql")
	}

}


func (r *AppPostgresqlResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	appType := "postgresql"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpUpdateAppAWS(r.client, ctx, req, resp, "Postgresql")
	} else {
		common.ResourceOpUpdateAppDefault(r.client, ctx, req, resp, "Postgresql")
	}

}


func (r *AppPostgresqlResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	appType := "postgresql"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpDeleteAppAWS(r.client, ctx, req, resp, "Postgresql")
	} else { 
		common.ResourceOpDeleteAppDefault(r.client, ctx, req, resp, "Postgresql")
	}
	

}


func (r *AppPostgresqlResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	appType := "postgresql"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
  		common.ResourceOpImportAppAWS(r.client, ctx, req, resp, "Postgresql")
	} else {
		common.ResourceOpImportAppDefault(r.client, ctx, req, resp, "Postgresql")
	}

}

