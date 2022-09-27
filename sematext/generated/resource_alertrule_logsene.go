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

// resourceAlertRuleLogsene is the resource class that handles sematext_alertrule_logsene
func resourceAlertRuleLogsene() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaAlertRule("Logsene"),
		CreateContext: resourceOperationCreateAlertRuleLogsene,
		ReadContext:   resourceOperationReadAlertRuleLogsene,
		UpdateContext: resourceOperationUpdateAlertRuleLogsene,
		DeleteContext: resourceOperationDeleteAlertRuleLogsene,
		Importer:      resourceOperationImportAlertRuleLogsene(),
	}
}

// resourceOperationCreateAlertRuleLogsene creates the sematext_alertrule_logsene resource.
func resourceOperationCreateAlertRuleLogsene(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationCreateAlertRule(ctx, d, meta, "Logsene")
}

// resourceOperationReadAlertRuleLogsene reads the sematext_alertrule_logsene resource from Sematext Cloud.
func resourceOperationReadAlertRuleLogsene(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationReadAlertRule(ctx, d, meta, "Logsene")
}

// resourceOperationUpdateAlertRuleLogsene updates Sematext Cloud from the sematext_alertrule_logsene resource.
func resourceOperationUpdateAlertRuleLogsene(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationUpdateAlertRule(ctx, d, meta, "Logsene")
}

// resourceOperationDeleteAlertRuleLogsene marks a sematext_alertrule_logsene resource as retired.
func resourceOperationDeleteAlertRuleLogsene(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationDeleteAlertRule(ctx, d, meta, "Logsene")
}

// resourceOperationImportAlertRuleLogsene imports a sematext_alertrule_logsene resource into the state file.
func resourceOperationImportAlertRuleLogsene() *schema.ResourceImporter {
	apptype := "Logsene"
	return sematext.ResourceOperationImportAlertRule()
}
