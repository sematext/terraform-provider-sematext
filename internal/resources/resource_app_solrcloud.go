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


type AppSolrcloudResource struct {
	client *stcloud.APIClient
}


// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppSolrcloudResource{}
var _ resource.ResourceWithImportState = &AppSolrcloudResource{}


func NewAppSolrcloudResource() resource.Resource {

	return &AppSolrcloudResource{}

}


func (r *AppSolrcloudResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {

	resp.TypeName = "sematext_app_solrcloud"	

}


func (r *AppSolrcloudResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

	appType := "SolrCloud"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		resp.Schema = common.ResourceSchemaAppAWS("SolrCloud")
	} else{
		resp.Schema = common.ResourceSchemaAppDefault("SolrCloud")
	}

}


func (r *AppSolrcloudResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	appType := "SolrCloud"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpConfigureAppAWS(r.client, ctx, req, resp, "Solrcloud")
	} else {
		common.ResourceOpConfigureAppDefault(r.client, ctx, req, resp, "Solrcloud")
	}

}


func (r *AppSolrcloudResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	appType := "SolrCloud"
		
	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpCreateAppAWS(r.client, ctx, req, resp, "Solrcloud")
	} else{
		common.ResourceOpCreateAppDefault(r.client, ctx, req, resp, "Solrcloud")
	}

}


func (r *AppSolrcloudResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	appType := "SolrCloud"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpReadAppAWS(r.client, ctx, req, resp, "Solrcloud")
	} else {
		common.ResourceOpReadAppDefault(r.client, ctx, req, resp, "Solrcloud")
	}

}


func (r *AppSolrcloudResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	appType := "SolrCloud"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpUpdateAppAWS(r.client, ctx, req, resp, "Solrcloud")
	} else {
		common.ResourceOpUpdateAppDefault(r.client, ctx, req, resp, "Solrcloud")
	}

}


func (r *AppSolrcloudResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	appType := "SolrCloud"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpDeleteAppAWS(r.client, ctx, req, resp, "Solrcloud")
	} else { 
		common.ResourceOpDeleteAppDefault(r.client, ctx, req, resp, "Solrcloud")
	}
	

}


func (r *AppSolrcloudResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	appType := "SolrCloud"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
  		common.ResourceOpImportAppAWS(r.client, ctx, req, resp, "Solrcloud")
	} else {
		common.ResourceOpImportAppDefault(r.client, ctx, req, resp, "Solrcloud")
	}

}

