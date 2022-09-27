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

// resourceAlertRuleSolrcloud is the resource class that handles sematext_alertrule_solrcloud
func resourceAlertRuleSolrcloud() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaAlertRule("SolrCloud"),
		CreateContext: resourceOperationCreateAlertRuleSolrcloud,
		ReadContext:   resourceOperationReadAlertRuleSolrcloud,
		UpdateContext: resourceOperationUpdateAlertRuleSolrcloud,
		DeleteContext: resourceOperationDeleteAlertRuleSolrcloud,
		Importer:      resourceOperationImportAlertRuleSolrcloud(),
	}
}

// resourceOperationCreateAlertRuleSolrcloud creates the sematext_alertrule_solrcloud resource.
func resourceOperationCreateAlertRuleSolrcloud(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationCreateAlertRule(ctx, d, meta, "SolrCloud")
}

// resourceOperationReadAlertRuleSolrcloud reads the sematext_alertrule_solrcloud resource from Sematext Cloud.
func resourceOperationReadAlertRuleSolrcloud(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationReadAlertRule(ctx, d, meta, "SolrCloud")
}

// resourceOperationUpdateAlertRuleSolrcloud updates Sematext Cloud from the sematext_alertrule_solrcloud resource.
func resourceOperationUpdateAlertRuleSolrcloud(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationUpdateAlertRule(ctx, d, meta, "SolrCloud")
}

// resourceOperationDeleteAlertRuleSolrcloud marks a sematext_alertrule_solrcloud resource as retired.
func resourceOperationDeleteAlertRuleSolrcloud(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationDeleteAlertRule(ctx, d, meta, "SolrCloud")
}

// resourceOperationImportAlertRuleSolrcloud imports a sematext_alertrule_solrcloud resource into the state file.
func resourceOperationImportAlertRuleSolrcloud() *schema.ResourceImporter {
	apptype := "SolrCloud"
	return sematext.ResourceOperationImportAlertRule()
}
