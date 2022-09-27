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

// resourceAlertRuleSolr is the resource class that handles sematext_alertrule_solr
func resourceAlertRuleSolr() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaAlertRule("Solr"),
		CreateContext: resourceOperationCreateAlertRuleSolr,
		ReadContext:   resourceOperationReadAlertRuleSolr,
		UpdateContext: resourceOperationUpdateAlertRuleSolr,
		DeleteContext: resourceOperationDeleteAlertRuleSolr,
		Importer:      resourceOperationImportAlertRuleSolr(),
	}
}

// resourceOperationCreateAlertRuleSolr creates the sematext_alertrule_solr resource.
func resourceOperationCreateAlertRuleSolr(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationCreateAlertRule(ctx, d, meta, "Solr")
}

// resourceOperationReadAlertRuleSolr reads the sematext_alertrule_solr resource from Sematext Cloud.
func resourceOperationReadAlertRuleSolr(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationReadAlertRule(ctx, d, meta, "Solr")
}

// resourceOperationUpdateAlertRuleSolr updates Sematext Cloud from the sematext_alertrule_solr resource.
func resourceOperationUpdateAlertRuleSolr(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationUpdateAlertRule(ctx, d, meta, "Solr")
}

// resourceOperationDeleteAlertRuleSolr marks a sematext_alertrule_solr resource as retired.
func resourceOperationDeleteAlertRuleSolr(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationDeleteAlertRule(ctx, d, meta, "Solr")
}

// resourceOperationImportAlertRuleSolr imports a sematext_alertrule_solr resource into the state file.
func resourceOperationImportAlertRuleSolr() *schema.ResourceImporter {
	apptype := "Solr"
	return sematext.ResourceOperationImportAlertRule()
}
