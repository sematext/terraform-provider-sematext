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

// resourceAlertRuleElasticsearch is the resource class that handles sematext_alertrule_elasticsearch
func resourceAlertRuleElasticsearch() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaAlertRule("Elastic Search"),
		CreateContext: resourceOperationCreateAlertRuleElasticsearch,
		ReadContext:   resourceOperationReadAlertRuleElasticsearch,
		UpdateContext: resourceOperationUpdateAlertRuleElasticsearch,
		DeleteContext: resourceOperationDeleteAlertRuleElasticsearch,
		Importer:      resourceOperationImportAlertRuleElasticsearch(),
	}
}

// resourceOperationCreateAlertRuleElasticsearch creates the sematext_alertrule_elasticsearch resource.
func resourceOperationCreateAlertRuleElasticsearch(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationCreateAlertRule(ctx, d, meta, "Elastic Search")
}

// resourceOperationReadAlertRuleElasticsearch reads the sematext_alertrule_elasticsearch resource from Sematext Cloud.
func resourceOperationReadAlertRuleElasticsearch(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationReadAlertRule(ctx, d, meta, "Elastic Search")
}

// resourceOperationUpdateAlertRuleElasticsearch updates Sematext Cloud from the sematext_alertrule_elasticsearch resource.
func resourceOperationUpdateAlertRuleElasticsearch(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationUpdateAlertRule(ctx, d, meta, "Elastic Search")
}

// resourceOperationDeleteAlertRuleElasticsearch marks a sematext_alertrule_elasticsearch resource as retired.
func resourceOperationDeleteAlertRuleElasticsearch(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationDeleteAlertRule(ctx, d, meta, "Elastic Search")
}

// resourceOperationImportAlertRuleElasticsearch imports a sematext_alertrule_elasticsearch resource into the state file.
func resourceOperationImportAlertRuleElasticsearch() *schema.ResourceImporter {
	apptype := "Elastic Search"
	return sematext.ResourceOperationImportAlertRule()
}
