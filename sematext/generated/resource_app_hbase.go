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

// resourceOperationAppHbase is the resource class that handles sematext_app_hbase
func resourceOperationAppHbase() *schema.Resource {

	return &schema.Resource{
		CreateContext: resourceOperationCreateAppHbase,
		ReadContext:   resourceOperationReadAppHbase,
		UpdateContext: resourceOperationUpdateAppHbase,
		DeleteContext: resourceOperationDeleteAppHbase,
		Schema:        sematext.ResourceSchemaApp("HBase"),
		Importer:      sematext.ResourceImporterApp("HBase"),
	}
}

// resourceOperationCreateAppHbase creates the sematext_app_hbase resource.
func resourceOperationCreateAppHbase(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "HBase"
	err := sematext.ResourceOperationCreateApp(ctx, d, meta, apptype)
	return err
}

// resourceOperationReadAppHbase reads the sematext_app_hbase resource from Sematext Cloud.
func resourceOperationReadAppHbase(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "HBase"
	return sematext.ResourceOperationReadApp(ctx, d, meta, apptype)
}

// resourceOperationUpdateAppHbase updates Sematext Cloud from the sematext_app_hbase resource.
func resourceOperationUpdateAppHbase(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "HBase"
	return sematext.ResourceOperationUpdateApp(ctx, d, meta, apptype)
}

// resourceOperationDeleteAppHbase marks a sematext_app_hbase resource as retired.
func resourceOperationDeleteAppHbase(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "HBase"
	return sematext.ResourceOperationDeleteApp(ctx, d, meta, apptype)
}
