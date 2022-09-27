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

// resourceAlertRuleHadoopyarn is the resource class that handles sematext_alertrule_hadoopyarn
func resourceAlertRuleHadoopyarn() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaAlertRule("Hadoop-YARN"),
		CreateContext: resourceOperationCreateAlertRuleHadoopyarn,
		ReadContext:   resourceOperationReadAlertRuleHadoopyarn,
		UpdateContext: resourceOperationUpdateAlertRuleHadoopyarn,
		DeleteContext: resourceOperationDeleteAlertRuleHadoopyarn,
		Importer:      resourceOperationImportAlertRuleHadoopyarn(),
	}
}

// resourceOperationCreateAlertRuleHadoopyarn creates the sematext_alertrule_hadoopyarn resource.
func resourceOperationCreateAlertRuleHadoopyarn(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationCreateAlertRule(ctx, d, meta, "Hadoop-YARN")
}

// resourceOperationReadAlertRuleHadoopyarn reads the sematext_alertrule_hadoopyarn resource from Sematext Cloud.
func resourceOperationReadAlertRuleHadoopyarn(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationReadAlertRule(ctx, d, meta, "Hadoop-YARN")
}

// resourceOperationUpdateAlertRuleHadoopyarn updates Sematext Cloud from the sematext_alertrule_hadoopyarn resource.
func resourceOperationUpdateAlertRuleHadoopyarn(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationUpdateAlertRule(ctx, d, meta, "Hadoop-YARN")
}

// resourceOperationDeleteAlertRuleHadoopyarn marks a sematext_alertrule_hadoopyarn resource as retired.
func resourceOperationDeleteAlertRuleHadoopyarn(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationDeleteAlertRule(ctx, d, meta, "Hadoop-YARN")
}

// resourceOperationImportAlertRuleHadoopyarn imports a sematext_alertrule_hadoopyarn resource into the state file.
func resourceOperationImportAlertRuleHadoopyarn() *schema.ResourceImporter {
	apptype := "Hadoop-YARN"
	return sematext.ResourceOperationImportAlertRule()
}
