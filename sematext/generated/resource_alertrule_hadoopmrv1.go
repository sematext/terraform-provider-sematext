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

// resourceAlertRuleHadoopmrv1 is the resource class that handles sematext_alertrule_hadoopmrv1
func resourceAlertRuleHadoopmrv1() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaAlertRule("Hadoop-MRv1"),
		CreateContext: resourceOperationCreateAlertRuleHadoopmrv1,
		ReadContext:   resourceOperationReadAlertRuleHadoopmrv1,
		UpdateContext: resourceOperationUpdateAlertRuleHadoopmrv1,
		DeleteContext: resourceOperationDeleteAlertRuleHadoopmrv1,
		Importer:      resourceOperationImportAlertRuleHadoopmrv1(),
	}
}

// resourceOperationCreateAlertRuleHadoopmrv1 creates the sematext_alertrule_hadoopmrv1 resource.
func resourceOperationCreateAlertRuleHadoopmrv1(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationCreateAlertRule(ctx, d, meta, "Hadoop-MRv1")
}

// resourceOperationReadAlertRuleHadoopmrv1 reads the sematext_alertrule_hadoopmrv1 resource from Sematext Cloud.
func resourceOperationReadAlertRuleHadoopmrv1(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationReadAlertRule(ctx, d, meta, "Hadoop-MRv1")
}

// resourceOperationUpdateAlertRuleHadoopmrv1 updates Sematext Cloud from the sematext_alertrule_hadoopmrv1 resource.
func resourceOperationUpdateAlertRuleHadoopmrv1(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationUpdateAlertRule(ctx, d, meta, "Hadoop-MRv1")
}

// resourceOperationDeleteAlertRuleHadoopmrv1 marks a sematext_alertrule_hadoopmrv1 resource as retired.
func resourceOperationDeleteAlertRuleHadoopmrv1(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationDeleteAlertRule(ctx, d, meta, "Hadoop-MRv1")
}

// resourceOperationImportAlertRuleHadoopmrv1 imports a sematext_alertrule_hadoopmrv1 resource into the state file.
func resourceOperationImportAlertRuleHadoopmrv1() *schema.ResourceImporter {
	apptype := "Hadoop-MRv1"
	return sematext.ResourceOperationImportAlertRule()
}
