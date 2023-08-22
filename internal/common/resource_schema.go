package sematext

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/sematext/sematext-api-client-go/stcloud"
)

// ResourceSchemaApp contains common resource fields
func ResourceSchemaApp(appType string) schema.Schema {

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

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" { // @TODO How to make this conditional on appType?

		// TODO Pull aws credentials out of .config in OS independant manner, override with env.

		resourceSchema.Attributes["aws_region"] = schema.StringAttribute{
			Description:         "The region where AWS operations will take place. Examples\nare us-east-1, us-west-2, etc.",
			MarkdownDescription: "The region where AWS operations will take place. Examples\nare us-east-1, us-west-2, etc.",
			Required:            true,
			Sensitive:           true,
			Validators: []validator.String{
				stringvalidator.OneOf(stringStringMapKeys(stcloud.AWSRegion2STRegion)...),
			},
		}

		resourceSchema.Attributes["aws_access_key"] = schema.StringAttribute{
			Description:         "The access key for retrieval of stats from AWS Cloudwatch. You can retrieve this\nfrom the 'Security & Credentials' section of the AWS console.",
			MarkdownDescription: "The access key for retrieval of stats from AWS Cloudwatch. You can retrieve this\nfrom the 'Security & Credentials' section of the AWS console.",
			Required:            true,
			Sensitive:           true,
			Validators: []validator.String{
				stringvalidator.RegexMatches(regexp.MustCompile(`([0-9A-Za-z]{20})`), "invalid aws_access_key"),
			},
		}

		resourceSchema.Attributes["aws_secret_key"] = schema.StringAttribute{
			Description:         "The access key for retrieval of stats from AWS Cloudwatch. You can retrieve this\nfrom the 'Security & Credentials' section of the AWS console.",
			MarkdownDescription: "The access key for retrieval of stats from AWS Cloudwatch. You can retrieve this\nfrom the 'Security & Credentials' section of the AWS console.",
			Required:            true,
			Sensitive:           true,
			Validators: []validator.String{
				stringvalidator.RegexMatches(regexp.MustCompile(`([0-9A-Za-z+/=]{40})`), "invalid aws_secret_key"),
			},
		}

		resourceSchema.Attributes["aws_fetch_frequency"] = schema.StringAttribute{
			Description:         "How frequently to fetch metrics. One of MINUTE|FIVE_MINUTES|FIFTEEN_MINUTES",
			MarkdownDescription: "How frequently to fetch metrics. One of MINUTE|FIVE_MINUTES|FIFTEEN_MINUTES",
			Required:            true,
			Sensitive:           true,
			Validators: []validator.String{
				stringvalidator.OneOf("MINUTE", "FIVE_MINUTES", "FIVE_MINUTES", "FIFTEEN_MINUTES"),
			},
			Default: stringdefault.StaticString("MINUTE"),
		}

	}

	return resourceSchema
}

type AppTokenType struct {
	Name []string `tdfsk:"name"`
}

type AppTokenEntryType struct {
	CreatedAt string `tdfsk:"createdAt"`
	Enabled   bool   `tdfsk:"enabled"`
	ID        int64  `tdfsk:"id"`
	Name      string `tdfsk:"name"`
	Readable  string `tdfsk:"readable"`
	Token     string `tdfsk:"token"`
	Writeable bool   `tdfsk:"writable"`
}

type ResourceAppModel struct {
	Id                string `tfsdk:"id"`
	Name              string `tfsdk:"name"`
	BillingPlanId     int64  `tfsdk:"billing_plan_id"`
	DiscountCode      string `tfsdk:"discount_code"`
	AppToken          []AppTokenType
	ApptokenEntry     []AppTokenEntryType
	ScAppTokenEntries map[string]string `tfsdk:"sc_apptoken_entries"`

	// @TODO make these conditional on apptype?.
	AwsAccessKey      types.String `tfsdk:"aws_access_key"`
	AwsSecretKey      types.String `tfsdk:"aws_secret_key"`
	AwsFetchFrequency types.String `tfsdk:"aws_fetch_frequency"`
	AwsRegion         types.String `tfsdk:"aws_region"`
}

type ResourceApp struct {
	client *stcloud.APIClient
}
