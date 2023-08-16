package sematext

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/sematext/sematext-api-client-go/stcloud"
)

// ResourceSchemaApp contains common resource fields
func ResourceSchemaMonitoringApp(appType string) schema.Schema {

	resourceSchema := schema.Schema{
		// This description is used by the documentation generator and the language server.
		Description:         appType, // TODO complete string
		MarkdownDescription: appType, // TODO complete string

		Attributes: map[string]schema.Attribute{

			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Example identifier",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(), /// huh?
				},
			},

			"name": schema.StringAttribute{
				Description:         "(Required) Label for the monitor app in Sematext Cloud.",
				MarkdownDescription: "(Required) Label for the monitor app in Sematext Cloud.",
				Required:            true,
				ForceNew:            false, // how to do this in new framework?  https://pkg.go.dev/github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
				PlanModifiers: []planmodifier.String{ //https://developer.hashicorp.com/terraform/plugin/framework/resources/plan-modification
					stringplanmodifier.RequiresReplace(),
				}				
			},

			"billing_plan_id": schema.Int64Attribute{
				Description:         "(Required) Plan ID attached to the monitor app in Sematext Cloud.",
				MarkdownDescription: "(Required) Plan ID attached to the monitor app in Sematext Cloud.",
				Required:            true,
				ForceNew:            false, // how to do this in new framework?
				Validators: []validator.Int64{
					int64validator.OneOf(intArraytoInt64array(stcloud.LookupAppType2PlanID[appType])...),
				},
			},


			// At this point the association is with all app-tokens in the SC app. Association of individual app-token with container instance is set by the provisioner(s).
			"apptoken": schema.SetNestedAttribute{

				Description:         "(Required) apptoken(s) attached to the monitor app in Sematext Cloud.",
				MarkdownDescription: "(Required) apptoken(s) attached to the monitor app in Sematext Cloud.",
				Required:            true,  // TODO isint this optional?
				ForceNew:            false, // how to do this in new framework?
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.ListAttribute{
							ElementType: types.StringType,
							Required:    true,
							ForceNew:    false,
						},
					},
				},
			},

			// Once the app is created this is pulled back from SC and lives only in state file. This is checked each time to see ids line up with names.
			"apptoken_entries": schema.SetNestedAttribute{
				Description:         "One or more apptoken entries. Calculated, supplied by SC Cloud.",
				MarkdownDescription: "One or more apptoken entries. Calculated, supplied by SC Cloud.",
				Computed:            true,
				ForceNew:            false, // how to do this in new framework?
				NestedObject: schema.NestedAttributeObject{

					Attributes: map[string]schema.Attribute{

						"CreatedAt": schema.StringAttribute{
							Computed: true,
						},
						"Enabled": schema.BoolAttribute{
							Computed: true,
						},
						"ID": schema.Int64Attribute{
							Computed: true,
						},
						"Name": schema.StringAttribute{
							Computed: true,
						},
						"Readable": schema.StringAttribute{
							Computed: true,
						},
						"Token": schema.StringAttribute{
							Computed: true,
						},
						"Writeable": schema.BoolAttribute{
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
				ForceNew:            false, // how to do this in new framework?
			},
		},
	}

	if appType == "AWS EBS" || appType == "AWS EC2" || appType == "AWS ELB" {

		// TODO Pull aws credentials out of .config in OS dependant manner, override with env.

		resourceSchema.Attributes["aws_access_key"] = schema.StringAttribute{
			Description:         "The access key for retrieval of stats from AWS Cloudwatch. You can retrieve this\nfrom the 'Security & Credentials' section of the AWS console.",
			MarkdownDescription: "The access key for retrieval of stats from AWS Cloudwatch. You can retrieve this\nfrom the 'Security & Credentials' section of the AWS console.",
			Required:            true,
			ForceNew:            false, // how to do this in new framework?
			Validators: []validator.String{
				stringvalidator.RegexMatches(regexp.MustCompile(`([0-9A-Za-z]{20})`), "invalid aws_access_key"),
			},
		}

		resourceSchema.Attributes["aws_secret_key"] = schema.StringAttribute{
			Description:         "The access key for retrieval of stats from AWS Cloudwatch. You can retrieve this\nfrom the 'Security & Credentials' section of the AWS console.",
			MarkdownDescription: "The access key for retrieval of stats from AWS Cloudwatch. You can retrieve this\nfrom the 'Security & Credentials' section of the AWS console.",
			Required:            true,
			ForceNew:            false, // how to do this in new framework?
			Validators: []validator.String{
				stringvalidator.RegexMatches(regexp.MustCompile(`([0-9A-Za-z+/=]{40})`), "invalid aws_secret_key"),
			},
		}

		resourceSchema.Attributes["aws_fetch_frequency"] = schema.StringAttribute{
			Description:         "How frequently to fetch metrics. One of MINUTE|FIVE_MINUTES|FIFTEEN_MINUTES",
			MarkdownDescription: "How frequently to fetch metrics. One of MINUTE|FIVE_MINUTES|FIFTEEN_MINUTES",
			Required:            true,
			ForceNew:            false, // how to do this in new framework?
			Validators: []validator.String{
				stringvalidator.OneOf("MINUTE", "FIVE_MINUTES", "FIVE_MINUTES", "FIFTEEN_MINUTES"),
			},
		}

		resourceSchema.Attributes["aws_region"] = schema.StringAttribute{
			Description:         "The region where AWS operations will take place. Examples\nare us-east-1, us-west-2, etc.",
			MarkdownDescription: "The region where AWS operations will take place. Examples\nare us-east-1, us-west-2, etc.",
			Required:            true,
			ForceNew:            false, // how to do this in new framework?
			Validators: []validator.String{
				stringvalidator.OneOf(stringStringMapKeys(stcloud.AWSRegion2STRegion)...),
			},
			InputDefault: "us-east-1",                                                                   // how to do this in new framework?
			Default:      schema.MultiEnvDefaultFunc([]string{"AWS_REGION", "AWS_DEFAULT_REGION"}, nil), // TODO how to do this in new framework?
		}

	}

	return resourceSchema
}

type ResourceModel struct {
	Id            types.String `tfsdk:"id"`
	Name          types.String `tfsdk:"name"`
	BillingPlanId types.Int64  `tfsdk:"billing_plan_id"`
	DiscountCode  types.String `tfsdk:"discount_code"`
	AppToken      types.String `tfsdk:"apptoken"`

	// @TODO make these conditional on apptype.
	AwsAccessKey      types.String `tfsdk:"aws_access_key"`
	AwsSecretKey      types.String `tfsdk:"aws_secret_key"`
	AwsFetchFrequency types.String `tfsdk:"aws_fetch_frequency"`
	AwsRegion         types.String `tfsdk:"aws_region"`
}
