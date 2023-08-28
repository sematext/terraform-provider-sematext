package common

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/sematext/sematext-api-client-go/stcloud"
)

func ResourceSchemaAppDefault(appType string) schema.Schema {

	resourceSchema := schema.Schema{
		// This description is used by the documentation generator and the language server.
		Description:         appType, // TODO complete string
		MarkdownDescription: appType, // TODO complete string

		Attributes: map[string]schema.Attribute{

			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Example identifier",
			},

			"name": schema.StringAttribute{
				Description:         "(Required) Label for the monitor app in Sematext Cloud.",
				MarkdownDescription: "(Required) Label for the monitor app in Sematext Cloud.",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},

			"billing_plan_id": schema.Int64Attribute{
				Description:         "(Required) Plan ID attached to the monitor app in Sematext Cloud.",
				MarkdownDescription: "(Required) Plan ID attached to the monitor app in Sematext Cloud.",
				Required:            true,
				Validators: []validator.Int64{
					int64validator.OneOf(intArraytoInt64array(stcloud.LookupAppType2PlanID[appType])...),
				},
			},

			// At this point the association is with all app-tokens in the SC app. Association of individual app-token with container instance is set by the provisioner(s).
			"apptoken": schema.SetNestedAttribute{
				Description:         "(Required) apptoken(s) attached to the monitor app in Sematext Cloud.",
				MarkdownDescription: "(Required) apptoken(s) attached to the monitor app in Sematext Cloud.",
				Required:            true, // TODO isint this optional?
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.ListAttribute{
							ElementType: types.StringType,
							Required:    true,
						},
					},
				},
			},

			// Once the app is created this is pulled back from SC and lives only in state file. This is checked each time to see ids line up with names.
			"apptoken_entries": schema.SetNestedAttribute{
				Description:         "One or more apptoken entries. Calculated, supplied by SC Cloud.",
				MarkdownDescription: "One or more apptoken entries. Calculated, supplied by SC Cloud.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{

					Attributes: map[string]schema.Attribute{

						"createdAt": schema.StringAttribute{
							Computed: true,
						},
						"enabled": schema.BoolAttribute{
							Computed: true,
						},
						"id": schema.Int64Attribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"readable": schema.StringAttribute{
							Computed: true,
						},
						"token": schema.StringAttribute{
							Computed: true,
						},
						"writeable": schema.BoolAttribute{
							Computed: true,
						},
					},
				},
			},

			// Once the app is created this is pulled back from SC and lives only in state file. This is checked each time to see ids line up with names.
			// Note : Simplified from above to make lookup more readible in scripts that use this. // TODO - is this still valid to use?
			"sc_apptoken_entries": schema.MapAttribute{
				Description:         "Map of apptoken name -> id. Calculated, supplied by SC Cloud.",
				MarkdownDescription: "Map of apptoken name -> id. Calculated, supplied by SC Cloud.",
				ElementType:         types.StringType,
				Computed:            true,
			},
		},
	}

	return resourceSchema
}
