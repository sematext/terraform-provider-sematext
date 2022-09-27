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

// resourceAlertRuleZookeeper is the resource class that handles sematext_alertrule_zookeeper
func resourceAlertRuleZookeeper() *schema.Resource {

	return &schema.Resource{
		Schema:        sematext.ResourceSchemaAlertRule("ZooKeeper"),
		CreateContext: resourceOperationCreateAlertRuleZookeeper,
		ReadContext:   resourceOperationReadAlertRuleZookeeper,
		UpdateContext: resourceOperationUpdateAlertRuleZookeeper,
		DeleteContext: resourceOperationDeleteAlertRuleZookeeper,
		Importer:      resourceOperationImportAlertRuleZookeeper(),
	}
}

// resourceOperationCreateAlertRuleZookeeper creates the sematext_alertrule_zookeeper resource.
func resourceOperationCreateAlertRuleZookeeper(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationCreateAlertRule(ctx, d, meta, "ZooKeeper")
}

// resourceOperationReadAlertRuleZookeeper reads the sematext_alertrule_zookeeper resource from Sematext Cloud.
func resourceOperationReadAlertRuleZookeeper(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationReadAlertRule(ctx, d, meta, "ZooKeeper")
}

// resourceOperationUpdateAlertRuleZookeeper updates Sematext Cloud from the sematext_alertrule_zookeeper resource.
func resourceOperationUpdateAlertRuleZookeeper(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationUpdateAlertRule(ctx, d, meta, "ZooKeeper")
}

// resourceOperationDeleteAlertRuleZookeeper marks a sematext_alertrule_zookeeper resource as retired.
func resourceOperationDeleteAlertRuleZookeeper(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return sematext.ResourceOperationDeleteAlertRule(ctx, d, meta, "ZooKeeper")
}

// resourceOperationImportAlertRuleZookeeper imports a sematext_alertrule_zookeeper resource into the state file.
func resourceOperationImportAlertRuleZookeeper() *schema.ResourceImporter {
	apptype := "ZooKeeper"
	return sematext.ResourceOperationImportAlertRule()
}
