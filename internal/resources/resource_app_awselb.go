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


type AppAwselbResource struct {
	client *stcloud.APIClient
}


// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppAwselbResource{}
var _ resource.ResourceWithImportState = &AppAwselbResource{}


func NewAppAwselbResource() resource.Resource {

	return &AppAwselbResource{}

}


func (r *AppAwselbResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {

	resp.TypeName = "sematext_app_awselb"	

}


func (r *AppAwselbResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

	appType := "AWS ELB"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		resp.Schema = common.ResourceSchemaAppAWS("AWS ELB")
	} else{
		resp.Schema = common.ResourceSchemaAppDefault("AWS ELB")
	}

}


func (r *AppAwselbResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	appType := "AWS ELB"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpConfigureAppAWS(r.client, ctx, req, resp, "Awselb")
	} else {
		common.ResourceOpConfigureAppDefault(r.client, ctx, req, resp, "Awselb")
	}

}


func (r *AppAwselbResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	appType := "AWS ELB"
		
	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpCreateAppAWS(r.client, ctx, req, resp, "Awselb")
	} else{
		common.ResourceOpCreateAppDefault(r.client, ctx, req, resp, "Awselb")
	}

}


func (r *AppAwselbResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	appType := "AWS ELB"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpReadAppAWS(r.client, ctx, req, resp, "Awselb")
	} else {
		common.ResourceOpReadAppDefault(r.client, ctx, req, resp, "Awselb")
	}

}


func (r *AppAwselbResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	appType := "AWS ELB"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpUpdateAppAWS(r.client, ctx, req, resp, "Awselb")
	} else {
		common.ResourceOpUpdateAppDefault(r.client, ctx, req, resp, "Awselb")
	}

}


func (r *AppAwselbResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	appType := "AWS ELB"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpDeleteAppAWS(r.client, ctx, req, resp, "Awselb")
	} else { 
		common.ResourceOpDeleteAppDefault(r.client, ctx, req, resp, "Awselb")
	}
	

}


func (r *AppAwselbResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	appType := "AWS ELB"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
  		common.ResourceOpImportAppAWS(r.client, ctx, req, resp, "Awselb")
	} else {
		common.ResourceOpImportAppDefault(r.client, ctx, req, resp, "Awselb")
	}

}

