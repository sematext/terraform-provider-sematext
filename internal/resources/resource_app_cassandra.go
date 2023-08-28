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


type AppCassandraResource struct {
	client *stcloud.APIClient
}


// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppCassandraResource{}
var _ resource.ResourceWithImportState = &AppCassandraResource{}


func NewAppCassandraResource() resource.Resource {

	return &AppCassandraResource{}

}


func (r *AppCassandraResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {

	resp.TypeName = "sematext_app_cassandra"	

}


func (r *AppCassandraResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

	appType := "Cassandra"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		resp.Schema = common.ResourceSchemaAppAWS("Cassandra")
	} else{
		resp.Schema = common.ResourceSchemaAppDefault("Cassandra")
	}

}


func (r *AppCassandraResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	appType := "Cassandra"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpConfigureAppAWS(r.client, ctx, req, resp, "Cassandra")
	} else {
		common.ResourceOpConfigureAppDefault(r.client, ctx, req, resp, "Cassandra")
	}

}


func (r *AppCassandraResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	appType := "Cassandra"
		
	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpCreateAppAWS(r.client, ctx, req, resp, "Cassandra")
	} else{
		common.ResourceOpCreateAppDefault(r.client, ctx, req, resp, "Cassandra")
	}

}


func (r *AppCassandraResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	appType := "Cassandra"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpReadAppAWS(r.client, ctx, req, resp, "Cassandra")
	} else {
		common.ResourceOpReadAppDefault(r.client, ctx, req, resp, "Cassandra")
	}

}


func (r *AppCassandraResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	appType := "Cassandra"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpUpdateAppAWS(r.client, ctx, req, resp, "Cassandra")
	} else {
		common.ResourceOpUpdateAppDefault(r.client, ctx, req, resp, "Cassandra")
	}

}


func (r *AppCassandraResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	appType := "Cassandra"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpDeleteAppAWS(r.client, ctx, req, resp, "Cassandra")
	} else { 
		common.ResourceOpDeleteAppDefault(r.client, ctx, req, resp, "Cassandra")
	}
	

}


func (r *AppCassandraResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	appType := "Cassandra"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
  		common.ResourceOpImportAppAWS(r.client, ctx, req, resp, "Cassandra")
	} else {
		common.ResourceOpImportAppDefault(r.client, ctx, req, resp, "Cassandra")
	}

}

