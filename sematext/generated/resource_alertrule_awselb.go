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

// resourceAlertRuleAwselb is the resource class that handles sematext_alertrule_awselb
func resourceAlertRuleAwselb() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaAlertRule("AWS ELB"),
		CreateContext: resourceOperationCreateAlertRuleAwselb,
		ReadContext:   resourceOperationReadAlertRuleAwselb,
		UpdateContext: resourceOperationUpdateAlertRuleAwselb,
		DeleteContext: resourceOperationDeleteAlertRuleAwselb,
		Importer:      resourceOperationImportAlertRuleAwselb(),
	}
}

// resourceOperationCreateAlertRuleAwselb creates the sematext_alertrule_awselb resource.
func resourceOperationCreateAlertRuleAwselb(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationCreateAlertRule(ctx, d, meta, "AWS ELB")
}

// resourceOperationReadAlertRuleAwselb reads the sematext_alertrule_awselb resource from Sematext Cloud.
func resourceOperationReadAlertRuleAwselb(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationReadAlertRule(ctx, d, meta, "AWS ELB")
}

// resourceOperationUpdateAlertRuleAwselb updates Sematext Cloud from the sematext_alertrule_awselb resource.
func resourceOperationUpdateAlertRuleAwselb(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationUpdateAlertRule(ctx, d, meta, "AWS ELB")
}

// resourceOperationDeleteAlertRuleAwselb marks a sematext_alertrule_awselb resource as retired.
func resourceOperationDeleteAlertRuleAwselb(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationDeleteAlertRule(ctx, d, meta, "AWS ELB")
}

// resourceOperationImportAlertRuleAwselb imports a sematext_alertrule_awselb resource into the state file.
func resourceOperationImportAlertRuleAwselb() *schema.ResourceImporter {
	apptype := "AWS ELB"
	return sematext.ResourceOperationImportAlertRule()
}
