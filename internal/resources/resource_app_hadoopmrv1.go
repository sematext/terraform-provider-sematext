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


type AppHadoopmrv1Resource struct {
	client *stcloud.APIClient
}


// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppHadoopmrv1Resource{}
var _ resource.ResourceWithImportState = &AppHadoopmrv1Resource{}


func NewAppHadoopmrv1Resource() resource.Resource {

	return &AppHadoopmrv1Resource{}

}


func (r *AppHadoopmrv1Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {

	resp.TypeName = "sematext_app_hadoopmrv1"	

}


func (r *AppHadoopmrv1Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

	appType := "Hadoop-MRv1"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		resp.Schema = common.ResourceSchemaAppAWS("Hadoop-MRv1")
	} else{
		resp.Schema = common.ResourceSchemaAppDefault("Hadoop-MRv1")
	}

}


func (r *AppHadoopmrv1Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	appType := "Hadoop-MRv1"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpConfigureAppAWS(r.client, ctx, req, resp, "Hadoopmrv1")
	} else {
		common.ResourceOpConfigureAppDefault(r.client, ctx, req, resp, "Hadoopmrv1")
	}

}


func (r *AppHadoopmrv1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	appType := "Hadoop-MRv1"
		
	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpCreateAppAWS(r.client, ctx, req, resp, "Hadoopmrv1")
	} else{
		common.ResourceOpCreateAppDefault(r.client, ctx, req, resp, "Hadoopmrv1")
	}

}


func (r *AppHadoopmrv1Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	appType := "Hadoop-MRv1"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpReadAppAWS(r.client, ctx, req, resp, "Hadoopmrv1")
	} else {
		common.ResourceOpReadAppDefault(r.client, ctx, req, resp, "Hadoopmrv1")
	}

}


func (r *AppHadoopmrv1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	appType := "Hadoop-MRv1"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpUpdateAppAWS(r.client, ctx, req, resp, "Hadoopmrv1")
	} else {
		common.ResourceOpUpdateAppDefault(r.client, ctx, req, resp, "Hadoopmrv1")
	}

}


func (r *AppHadoopmrv1Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	appType := "Hadoop-MRv1"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpDeleteAppAWS(r.client, ctx, req, resp, "Hadoopmrv1")
	} else { 
		common.ResourceOpDeleteAppDefault(r.client, ctx, req, resp, "Hadoopmrv1")
	}
	

}


func (r *AppHadoopmrv1Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	appType := "Hadoop-MRv1"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
  		common.ResourceOpImportAppAWS(r.client, ctx, req, resp, "Hadoopmrv1")
	} else {
		common.ResourceOpImportAppDefault(r.client, ctx, req, resp, "Hadoopmrv1")
	}

}

