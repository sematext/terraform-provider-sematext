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


type AppSolrResource struct {
	client *stcloud.APIClient
}


// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppSolrResource{}
var _ resource.ResourceWithImportState = &AppSolrResource{}


func NewAppSolrResource() resource.Resource {

	return &AppSolrResource{}

}


func (r *AppSolrResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {

	resp.TypeName = "sematext_app_solr"	

}


func (r *AppSolrResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

	appType := "Solr"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		resp.Schema = common.ResourceSchemaAppAWS("Solr")
	} else{
		resp.Schema = common.ResourceSchemaAppDefault("Solr")
	}

}


func (r *AppSolrResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	appType := "Solr"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpConfigureAppAWS(r.client, ctx, req, resp, "Solr")
	} else {
		common.ResourceOpConfigureAppDefault(r.client, ctx, req, resp, "Solr")
	}

}


func (r *AppSolrResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	appType := "Solr"
		
	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpCreateAppAWS(r.client, ctx, req, resp, "Solr")
	} else{
		common.ResourceOpCreateAppDefault(r.client, ctx, req, resp, "Solr")
	}

}


func (r *AppSolrResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	appType := "Solr"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpReadAppAWS(r.client, ctx, req, resp, "Solr")
	} else {
		common.ResourceOpReadAppDefault(r.client, ctx, req, resp, "Solr")
	}

}


func (r *AppSolrResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	appType := "Solr"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpUpdateAppAWS(r.client, ctx, req, resp, "Solr")
	} else {
		common.ResourceOpUpdateAppDefault(r.client, ctx, req, resp, "Solr")
	}

}


func (r *AppSolrResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	appType := "Solr"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpDeleteAppAWS(r.client, ctx, req, resp, "Solr")
	} else { 
		common.ResourceOpDeleteAppDefault(r.client, ctx, req, resp, "Solr")
	}
	

}


func (r *AppSolrResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	appType := "Solr"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
  		common.ResourceOpImportAppAWS(r.client, ctx, req, resp, "Solr")
	} else {
		common.ResourceOpImportAppDefault(r.client, ctx, req, resp, "Solr")
	}

}

