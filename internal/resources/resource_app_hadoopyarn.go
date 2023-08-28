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


type AppHadoopyarnResource struct {
	client *stcloud.APIClient
}


// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppHadoopyarnResource{}
var _ resource.ResourceWithImportState = &AppHadoopyarnResource{}


func NewAppHadoopyarnResource() resource.Resource {

	return &AppHadoopyarnResource{}

}


func (r *AppHadoopyarnResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {

	resp.TypeName = "sematext_app_hadoopyarn"	

}


func (r *AppHadoopyarnResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

	appType := "Hadoop-YARN"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		resp.Schema = common.ResourceSchemaAppAWS("Hadoop-YARN")
	} else{
		resp.Schema = common.ResourceSchemaAppDefault("Hadoop-YARN")
	}

}


func (r *AppHadoopyarnResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	appType := "Hadoop-YARN"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpConfigureAppAWS(r.client, ctx, req, resp, "Hadoopyarn")
	} else {
		common.ResourceOpConfigureAppDefault(r.client, ctx, req, resp, "Hadoopyarn")
	}

}


func (r *AppHadoopyarnResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	appType := "Hadoop-YARN"
		
	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpCreateAppAWS(r.client, ctx, req, resp, "Hadoopyarn")
	} else{
		common.ResourceOpCreateAppDefault(r.client, ctx, req, resp, "Hadoopyarn")
	}

}


func (r *AppHadoopyarnResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	appType := "Hadoop-YARN"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpReadAppAWS(r.client, ctx, req, resp, "Hadoopyarn")
	} else {
		common.ResourceOpReadAppDefault(r.client, ctx, req, resp, "Hadoopyarn")
	}

}


func (r *AppHadoopyarnResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	appType := "Hadoop-YARN"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpUpdateAppAWS(r.client, ctx, req, resp, "Hadoopyarn")
	} else {
		common.ResourceOpUpdateAppDefault(r.client, ctx, req, resp, "Hadoopyarn")
	}

}


func (r *AppHadoopyarnResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	appType := "Hadoop-YARN"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpDeleteAppAWS(r.client, ctx, req, resp, "Hadoopyarn")
	} else { 
		common.ResourceOpDeleteAppDefault(r.client, ctx, req, resp, "Hadoopyarn")
	}
	

}


func (r *AppHadoopyarnResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	appType := "Hadoop-YARN"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
  		common.ResourceOpImportAppAWS(r.client, ctx, req, resp, "Hadoopyarn")
	} else {
		common.ResourceOpImportAppDefault(r.client, ctx, req, resp, "Hadoopyarn")
	}

}

