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

// resourceAlertRuleNginxplus is the resource class that handles sematext_alertrule_nginxplus
func resourceAlertRuleNginxplus() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaAlertRule("Nginx-Plus"),
		CreateContext: resourceOperationCreateAlertRuleNginxplus,
		ReadContext:   resourceOperationReadAlertRuleNginxplus,
		UpdateContext: resourceOperationUpdateAlertRuleNginxplus,
		DeleteContext: resourceOperationDeleteAlertRuleNginxplus,
		Importer:      resourceOperationImportAlertRuleNginxplus(),
	}
}

// resourceOperationCreateAlertRuleNginxplus creates the sematext_alertrule_nginxplus resource.
func resourceOperationCreateAlertRuleNginxplus(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationCreateAlertRule(ctx, d, meta, "Nginx-Plus")
}

// resourceOperationReadAlertRuleNginxplus reads the sematext_alertrule_nginxplus resource from Sematext Cloud.
func resourceOperationReadAlertRuleNginxplus(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationReadAlertRule(ctx, d, meta, "Nginx-Plus")
}

// resourceOperationUpdateAlertRuleNginxplus updates Sematext Cloud from the sematext_alertrule_nginxplus resource.
func resourceOperationUpdateAlertRuleNginxplus(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationUpdateAlertRule(ctx, d, meta, "Nginx-Plus")
}

// resourceOperationDeleteAlertRuleNginxplus marks a sematext_alertrule_nginxplus resource as retired.
func resourceOperationDeleteAlertRuleNginxplus(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationDeleteAlertRule(ctx, d, meta, "Nginx-Plus")
}

// resourceOperationImportAlertRuleNginxplus imports a sematext_alertrule_nginxplus resource into the state file.
func resourceOperationImportAlertRuleNginxplus() *schema.ResourceImporter {
	apptype := "Nginx-Plus"
	return sematext.ResourceOperationImportAlertRule()
}
