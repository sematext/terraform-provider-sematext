package generated

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_App.go.template
	Then run generate/generate.sh
*/

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/sematext/terraform-provider-sematext/sematext"
)

// resourceAppElasticsearch is the resource class that handles sematext_app_elasticsearch
func resourceAppElasticsearch() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaApp("Elastic Search"),
		CreateContext: resourceOperationCreateAppElasticsearch,
		ReadContext:   resourceOperationReadAppElasticsearch,
		UpdateContext: resourceOperationUpdateAppElasticsearch,
		DeleteContext: resourceOperationDeleteAppElasticsearch,
		Importer:      resourceOperationImportAppElasticsearch(),
	}
}

// resourceOperationCreateAppElasticsearch creates the sematext_app_elasticsearch resource.
func resourceOperationCreateAppElasticsearch(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Elastic Search"
	err := sematext.ResourceOperationCreateApp(ctx, d, meta, apptype)
	return err
}

// resourceOperationReadAppElasticsearch reads the sematext_app_elasticsearch resource from Sematext Cloud.
func resourceOperationReadAppElasticsearch(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Elastic Search"
	return sematext.ResourceOperationReadApp(ctx, d, meta, apptype)
}

// resourceOperationUpdateAppElasticsearch updates Sematext Cloud from the sematext_app_elasticsearch resource.
func resourceOperationUpdateAppElasticsearch(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Elastic Search"
	return sematext.ResourceOperationUpdateApp(ctx, d, meta, apptype)
}

// resourceOperationDeleteAppElasticsearch marks a sematext_app_elasticsearch resource as retired.
func resourceOperationDeleteAppElasticsearch(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Elastic Search"
	return sematext.ResourceOperationDeleteApp(ctx, d, meta, apptype)
}

// resourceOperationImportAppElasticsearch imports a sematext_app_elasticsearch resource into the state file.
func resourceOperationImportAppElasticsearch() *schema.ResourceImporter {
	apptype := "Elastic Search"
	return sematext.ResourceOperationImportApp(apptype)
}
