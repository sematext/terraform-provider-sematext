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
var _ resource.Resource = &AppSolrcloudResource{}
var _ resource.ResourceWithImportState = &AppSolrcloudResource{}


func NewAppSolrcloudResource() resource.Resource {
	return &AppSolrcloudResource{}
}

// AppSolrcloudResourceModel describes the resource data model.
type AppSolrcloudResourceModel = AppResourceModel


// AppSolrcloudResource defines the resource implementation.
type AppSolrcloudResource = AppResource


func (r *AppSolrcloudResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_AppSolrcloudResource"	
}


func (r *AppSolrcloudResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchemaApp("SolrCloud")
}


func (r *AppSolrcloudResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	ResourceOpConfigureApp(r, ctx, req, resp, "Solrcloud")
}


func (r *AppSolrcloudResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
		
	ResourceOpCreateApp(r, ctx, req, resp, "Solrcloud")

}


func (r *AppSolrcloudResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	ResourceOpReadApp(r, ctx, req, resp, "Solrcloud")

}

func (r *AppSolrcloudResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	ResourceOpUpdateApp(r, ctx, req, resp, "Solrcloud")

}

func (r *AppSolrcloudResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	ResourceOpDeleteApp(r, ctx, req, resp, "Solrcloud")

}

func (r *AppSolrcloudResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	ResourceOpImportApp(r, ctx, req, resp, "Solrcloud")

}

