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


type AppAwsec2Resource struct {
	client *stcloud.APIClient
}


// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppAwsec2Resource{}
var _ resource.ResourceWithImportState = &AppAwsec2Resource{}


func NewAppAwsec2Resource() resource.Resource {

	return &AppAwsec2Resource{}

}


func (r *AppAwsec2Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {

	resp.TypeName = "sematext_app_awsec2"	

}


func (r *AppAwsec2Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

	appType := "AWS EC2"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		resp.Schema = common.ResourceSchemaAppAWS("AWS EC2")
	} else{
		resp.Schema = common.ResourceSchemaAppDefault("AWS EC2")
	}

}


func (r *AppAwsec2Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	appType := "AWS EC2"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpConfigureAppAWS(r.client, ctx, req, resp, "Awsec2")
	} else {
		common.ResourceOpConfigureAppDefault(r.client, ctx, req, resp, "Awsec2")
	}

}


func (r *AppAwsec2Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	appType := "AWS EC2"
		
	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpCreateAppAWS(r.client, ctx, req, resp, "Awsec2")
	} else{
		common.ResourceOpCreateAppDefault(r.client, ctx, req, resp, "Awsec2")
	}

}


func (r *AppAwsec2Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	appType := "AWS EC2"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpReadAppAWS(r.client, ctx, req, resp, "Awsec2")
	} else {
		common.ResourceOpReadAppDefault(r.client, ctx, req, resp, "Awsec2")
	}

}


func (r *AppAwsec2Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	appType := "AWS EC2"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpUpdateAppAWS(r.client, ctx, req, resp, "Awsec2")
	} else {
		common.ResourceOpUpdateAppDefault(r.client, ctx, req, resp, "Awsec2")
	}

}


func (r *AppAwsec2Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	appType := "AWS EC2"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpDeleteAppAWS(r.client, ctx, req, resp, "Awsec2")
	} else { 
		common.ResourceOpDeleteAppDefault(r.client, ctx, req, resp, "Awsec2")
	}
	

}


func (r *AppAwsec2Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	appType := "AWS EC2"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
  		common.ResourceOpImportAppAWS(r.client, ctx, req, resp, "Awsec2")
	} else {
		common.ResourceOpImportAppDefault(r.client, ctx, req, resp, "Awsec2")
	}

}

