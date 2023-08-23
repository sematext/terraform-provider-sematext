package sematext

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_App.go.template
	Then run generate/generate.sh
*/

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"

)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppElasticsearchResource{}
var _ resource.ResourceWithImportState = &AppElasticsearchResource{}


func NewAppElasticsearchResource() resource.Resource {
	return &AppElasticsearchResource{}
}

// AppElasticsearchResourceModel describes the resource data model.
type AppElasticsearchResourceModel = AppResourceModel


// AppElasticsearchResource defines the resource implementation.
type AppElasticsearchResource = AppResource


func (r *AppElasticsearchResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_AppElasticsearchResource"	
}


func (r *AppElasticsearchResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchemaApp("Elastic Search")
}


func (r *AppElasticsearchResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	ResourceOpConfigureApp(r, ctx, req, resp, "Elasticsearch")
}


func (r *AppElasticsearchResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
		
	ResourceOpCreateApp(r, ctx, req, resp, "Elasticsearch")

}


func (r *AppElasticsearchResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	ResourceOpReadApp(r, ctx, req, resp, "Elasticsearch")

}

func (r *AppElasticsearchResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	ResourceOpUpdateApp(r, ctx, req, resp, "Elasticsearch")

}

func (r *AppElasticsearchResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	ResourceOpDeleteApp(r, ctx, req, resp, "Elasticsearch")

}

func (r *AppElasticsearchResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	ResourceOpImportApp(r, ctx, req, resp, "Elasticsearch")

}

