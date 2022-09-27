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

// resourceAlertRuleNginx is the resource class that handles sematext_alertrule_nginx
func resourceAlertRuleNginx() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaAlertRule("Nginx"),
		CreateContext: resourceOperationCreateAlertRuleNginx,
		ReadContext:   resourceOperationReadAlertRuleNginx,
		UpdateContext: resourceOperationUpdateAlertRuleNginx,
		DeleteContext: resourceOperationDeleteAlertRuleNginx,
		Importer:      resourceOperationImportAlertRuleNginx(),
	}
}

// resourceOperationCreateAlertRuleNginx creates the sematext_alertrule_nginx resource.
func resourceOperationCreateAlertRuleNginx(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationCreateAlertRule(ctx, d, meta, "Nginx")
}

// resourceOperationReadAlertRuleNginx reads the sematext_alertrule_nginx resource from Sematext Cloud.
func resourceOperationReadAlertRuleNginx(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationReadAlertRule(ctx, d, meta, "Nginx")
}

// resourceOperationUpdateAlertRuleNginx updates Sematext Cloud from the sematext_alertrule_nginx resource.
func resourceOperationUpdateAlertRuleNginx(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationUpdateAlertRule(ctx, d, meta, "Nginx")
}

// resourceOperationDeleteAlertRuleNginx marks a sematext_alertrule_nginx resource as retired.
func resourceOperationDeleteAlertRuleNginx(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationDeleteAlertRule(ctx, d, meta, "Nginx")
}

// resourceOperationImportAlertRuleNginx imports a sematext_alertrule_nginx resource into the state file.
func resourceOperationImportAlertRuleNginx() *schema.ResourceImporter {
	apptype := "Nginx"
	return sematext.ResourceOperationImportAlertRule()
}
