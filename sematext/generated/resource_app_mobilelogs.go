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

// resourceAppMobilelogs is the resource class that handles sematext_app_mobilelogs
func resourceAppMobilelogs() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaApp("mobile-logs"),
		CreateContext: resourceOperationCreateAppMobilelogs,
		ReadContext:   resourceOperationReadAppMobilelogs,
		UpdateContext: resourceOperationUpdateAppMobilelogs,
		DeleteContext: resourceOperationDeleteAppMobilelogs,
		Importer:      resourceOperationImportAppMobilelogs(),
	}
}

// resourceOperationCreateAppMobilelogs creates the sematext_app_mobilelogs resource.
func resourceOperationCreateAppMobilelogs(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "mobile-logs"
	err := sematext.ResourceOperationCreateApp(ctx, d, meta, apptype)
	return err
}

// resourceOperationReadAppMobilelogs reads the sematext_app_mobilelogs resource from Sematext Cloud.
func resourceOperationReadAppMobilelogs(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "mobile-logs"
	return sematext.ResourceOperationReadApp(ctx, d, meta, apptype)
}

// resourceOperationUpdateAppMobilelogs updates Sematext Cloud from the sematext_app_mobilelogs resource.
func resourceOperationUpdateAppMobilelogs(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "mobile-logs"
	return sematext.ResourceOperationUpdateApp(ctx, d, meta, apptype)
}

// resourceOperationDeleteAppMobilelogs marks a sematext_app_mobilelogs resource as retired.
func resourceOperationDeleteAppMobilelogs(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "mobile-logs"
	return sematext.ResourceOperationDeleteApp(ctx, d, meta, apptype)
}

// resourceOperationImportAppMobilelogs imports a sematext_app_mobilelogs resource into the state file.
func resourceOperationImportAppMobilelogs() *schema.ResourceImporter {
	apptype := "mobile-logs"
	return sematext.ResourceOperationImportApp(apptype)
}
