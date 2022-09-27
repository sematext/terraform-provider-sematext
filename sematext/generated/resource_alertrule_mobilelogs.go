package generated

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_app.go.template
	Then run generate/generate.sh
*/

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/sematext/terraform-provider-sematext/sematext"
)

// resourceAlertRuleMobilelogs is the resource class that handles sematext_alertrule_mobilelogs
func resourceAlertRuleMobilelogs() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaAlertRule("mobile-logs"),
		CreateContext: resourceOperationCreateAlertRuleMobilelogs,
		ReadContext:   resourceOperationReadAlertRuleMobilelogs,
		UpdateContext: resourceOperationUpdateAlertRuleMobilelogs,
		DeleteContext: resourceOperationDeleteAlertRuleMobilelogs,
		Importer:      resourceOperationImportAlertRuleMobilelogs(),
	}
}

// resourceOperationCreateAlertRuleMobilelogs creates the sematext_alertrule_mobilelogs resource.
func resourceOperationCreateAlertRuleMobilelogs(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationCreateAlertRule(ctx, d, meta, "mobile-logs")
}

// resourceOperationReadAlertRuleMobilelogs reads the sematext_alertrule_mobilelogs resource from Sematext Cloud.
func resourceOperationReadAlertRuleMobilelogs(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationReadAlertRule(ctx, d, meta, "mobile-logs")
}

// resourceOperationUpdateAlertRuleMobilelogs updates Sematext Cloud from the sematext_alertrule_mobilelogs resource.
func resourceOperationUpdateAlertRuleMobilelogs(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationUpdateAlertRule(ctx, d, meta, "mobile-logs")
}

// resourceOperationDeleteAlertRuleMobilelogs marks a sematext_alertrule_mobilelogs resource as retired.
func resourceOperationDeleteAlertRuleMobilelogs(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationDeleteAlertRule(ctx, d, meta, "mobile-logs")
}

// resourceOperationImportAlertRuleMobilelogs imports a sematext_alertrule_mobilelogs resource into the state file.
func resourceOperationImportAlertRuleMobilelogs() *schema.ResourceImporter {
	apptype := "mobile-logs"
	return sematext.ResourceOperationImportAlertRule()
}
