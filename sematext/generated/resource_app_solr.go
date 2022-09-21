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

// resourceOperationAppSolr is the resource class that handles sematext_app_solr
func resourceOperationAppSolr() *schema.Resource {

	return &schema.Resource{
		CreateContext: resourceOperationCreateAppSolr,
		ReadContext:   resourceOperationReadAppSolr,
		UpdateContext: resourceOperationUpdateAppSolr,
		DeleteContext: resourceOperationDeleteAppSolr,
		Schema:        sematext.ResourceSchemaApp("Solr"),
		Importer:      sematext.ResourceImporterApp("Solr"),
	}
}

// resourceOperationCreateAppSolr creates the sematext_app_solr resource.
func resourceOperationCreateAppSolr(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Solr"
	err := sematext.ResourceOperationCreateApp(ctx, d, meta, apptype)
	return err
}

// resourceOperationReadAppSolr reads the sematext_app_solr resource from Sematext Cloud.
func resourceOperationReadAppSolr(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Solr"
	return sematext.ResourceOperationReadApp(ctx, d, meta, apptype)
}

// resourceOperationUpdateAppSolr updates Sematext Cloud from the sematext_app_solr resource.
func resourceOperationUpdateAppSolr(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Solr"
	return sematext.ResourceOperationUpdateApp(ctx, d, meta, apptype)
}

// resourceOperationDeleteAppSolr marks a sematext_app_solr resource as retired.
func resourceOperationDeleteAppSolr(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Solr"
	return sematext.ResourceOperationDeleteApp(ctx, d, meta, apptype)
}
