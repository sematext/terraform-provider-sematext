package generated

/*
	Note: Generated file, any edits will be overwritten!
	Correct way to alter is to edit generate/provider.go.template
	Then run generate/generate.sh
*/

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"os"

	"github.com/blang/semver/v4"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/meta"
	"github.com/sematext/sematext-api-client-go/stcloud"
	"github.com/sematext/terraform-provider-sematext/sematext"
)

// Provider is the connection to Sematext Cloud.
func Provider() *schema.Provider {

	provider := &schema.Provider{

		Schema: map[string]*schema.Schema{
			"sematext_region": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SEMATEXT_REGION", "US"),
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					region := val.(string)
					if val == nil || !sematext.IsValidSematextRegion(region) {
						errs = append(errs, fmt.Errorf("ERROR  : sematext_region missing or invalid in sematext provider"))
					}
					return
				},
				Description: "The Sematext region, either US or EU.",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"sematext_app_akka":                resourceAppAkka(),
			"sematext_alertrule_akka":          resourceAlertRuleAkka(),
			"sematext_app_apache":              resourceAppApache(),
			"sematext_alertrule_apache":        resourceAlertRuleApache(),
			"sematext_app_awsebs":              resourceAppAwsebs(),
			"sematext_alertrule_awsebs":        resourceAlertRuleAwsebs(),
			"sematext_app_awsec2":              resourceAppAwsec2(),
			"sematext_alertrule_awsec2":        resourceAlertRuleAwsec2(),
			"sematext_app_awselb":              resourceAppAwselb(),
			"sematext_alertrule_awselb":        resourceAlertRuleAwselb(),
			"sematext_app_cassandra":           resourceAppCassandra(),
			"sematext_alertrule_cassandra":     resourceAlertRuleCassandra(),
			"sematext_app_clickhouse":          resourceAppClickhouse(),
			"sematext_alertrule_clickhouse":    resourceAlertRuleClickhouse(),
			"sematext_app_elasticsearch":       resourceAppElasticsearch(),
			"sematext_alertrule_elasticsearch": resourceAlertRuleElasticsearch(),
			"sematext_app_hadoopmrv1":          resourceAppHadoopmrv1(),
			"sematext_alertrule_hadoopmrv1":    resourceAlertRuleHadoopmrv1(),
			"sematext_app_hadoopyarn":          resourceAppHadoopyarn(),
			"sematext_alertrule_hadoopyarn":    resourceAlertRuleHadoopyarn(),
			"sematext_app_haproxy":             resourceAppHaproxy(),
			"sematext_alertrule_haproxy":       resourceAlertRuleHaproxy(),
			"sematext_app_hbase":               resourceAppHbase(),
			"sematext_alertrule_hbase":         resourceAlertRuleHbase(),
			"sematext_app_infra":               resourceAppInfra(),
			"sematext_alertrule_infra":         resourceAlertRuleInfra(),
			"sematext_app_jvm":                 resourceAppJvm(),
			"sematext_alertrule_jvm":           resourceAlertRuleJvm(),
			"sematext_app_kafka":               resourceAppKafka(),
			"sematext_alertrule_kafka":         resourceAlertRuleKafka(),
			"sematext_app_logsene":             resourceAppLogsene(),
			"sematext_alertrule_logsene":       resourceAlertRuleLogsene(),
			"sematext_app_mongodb":             resourceAppMongodb(),
			"sematext_alertrule_mongodb":       resourceAlertRuleMongodb(),
			"sematext_app_mysql":               resourceAppMysql(),
			"sematext_alertrule_mysql":         resourceAlertRuleMysql(),
			"sematext_app_nginx":               resourceAppNginx(),
			"sematext_alertrule_nginx":         resourceAlertRuleNginx(),
			"sematext_app_nginxplus":           resourceAppNginxplus(),
			"sematext_alertrule_nginxplus":     resourceAlertRuleNginxplus(),
			"sematext_app_nodejs":              resourceAppNodejs(),
			"sematext_alertrule_nodejs":        resourceAlertRuleNodejs(),
			"sematext_app_redis":               resourceAppRedis(),
			"sematext_alertrule_redis":         resourceAlertRuleRedis(),
			"sematext_app_solr":                resourceAppSolr(),
			"sematext_alertrule_solr":          resourceAlertRuleSolr(),
			"sematext_app_solrcloud":           resourceAppSolrcloud(),
			"sematext_alertrule_solrcloud":     resourceAlertRuleSolrcloud(),
			"sematext_app_spark":               resourceAppSpark(),
			"sematext_alertrule_spark":         resourceAlertRuleSpark(),
			"sematext_app_storm":               resourceAppStorm(),
			"sematext_alertrule_storm":         resourceAlertRuleStorm(),
			"sematext_app_tomcat":              resourceAppTomcat(),
			"sematext_alertrule_tomcat":        resourceAlertRuleTomcat(),
			"sematext_app_zookeeper":           resourceAppZookeeper(),
			"sematext_alertrule_zookeeper":     resourceAlertRuleZookeeper(),
			"sematext_app_postgresql":          resourceAppPostgresql(),
			"sematext_alertrule_postgresql":    resourceAlertRulePostgresql(),
			"sematext_app_rabbitmq":            resourceAppRabbitmq(),
			"sematext_alertrule_rabbitmq":      resourceAlertRuleRabbitmq(),
			"sematext_app_mobilelogs":          resourceAppMobilelogs(),
			"sematext_alertrule_mobilelogs":    resourceAlertRuleMobilelogs(),
		},
	}

	provider.ConfigureContextFunc = func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {

		//func(d *schema.ResourceData) (interface{}, error) {
		var diags diag.Diagnostics
		var err error
		var tfVersionString string
		var tfVersionSemver semver.Version
		var tfVersionSupportedRange semver.Range
		var sdkVersionString string
		var sdkVersionSemver semver.Version
		var sdkVersionSupportedRange semver.Range
		var baseURL *url.URL

		tfVersionString = provider.TerraformVersion
		if tfVersionString == "" {

			sdkVersionString = meta.SDKVersion

			if sdkVersionString == "" {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "failed to parse Terraform version",
					Detail:   "failed to parse Terraform version",
				})

				return nil, diags
			}

			sdkVersionSemver, err = semver.Parse(sdkVersionString)
			if err != nil {
				return nil, diag.FromErr(err)
			}

			sdkVersionSupportedRange, err = semver.ParseRange(">=2.0.3")
			if err != nil {
				return nil, diag.FromErr(err)
			}

			if !sdkVersionSupportedRange(sdkVersionSemver) {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Terraform sdk version not supported",
					Detail:   "Terraform sdk version not supported",
				})

				return nil, diags
			}

			tfVersionString = "0.13.0"

		}

		tfVersionSemver, err = semver.Parse(tfVersionString)
		if err != nil {
			return nil, diag.FromErr(err)
		}

		tfVersionSupportedRange, err = semver.ParseRange(">=0.13.0")
		if err != nil {
			return nil, diag.FromErr(err)
		}

		if !tfVersionSupportedRange(tfVersionSemver) {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Terraform version must be >=0.13.0",
				Detail:   "Terraform version must be >=0.13.0",
			})

			return nil, diags

		}

		region := d.Get("sematext_region").(string)
		if !sematext.IsValidSematextRegion(region) {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Missing or invalid sematext_region parameter in provider stanza",
				Detail:   "Missing or invalid sematext_region parameter in provider stanza",
			})

			return nil, diags

		}

		token := os.Getenv("SEMATEXT_API_KEY")
		if !sematext.IsValidUUID(token) {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Missing or invalid env SEMATEXT_API_KEY",
				Detail:   "Missing or invalid env SEMATEXT_API_KEY",
			})

			return nil, diags

		}

		cfg := stcloud.NewConfiguration()

		if baseURL, err = url.Parse("https://apps.sematext.com"); err != nil {
			return nil, diag.FromErr(err)
		}

		switch region {
		case "US":
			baseURL, err = url.Parse("https://apps.sematext.com")
		case "EU":
			baseURL, err = url.Parse("https://apps.eu.sematext.com")
		default:
			err = errors.New("sematext_region must be one of EU, US")
		}

		if err != nil {
			return nil, diag.FromErr(err)
		}

		cfg.BasePath = baseURL.String()
		cfg.UserAgent = fmt.Sprintf("HashiCorp/1.0 Terraform/%s", tfVersionString)
		cfg.AddDefaultHeader("Authorization", fmt.Sprintf("apiKey %s", token))

		client := stcloud.NewAPIClient(cfg)

		return client, nil
	}

	return provider
}
