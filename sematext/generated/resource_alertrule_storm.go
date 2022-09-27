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

// resourceAlertRuleStorm is the resource class that handles sematext_alertrule_storm
func resourceAlertRuleStorm() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaAlertRule("Storm"),
		CreateContext: resourceOperationCreateAlertRuleStorm,
		ReadContext:   resourceOperationReadAlertRuleStorm,
		UpdateContext: resourceOperationUpdateAlertRuleStorm,
		DeleteContext: resourceOperationDeleteAlertRuleStorm,
		Importer:      resourceOperationImportAlertRuleStorm(),
	}
}

// resourceOperationCreateAlertRuleStorm creates the sematext_alertrule_storm resource.
func resourceOperationCreateAlertRuleStorm(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationCreateAlertRule(ctx, d, meta, "Storm")
}

// resourceOperationReadAlertRuleStorm reads the sematext_alertrule_storm resource from Sematext Cloud.
func resourceOperationReadAlertRuleStorm(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationReadAlertRule(ctx, d, meta, "Storm")
}

// resourceOperationUpdateAlertRuleStorm updates Sematext Cloud from the sematext_alertrule_storm resource.
func resourceOperationUpdateAlertRuleStorm(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationUpdateAlertRule(ctx, d, meta, "Storm")
}

// resourceOperationDeleteAlertRuleStorm marks a sematext_alertrule_storm resource as retired.
func resourceOperationDeleteAlertRuleStorm(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationDeleteAlertRule(ctx, d, meta, "Storm")
}

// resourceOperationImportAlertRuleStorm imports a sematext_alertrule_storm resource into the state file.
func resourceOperationImportAlertRuleStorm() *schema.ResourceImporter {
	apptype := "Storm"
	return sematext.ResourceOperationImportAlertRule()
}
