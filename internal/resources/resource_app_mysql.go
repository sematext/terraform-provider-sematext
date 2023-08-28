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


type AppMysqlResource struct {
	client *stcloud.APIClient
}


// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppMysqlResource{}
var _ resource.ResourceWithImportState = &AppMysqlResource{}


func NewAppMysqlResource() resource.Resource {

	return &AppMysqlResource{}

}


func (r *AppMysqlResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {

	resp.TypeName = "sematext_app_mysql"	

}


func (r *AppMysqlResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

	appType := "MySQL"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		resp.Schema = common.ResourceSchemaAppAWS("MySQL")
	} else{
		resp.Schema = common.ResourceSchemaAppDefault("MySQL")
	}

}


func (r *AppMysqlResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	appType := "MySQL"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpConfigureAppAWS(r.client, ctx, req, resp, "Mysql")
	} else {
		common.ResourceOpConfigureAppDefault(r.client, ctx, req, resp, "Mysql")
	}

}


func (r *AppMysqlResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	appType := "MySQL"
		
	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpCreateAppAWS(r.client, ctx, req, resp, "Mysql")
	} else{
		common.ResourceOpCreateAppDefault(r.client, ctx, req, resp, "Mysql")
	}

}


func (r *AppMysqlResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	appType := "MySQL"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpReadAppAWS(r.client, ctx, req, resp, "Mysql")
	} else {
		common.ResourceOpReadAppDefault(r.client, ctx, req, resp, "Mysql")
	}

}


func (r *AppMysqlResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	appType := "MySQL"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpUpdateAppAWS(r.client, ctx, req, resp, "Mysql")
	} else {
		common.ResourceOpUpdateAppDefault(r.client, ctx, req, resp, "Mysql")
	}

}


func (r *AppMysqlResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	appType := "MySQL"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpDeleteAppAWS(r.client, ctx, req, resp, "Mysql")
	} else { 
		common.ResourceOpDeleteAppDefault(r.client, ctx, req, resp, "Mysql")
	}
	

}


func (r *AppMysqlResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	appType := "MySQL"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
  		common.ResourceOpImportAppAWS(r.client, ctx, req, resp, "Mysql")
	} else {
		common.ResourceOpImportAppDefault(r.client, ctx, req, resp, "Mysql")
	}

}

