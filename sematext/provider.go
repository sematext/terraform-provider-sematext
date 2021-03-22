package sematext

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
					if val == nil || !IsValidSematextRegion(region) {
						errs = append(errs, fmt.Errorf("ERROR  : sematext_region missing or invalid in sematext provider"))
					}
					return
				},
				Description: "The Sematext region, either US or EU.",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"sematext_monitor_akka":          resourceSematextMonitorAkka(),
			"sematext_monitor_apache":        resourceSematextMonitorApache(),
			"sematext_monitor_awsebs":        resourceSematextMonitorAwsebs(),
			"sematext_monitor_awsec2":        resourceSematextMonitorAwsec2(),
			"sematext_monitor_awselb":        resourceSematextMonitorAwselb(),
			"sematext_monitor_cassandra":     resourceSematextMonitorCassandra(),
			"sematext_monitor_clickhouse":    resourceSematextMonitorClickhouse(),
			"sematext_monitor_elasticsearch": resourceSematextMonitorElasticsearch(),
			"sematext_monitor_hadoopmrv1":    resourceSematextMonitorHadoopmrv1(),
			"sematext_monitor_hadoopyarn":    resourceSematextMonitorHadoopyarn(),
			"sematext_monitor_haproxy":       resourceSematextMonitorHaproxy(),
			"sematext_monitor_hbase":         resourceSematextMonitorHbase(),
			"sematext_monitor_infra":         resourceSematextMonitorInfra(),
			"sematext_monitor_jvm":           resourceSematextMonitorJvm(),
			"sematext_monitor_kafka":         resourceSematextMonitorKafka(),
			"sematext_monitor_logsene":       resourceSematextMonitorLogsene(),
			"sematext_monitor_mongodb":       resourceSematextMonitorMongodb(),
			"sematext_monitor_mysql":         resourceSematextMonitorMysql(),
			"sematext_monitor_nginx":         resourceSematextMonitorNginx(),
			"sematext_monitor_nginxplus":     resourceSematextMonitorNginxplus(),
			"sematext_monitor_nodejs":        resourceSematextMonitorNodejs(),
			"sematext_monitor_redis":         resourceSematextMonitorRedis(),
			"sematext_monitor_solr":          resourceSematextMonitorSolr(),
			"sematext_monitor_solrcloud":     resourceSematextMonitorSolrcloud(),
			"sematext_monitor_spark":         resourceSematextMonitorSpark(),
			"sematext_monitor_storm":         resourceSematextMonitorStorm(),
			"sematext_monitor_tomcat":        resourceSematextMonitorTomcat(),
			"sematext_monitor_zookeeper":     resourceSematextMonitorZookeeper(),
			"sematext_monitor_postgresql":    resourceSematextMonitorPostgresql(),
			"sematext_monitor_rabbitmq":      resourceSematextMonitorRabbitmq(),
			"sematext_monitor_mobilelogs":    resourceSematextMonitorMobilelogs(),
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
		if !IsValidSematextRegion(region) {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Missing or invalid sematext_region parameter in provider stanza",
				Detail:   "Missing or invalid sematext_region parameter in provider stanza",
			})

			return nil, diags

		}

		token := os.Getenv("SEMATEXT_API_KEY")
		if !IsValidUUID(token) {
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
