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

// resourceAlertRuleRedis is the resource class that handles sematext_alertrule_redis
func resourceAlertRuleRedis() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaAlertRule("Redis"),
		CreateContext: resourceOperationCreateAlertRuleRedis,
		ReadContext:   resourceOperationReadAlertRuleRedis,
		UpdateContext: resourceOperationUpdateAlertRuleRedis,
		DeleteContext: resourceOperationDeleteAlertRuleRedis,
		Importer:      resourceOperationImportAlertRuleRedis(),
	}
}

// resourceOperationCreateAlertRuleRedis creates the sematext_alertrule_redis resource.
func resourceOperationCreateAlertRuleRedis(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationCreateAlertRule(ctx, d, meta, "Redis")
}

// resourceOperationReadAlertRuleRedis reads the sematext_alertrule_redis resource from Sematext Cloud.
func resourceOperationReadAlertRuleRedis(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationReadAlertRule(ctx, d, meta, "Redis")
}

// resourceOperationUpdateAlertRuleRedis updates Sematext Cloud from the sematext_alertrule_redis resource.
func resourceOperationUpdateAlertRuleRedis(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationUpdateAlertRule(ctx, d, meta, "Redis")
}

// resourceOperationDeleteAlertRuleRedis marks a sematext_alertrule_redis resource as retired.
func resourceOperationDeleteAlertRuleRedis(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationDeleteAlertRule(ctx, d, meta, "Redis")
}

// resourceOperationImportAlertRuleRedis imports a sematext_alertrule_redis resource into the state file.
func resourceOperationImportAlertRuleRedis() *schema.ResourceImporter {
	apptype := "Redis"
	return sematext.ResourceOperationImportAlertRule()
}
