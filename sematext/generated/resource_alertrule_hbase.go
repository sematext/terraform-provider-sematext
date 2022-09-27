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

// resourceAlertRuleHbase is the resource class that handles sematext_alertrule_hbase
func resourceAlertRuleHbase() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaAlertRule("HBase"),
		CreateContext: resourceOperationCreateAlertRuleHbase,
		ReadContext:   resourceOperationReadAlertRuleHbase,
		UpdateContext: resourceOperationUpdateAlertRuleHbase,
		DeleteContext: resourceOperationDeleteAlertRuleHbase,
		Importer:      resourceOperationImportAlertRuleHbase(),
	}
}

// resourceOperationCreateAlertRuleHbase creates the sematext_alertrule_hbase resource.
func resourceOperationCreateAlertRuleHbase(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationCreateAlertRule(ctx, d, meta, "HBase")
}

// resourceOperationReadAlertRuleHbase reads the sematext_alertrule_hbase resource from Sematext Cloud.
func resourceOperationReadAlertRuleHbase(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationReadAlertRule(ctx, d, meta, "HBase")
}

// resourceOperationUpdateAlertRuleHbase updates Sematext Cloud from the sematext_alertrule_hbase resource.
func resourceOperationUpdateAlertRuleHbase(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationUpdateAlertRule(ctx, d, meta, "HBase")
}

// resourceOperationDeleteAlertRuleHbase marks a sematext_alertrule_hbase resource as retired.
func resourceOperationDeleteAlertRuleHbase(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationDeleteAlertRule(ctx, d, meta, "HBase")
}

// resourceOperationImportAlertRuleHbase imports a sematext_alertrule_hbase resource into the state file.
func resourceOperationImportAlertRuleHbase() *schema.ResourceImporter {
	apptype := "HBase"
	return sematext.ResourceOperationImportAlertRule()
}
