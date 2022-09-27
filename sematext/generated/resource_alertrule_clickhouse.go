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

// resourceAlertRuleClickhouse is the resource class that handles sematext_alertrule_clickhouse
func resourceAlertRuleClickhouse() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaAlertRule("ClickHouse"),
		CreateContext: resourceOperationCreateAlertRuleClickhouse,
		ReadContext:   resourceOperationReadAlertRuleClickhouse,
		UpdateContext: resourceOperationUpdateAlertRuleClickhouse,
		DeleteContext: resourceOperationDeleteAlertRuleClickhouse,
		Importer:      resourceOperationImportAlertRuleClickhouse(),
	}
}

// resourceOperationCreateAlertRuleClickhouse creates the sematext_alertrule_clickhouse resource.
func resourceOperationCreateAlertRuleClickhouse(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationCreateAlertRule(ctx, d, meta, "ClickHouse")
}

// resourceOperationReadAlertRuleClickhouse reads the sematext_alertrule_clickhouse resource from Sematext Cloud.
func resourceOperationReadAlertRuleClickhouse(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationReadAlertRule(ctx, d, meta, "ClickHouse")
}

// resourceOperationUpdateAlertRuleClickhouse updates Sematext Cloud from the sematext_alertrule_clickhouse resource.
func resourceOperationUpdateAlertRuleClickhouse(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationUpdateAlertRule(ctx, d, meta, "ClickHouse")
}

// resourceOperationDeleteAlertRuleClickhouse marks a sematext_alertrule_clickhouse resource as retired.
func resourceOperationDeleteAlertRuleClickhouse(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationDeleteAlertRule(ctx, d, meta, "ClickHouse")
}

// resourceOperationImportAlertRuleClickhouse imports a sematext_alertrule_clickhouse resource into the state file.
func resourceOperationImportAlertRuleClickhouse() *schema.ResourceImporter {
	apptype := "ClickHouse"
	return sematext.ResourceOperationImportAlertRule()
}
