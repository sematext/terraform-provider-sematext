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

// resourceOperationAppLogsene is the resource class that handles sematext_app_logsene
func resourceOperationAppLogsene() *schema.Resource {

	return &schema.Resource{
		CreateContext: resourceOperationCreateAppLogsene,
		ReadContext:   resourceOperationReadAppLogsene,
		UpdateContext: resourceOperationUpdateAppLogsene,
		DeleteContext: resourceOperationDeleteAppLogsene,
		Schema:        sematext.ResourceSchemaApp("Logsene"),
		Importer:      sematext.ResourceImporterApp("Logsene"),
	}
}

// resourceOperationCreateAppLogsene creates the sematext_app_logsene resource.
func resourceOperationCreateAppLogsene(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Logsene"
	err := sematext.ResourceOperationCreateApp(ctx, d, meta, apptype)
	return err
}

// resourceOperationReadAppLogsene reads the sematext_app_logsene resource from Sematext Cloud.
func resourceOperationReadAppLogsene(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Logsene"
	return sematext.ResourceOperationReadApp(ctx, d, meta, apptype)
}

// resourceOperationUpdateAppLogsene updates Sematext Cloud from the sematext_app_logsene resource.
func resourceOperationUpdateAppLogsene(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Logsene"
	return sematext.ResourceOperationUpdateApp(ctx, d, meta, apptype)
}

// resourceOperationDeleteAppLogsene marks a sematext_app_logsene resource as retired.
func resourceOperationDeleteAppLogsene(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Logsene"
	return sematext.ResourceOperationDeleteApp(ctx, d, meta, apptype)
}
