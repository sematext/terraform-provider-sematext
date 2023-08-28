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


type AppElasticsearchResource struct {
	client *stcloud.APIClient
}


// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppElasticsearchResource{}
var _ resource.ResourceWithImportState = &AppElasticsearchResource{}


func NewAppElasticsearchResource() resource.Resource {

	return &AppElasticsearchResource{}

}


func (r *AppElasticsearchResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {

	resp.TypeName = "sematext_app_elasticsearch"	

}


func (r *AppElasticsearchResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

	appType := "Elastic Search"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		resp.Schema = common.ResourceSchemaAppAWS("Elastic Search")
	} else{
		resp.Schema = common.ResourceSchemaAppDefault("Elastic Search")
	}

}


func (r *AppElasticsearchResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	appType := "Elastic Search"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpConfigureAppAWS(r.client, ctx, req, resp, "Elasticsearch")
	} else {
		common.ResourceOpConfigureAppDefault(r.client, ctx, req, resp, "Elasticsearch")
	}

}


func (r *AppElasticsearchResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	appType := "Elastic Search"
		
	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpCreateAppAWS(r.client, ctx, req, resp, "Elasticsearch")
	} else{
		common.ResourceOpCreateAppDefault(r.client, ctx, req, resp, "Elasticsearch")
	}

}


func (r *AppElasticsearchResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	appType := "Elastic Search"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpReadAppAWS(r.client, ctx, req, resp, "Elasticsearch")
	} else {
		common.ResourceOpReadAppDefault(r.client, ctx, req, resp, "Elasticsearch")
	}

}


func (r *AppElasticsearchResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	appType := "Elastic Search"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpUpdateAppAWS(r.client, ctx, req, resp, "Elasticsearch")
	} else {
		common.ResourceOpUpdateAppDefault(r.client, ctx, req, resp, "Elasticsearch")
	}

}


func (r *AppElasticsearchResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	appType := "Elastic Search"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpDeleteAppAWS(r.client, ctx, req, resp, "Elasticsearch")
	} else { 
		common.ResourceOpDeleteAppDefault(r.client, ctx, req, resp, "Elasticsearch")
	}
	

}


func (r *AppElasticsearchResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	appType := "Elastic Search"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
  		common.ResourceOpImportAppAWS(r.client, ctx, req, resp, "Elasticsearch")
	} else {
		common.ResourceOpImportAppDefault(r.client, ctx, req, resp, "Elasticsearch")
	}

}

