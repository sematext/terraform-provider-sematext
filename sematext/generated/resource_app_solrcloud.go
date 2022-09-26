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

// resourceAppSolrcloud is the resource class that handles sematext_app_solrcloud
func resourceAppSolrcloud() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaApp("SolrCloud"),
		CreateContext: resourceOperationCreateAppSolrcloud,
		ReadContext:   resourceOperationReadAppSolrcloud,
		UpdateContext: resourceOperationUpdateAppSolrcloud,
		DeleteContext: resourceOperationDeleteAppSolrcloud,
		Importer:      resourceOperationImportAppSolrcloud(),
	}
}

// resourceOperationCreateAppSolrcloud creates the sematext_app_solrcloud resource.
func resourceOperationCreateAppSolrcloud(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "SolrCloud"
	err := sematext.ResourceOperationCreateApp(ctx, d, meta, apptype)
	return err
}

// resourceOperationReadAppSolrcloud reads the sematext_app_solrcloud resource from Sematext Cloud.
func resourceOperationReadAppSolrcloud(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "SolrCloud"
	return sematext.ResourceOperationReadApp(ctx, d, meta, apptype)
}

// resourceOperationUpdateAppSolrcloud updates Sematext Cloud from the sematext_app_solrcloud resource.
func resourceOperationUpdateAppSolrcloud(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "SolrCloud"
	return sematext.ResourceOperationUpdateApp(ctx, d, meta, apptype)
}

// resourceOperationDeleteAppSolrcloud marks a sematext_app_solrcloud resource as retired.
func resourceOperationDeleteAppSolrcloud(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "SolrCloud"
	return sematext.ResourceOperationDeleteApp(ctx, d, meta, apptype)
}

// resourceOperationImportAppSolrcloud imports a sematext_app_solrcloud resource into the state file.
func resourceOperationImportAppSolrcloud() *schema.ResourceImporter {
	apptype := "SolrCloud"
	return sematext.ResourceOperationImportApp(apptype)
}
