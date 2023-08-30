package jvm

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_App.go.template
	Then run generate/generate.sh
*/

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/sematext/sematext-api-client-go/stcloud"
	"github.com/sematext/terraform-provider-sematext/internal/operation"
	"github.com/sematext/terraform-provider-sematext/internal/common"
)


type Resource struct {
	client *stcloud.APIClient
}


// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &Resource{}
var _ resource.ResourceWithImportState = &Resource{}


func NewResource() resource.Resource {

	return &Resource{}

}


func (r *Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {

	resp.TypeName = "sematext_app_jvm"	

}


func (r *Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

	appType := "JVM"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		resp.Schema = common.ResourceSchemaAWS("JVM")
	} else{
		resp.Schema = common.ResourceSchemaDefault("JVM")
	}

}


func (r *Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	appType := "JVM"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		operation.ResourceOpConfigureAWS(r.client, ctx, req, resp, "Jvm")
	} else {
		operation.ResourceOpConfigureDefault(r.client, ctx, req, resp, "Jvm")
	}

}


func (r *Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	appType := "JVM"
		
	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		operation.ResourceOpCreateAWS(r.client, ctx, req, resp, "Jvm")
	} else{
		operation.ResourceOpCreateDefault(r.client, ctx, req, resp, "Jvm")
	}

}


func (r *Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	appType := "JVM"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		operation.ResourceOpReadAWS(r.client, ctx, req, resp, "Jvm")
	} else {
		operation.ResourceOpReadDefault(r.client, ctx, req, resp, "Jvm")
	}

}


func (r *Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	appType := "JVM"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		operation.ResourceOpUpdateAWS(r.client, ctx, req, resp, "Jvm")
	} else {
		operation.ResourceOpUpdateDefault(r.client, ctx, req, resp, "Jvm")
	}

}


func (r *Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	appType := "JVM"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
		operation.ResourceOpDeleteAWS(r.client, ctx, req, resp, "Jvm")
	} else { 
		operation.ResourceOpDeleteDefault(r.client, ctx, req, resp, "Jvm")
	}
	

}


func (r *Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	appType := "JVM"

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { 
  		operation.ResourceOpImportAWS(r.client, ctx, req, resp, "Jvm")
	} else {
		operation.ResourceOpImportDefault(r.client, ctx, req, resp, "Jvm")
	}

}

