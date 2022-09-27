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

// resourceAlertRuleNodejs is the resource class that handles sematext_alertrule_nodejs
func resourceAlertRuleNodejs() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaAlertRule("Node.js"),
		CreateContext: resourceOperationCreateAlertRuleNodejs,
		ReadContext:   resourceOperationReadAlertRuleNodejs,
		UpdateContext: resourceOperationUpdateAlertRuleNodejs,
		DeleteContext: resourceOperationDeleteAlertRuleNodejs,
		Importer:      resourceOperationImportAlertRuleNodejs(),
	}
}

// resourceOperationCreateAlertRuleNodejs creates the sematext_alertrule_nodejs resource.
func resourceOperationCreateAlertRuleNodejs(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationCreateAlertRule(ctx, d, meta, "Node.js")
}

// resourceOperationReadAlertRuleNodejs reads the sematext_alertrule_nodejs resource from Sematext Cloud.
func resourceOperationReadAlertRuleNodejs(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationReadAlertRule(ctx, d, meta, "Node.js")
}

// resourceOperationUpdateAlertRuleNodejs updates Sematext Cloud from the sematext_alertrule_nodejs resource.
func resourceOperationUpdateAlertRuleNodejs(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationUpdateAlertRule(ctx, d, meta, "Node.js")
}

// resourceOperationDeleteAlertRuleNodejs marks a sematext_alertrule_nodejs resource as retired.
func resourceOperationDeleteAlertRuleNodejs(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationDeleteAlertRule(ctx, d, meta, "Node.js")
}

// resourceOperationImportAlertRuleNodejs imports a sematext_alertrule_nodejs resource into the state file.
func resourceOperationImportAlertRuleNodejs() *schema.ResourceImporter {
	apptype := "Node.js"
	return sematext.ResourceOperationImportAlertRule()
}
