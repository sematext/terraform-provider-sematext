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

// resourceAppApache is the resource class that handles sematext_app_apache
func resourceAppApache() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaApp("Apache"),
		CreateContext: resourceOperationCreateAppApache,
		ReadContext:   resourceOperationReadAppApache,
		UpdateContext: resourceOperationUpdateAppApache,
		DeleteContext: resourceOperationDeleteAppApache,
		Importer:      resourceOperationImportAppApache(),
	}
}

// resourceOperationCreateAppApache creates the sematext_app_apache resource.
func resourceOperationCreateAppApache(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Apache"
	err := sematext.ResourceOperationCreateApp(ctx, d, meta, apptype)
	return err
}

// resourceOperationReadAppApache reads the sematext_app_apache resource from Sematext Cloud.
func resourceOperationReadAppApache(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Apache"
	return sematext.ResourceOperationReadApp(ctx, d, meta, apptype)
}

// resourceOperationUpdateAppApache updates Sematext Cloud from the sematext_app_apache resource.
func resourceOperationUpdateAppApache(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Apache"
	return sematext.ResourceOperationUpdateApp(ctx, d, meta, apptype)
}

// resourceOperationDeleteAppApache marks a sematext_app_apache resource as retired.
func resourceOperationDeleteAppApache(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Apache"
	return sematext.ResourceOperationDeleteApp(ctx, d, meta, apptype)
}

// resourceOperationImportAppApache imports a sematext_app_apache resource into the state file.
func resourceOperationImportAppApache() *schema.ResourceImporter {
	apptype := "Apache"
	return sematext.ResourceOperationImportApp(apptype)
}
