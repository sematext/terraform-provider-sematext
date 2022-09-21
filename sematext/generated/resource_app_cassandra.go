package generated

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/resource_sematext_App.go.template
	Then run generate/generate.sh
*/

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/sematext/terraform-provider-sematext/sematext"
)

// resourceOperationAppCassandra is the resource class that handles sematext_app_cassandra
func resourceOperationAppCassandra() *schema.Resource {

	return &schema.Resource{
		CreateContext: resourceOperationCreateAppCassandra,
		ReadContext:   resourceOperationReadAppCassandra,
		UpdateContext: resourceOperationUpdateAppCassandra,
		DeleteContext: resourceOperationDeleteAppCassandra,
		Schema:        sematext.ResourceSchemaApp("Cassandra"),
		Importer:      sematext.ResourceImporterApp("Cassandra"),
	}
}

// resourceOperationCreateAppCassandra creates the sematext_app_cassandra resource.
func resourceOperationCreateAppCassandra(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Cassandra"
	err := sematext.ResourceOperationCreateApp(ctx, d, meta, apptype)
	return err
}

// resourceOperationReadAppCassandra reads the sematext_app_cassandra resource from Sematext Cloud.
func resourceOperationReadAppCassandra(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Cassandra"
	return sematext.ResourceOperationReadApp(ctx, d, meta, apptype)
}

// resourceOperationUpdateAppCassandra updates Sematext Cloud from the sematext_app_cassandra resource.
func resourceOperationUpdateAppCassandra(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Cassandra"
	return sematext.ResourceOperationUpdateApp(ctx, d, meta, apptype)
}

// resourceOperationDeleteAppCassandra marks a sematext_app_cassandra resource as retired.
func resourceOperationDeleteAppCassandra(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "Cassandra"
	return sematext.ResourceOperationDeleteApp(ctx, d, meta, apptype)
}
