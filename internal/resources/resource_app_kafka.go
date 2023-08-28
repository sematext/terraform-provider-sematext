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


type AppKafkaResource struct {
	client *stcloud.APIClient
}


// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppKafkaResource{}
var _ resource.ResourceWithImportState = &AppKafkaResource{}


func NewAppKafkaResource() resource.Resource {

	return &AppKafkaResource{}

}


func (r *AppKafkaResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {

	resp.TypeName = "sematext_app_kafka"	

}


func (r *AppKafkaResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

	appType := "Kafka"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		resp.Schema = common.ResourceSchemaAppAWS("Kafka")
	} else{
		resp.Schema = common.ResourceSchemaAppDefault("Kafka")
	}

}


func (r *AppKafkaResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	appType := "Kafka"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpConfigureAppAWS(r.client, ctx, req, resp, "Kafka")
	} else {
		common.ResourceOpConfigureAppDefault(r.client, ctx, req, resp, "Kafka")
	}

}


func (r *AppKafkaResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	appType := "Kafka"
		
	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpCreateAppAWS(r.client, ctx, req, resp, "Kafka")
	} else{
		common.ResourceOpCreateAppDefault(r.client, ctx, req, resp, "Kafka")
	}

}


func (r *AppKafkaResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	appType := "Kafka"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpReadAppAWS(r.client, ctx, req, resp, "Kafka")
	} else {
		common.ResourceOpReadAppDefault(r.client, ctx, req, resp, "Kafka")
	}

}


func (r *AppKafkaResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	appType := "Kafka"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpUpdateAppAWS(r.client, ctx, req, resp, "Kafka")
	} else {
		common.ResourceOpUpdateAppDefault(r.client, ctx, req, resp, "Kafka")
	}

}


func (r *AppKafkaResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	appType := "Kafka"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		common.ResourceOpDeleteAppAWS(r.client, ctx, req, resp, "Kafka")
	} else { 
		common.ResourceOpDeleteAppDefault(r.client, ctx, req, resp, "Kafka")
	}
	

}


func (r *AppKafkaResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	appType := "Kafka"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
  		common.ResourceOpImportAppAWS(r.client, ctx, req, resp, "Kafka")
	} else {
		common.ResourceOpImportAppDefault(r.client, ctx, req, resp, "Kafka")
	}

}

