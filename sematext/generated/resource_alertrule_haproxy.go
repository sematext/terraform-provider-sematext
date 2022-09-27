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

// resourceAlertRuleHaproxy is the resource class that handles sematext_alertrule_haproxy
func resourceAlertRuleHaproxy() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaAlertRule("HAProxy"),
		CreateContext: resourceOperationCreateAlertRuleHaproxy,
		ReadContext:   resourceOperationReadAlertRuleHaproxy,
		UpdateContext: resourceOperationUpdateAlertRuleHaproxy,
		DeleteContext: resourceOperationDeleteAlertRuleHaproxy,
		Importer:      resourceOperationImportAlertRuleHaproxy(),
	}
}

// resourceOperationCreateAlertRuleHaproxy creates the sematext_alertrule_haproxy resource.
func resourceOperationCreateAlertRuleHaproxy(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationCreateAlertRule(ctx, d, meta, "HAProxy")
}

// resourceOperationReadAlertRuleHaproxy reads the sematext_alertrule_haproxy resource from Sematext Cloud.
func resourceOperationReadAlertRuleHaproxy(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationReadAlertRule(ctx, d, meta, "HAProxy")
}

// resourceOperationUpdateAlertRuleHaproxy updates Sematext Cloud from the sematext_alertrule_haproxy resource.
func resourceOperationUpdateAlertRuleHaproxy(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationUpdateAlertRule(ctx, d, meta, "HAProxy")
}

// resourceOperationDeleteAlertRuleHaproxy marks a sematext_alertrule_haproxy resource as retired.
func resourceOperationDeleteAlertRuleHaproxy(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationDeleteAlertRule(ctx, d, meta, "HAProxy")
}

// resourceOperationImportAlertRuleHaproxy imports a sematext_alertrule_haproxy resource into the state file.
func resourceOperationImportAlertRuleHaproxy() *schema.ResourceImporter {
	apptype := "HAProxy"
	return sematext.ResourceOperationImportAlertRule()
}
