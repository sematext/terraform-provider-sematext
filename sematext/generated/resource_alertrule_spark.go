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

// resourceAlertRuleSpark is the resource class that handles sematext_alertrule_spark
func resourceAlertRuleSpark() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaAlertRule("Spark"),
		CreateContext: resourceOperationCreateAlertRuleSpark,
		ReadContext:   resourceOperationReadAlertRuleSpark,
		UpdateContext: resourceOperationUpdateAlertRuleSpark,
		DeleteContext: resourceOperationDeleteAlertRuleSpark,
		Importer:      resourceOperationImportAlertRuleSpark(),
	}
}

// resourceOperationCreateAlertRuleSpark creates the sematext_alertrule_spark resource.
func resourceOperationCreateAlertRuleSpark(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationCreateAlertRule(ctx, d, meta, "Spark")
}

// resourceOperationReadAlertRuleSpark reads the sematext_alertrule_spark resource from Sematext Cloud.
func resourceOperationReadAlertRuleSpark(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationReadAlertRule(ctx, d, meta, "Spark")
}

// resourceOperationUpdateAlertRuleSpark updates Sematext Cloud from the sematext_alertrule_spark resource.
func resourceOperationUpdateAlertRuleSpark(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationUpdateAlertRule(ctx, d, meta, "Spark")
}

// resourceOperationDeleteAlertRuleSpark marks a sematext_alertrule_spark resource as retired.
func resourceOperationDeleteAlertRuleSpark(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationDeleteAlertRule(ctx, d, meta, "Spark")
}

// resourceOperationImportAlertRuleSpark imports a sematext_alertrule_spark resource into the state file.
func resourceOperationImportAlertRuleSpark() *schema.ResourceImporter {
	apptype := "Spark"
	return sematext.ResourceOperationImportAlertRule()
}
