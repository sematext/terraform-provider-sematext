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

// resourceOperationAppRabbitmq is the resource class that handles sematext_app_rabbitmq
func resourceOperationAppRabbitmq() *schema.Resource {

	return &schema.Resource{
		CreateContext: resourceOperationCreateAppRabbitmq,
		ReadContext:   resourceOperationReadAppRabbitmq,
		UpdateContext: resourceOperationUpdateAppRabbitmq,
		DeleteContext: resourceOperationDeleteAppRabbitmq,
		Schema:        sematext.ResourceSchemaApp("rabbitmq"),
		Importer:      sematext.ResourceImporterApp("rabbitmq"),
	}
}

// resourceOperationCreateAppRabbitmq creates the sematext_app_rabbitmq resource.
func resourceOperationCreateAppRabbitmq(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "rabbitmq"
	err := sematext.ResourceOperationCreateApp(ctx, d, meta, apptype)
	return err
}

// resourceOperationReadAppRabbitmq reads the sematext_app_rabbitmq resource from Sematext Cloud.
func resourceOperationReadAppRabbitmq(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "rabbitmq"
	return sematext.ResourceOperationReadApp(ctx, d, meta, apptype)
}

// resourceOperationUpdateAppRabbitmq updates Sematext Cloud from the sematext_app_rabbitmq resource.
func resourceOperationUpdateAppRabbitmq(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "rabbitmq"
	return sematext.ResourceOperationUpdateApp(ctx, d, meta, apptype)
}

// resourceOperationDeleteAppRabbitmq marks a sematext_app_rabbitmq resource as retired.
func resourceOperationDeleteAppRabbitmq(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	apptype := "rabbitmq"
	return sematext.ResourceOperationDeleteApp(ctx, d, meta, apptype)
}
