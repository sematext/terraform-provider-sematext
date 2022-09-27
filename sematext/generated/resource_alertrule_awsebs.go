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

// resourceAlertRuleAwsebs is the resource class that handles sematext_alertrule_awsebs
func resourceAlertRuleAwsebs() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaAlertRule("AWS EBS"),
		CreateContext: resourceOperationCreateAlertRuleAwsebs,
		ReadContext:   resourceOperationReadAlertRuleAwsebs,
		UpdateContext: resourceOperationUpdateAlertRuleAwsebs,
		DeleteContext: resourceOperationDeleteAlertRuleAwsebs,
		Importer:      resourceOperationImportAlertRuleAwsebs(),
	}
}

// resourceOperationCreateAlertRuleAwsebs creates the sematext_alertrule_awsebs resource.
func resourceOperationCreateAlertRuleAwsebs(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationCreateAlertRule(ctx, d, meta, "AWS EBS")
}

// resourceOperationReadAlertRuleAwsebs reads the sematext_alertrule_awsebs resource from Sematext Cloud.
func resourceOperationReadAlertRuleAwsebs(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationReadAlertRule(ctx, d, meta, "AWS EBS")
}

// resourceOperationUpdateAlertRuleAwsebs updates Sematext Cloud from the sematext_alertrule_awsebs resource.
func resourceOperationUpdateAlertRuleAwsebs(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationUpdateAlertRule(ctx, d, meta, "AWS EBS")
}

// resourceOperationDeleteAlertRuleAwsebs marks a sematext_alertrule_awsebs resource as retired.
func resourceOperationDeleteAlertRuleAwsebs(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationDeleteAlertRule(ctx, d, meta, "AWS EBS")
}

// resourceOperationImportAlertRuleAwsebs imports a sematext_alertrule_awsebs resource into the state file.
func resourceOperationImportAlertRuleAwsebs() *schema.ResourceImporter {
	apptype := "AWS EBS"
	return sematext.ResourceOperationImportAlertRule()
}
