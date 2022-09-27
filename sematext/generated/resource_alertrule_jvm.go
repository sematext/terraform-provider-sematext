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

// resourceAlertRuleJvm is the resource class that handles sematext_alertrule_jvm
func resourceAlertRuleJvm() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaAlertRule("JVM"),
		CreateContext: resourceOperationCreateAlertRuleJvm,
		ReadContext:   resourceOperationReadAlertRuleJvm,
		UpdateContext: resourceOperationUpdateAlertRuleJvm,
		DeleteContext: resourceOperationDeleteAlertRuleJvm,
		Importer:      resourceOperationImportAlertRuleJvm(),
	}
}

// resourceOperationCreateAlertRuleJvm creates the sematext_alertrule_jvm resource.
func resourceOperationCreateAlertRuleJvm(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationCreateAlertRule(ctx, d, meta, "JVM")
}

// resourceOperationReadAlertRuleJvm reads the sematext_alertrule_jvm resource from Sematext Cloud.
func resourceOperationReadAlertRuleJvm(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationReadAlertRule(ctx, d, meta, "JVM")
}

// resourceOperationUpdateAlertRuleJvm updates Sematext Cloud from the sematext_alertrule_jvm resource.
func resourceOperationUpdateAlertRuleJvm(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationUpdateAlertRule(ctx, d, meta, "JVM")
}

// resourceOperationDeleteAlertRuleJvm marks a sematext_alertrule_jvm resource as retired.
func resourceOperationDeleteAlertRuleJvm(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationDeleteAlertRule(ctx, d, meta, "JVM")
}

// resourceOperationImportAlertRuleJvm imports a sematext_alertrule_jvm resource into the state file.
func resourceOperationImportAlertRuleJvm() *schema.ResourceImporter {
	apptype := "JVM"
	return sematext.ResourceOperationImportAlertRule()
}
