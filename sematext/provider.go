package sematext

import (
	//"github.com/hashicorp/terraform-plugin-sdk/helper/logging" // TODO Decide use of logging lib
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const (
	defaultProviderMaxOpenConnections = 4
	defaultExpectedAPIVersion         = "" //TODO Set expected API version
)

type ProviderConfig struct {
	Api APIConfig
}

func Provider() terraform.ResourceProvider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"sematext_region": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SEMATEXT_REGION", ""),
				Description: "The Sematext region, either US or EU.",
			},
		},

		// TODO Write data source
		// DataSourcesMap: map[string]*schema.Resource{
		//	"sematext_application_data": dataSourceSematextApp(),
		// },

		ResourcesMap: map[string]*schema.Resource{
			"sematext_monitor_infra":         resourceSematextMonitorInfra(),
			"sematext_monitor_docker":        resourceSematextMonitorDocker(),
			"sematext_monitor_solr":          resourceSematextMonitorSolr(),
			"sematext_monitor_solrcloud":     resourceSematextMonitorSolrCloud(),
			"sematext_monitor_elasticsearch": resourceSematextMonitorElasticsearch(),
			"sematext_monitor_akka":          resourceSematextMonitorAkka(),
			"sematext_monitor_java":          resourceSematextMonitorJava(),
			"sematext_monitor_tomcat":        resourceSematextMonitorTomcat(),
			"sematext_monitor_zookeeper":     resourceSematextMonitorZookeeper(),
			"sematext_monitor_kafka":         resourceSematextMonitorKafka(),
			"sematext_monitor_clickhouse":    resourceSematextMonitorClickhouse(),
			"sematext_monitor_aws":           resourceSematextMonitorAws(),
			"sematext_monitor_cassandra":     resourceSematextMonitorCassandra(),
			"sematext_monitor_nginx":         resourceSematextMonitorNginx(),
			"sematext_monitor_nginxplus":     resourceSematextMonitorNginxPlus(),
			"sematext_monitor_mongodb":       resourceSematextMonitorMongoDb(),
			"sematext_monitor_nodejs":        resourceSematextMonitorNodejs(),
			"sematext_monitor_apache":        resourceSematextMonitorApache(),
			"sematext_monitor_mysql":         resourceSematextMonitorMysql(),
			"sematext_monitor_hbase":         resourceSematextMonitorHbase(),
			"sematext_monitor_redis":         resourceSematextMonitorRedis(),
			"sematext_monitor_storm":         resourceSematextMonitorStorm(),
			"sematext_monitor_haproxy":       resourceSematextMonitorHaproxy(),
			"sematext_monitor_spark":         resourceSematextMonitorSpark(),
			"sematext_monitor_yarn":          resourceSematextMonitorYarn(),
			"sematext_monitor_hadoop":        resourceSematextMonitorHadoop(),
		},

		ConfigureFunc: providerConfigure,
	}

	p.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {

		terraformVersion := p.TerraformVersion
		if terraformVersion == "" {
			// Terraform 0.12 introduced this field to the protocol
			// We can therefore assume that if it's missing it's 0.10 or 0.11
			terraformVersion = "0.11+compatible"
		}
		region := d.Get("sematext_region").(string)
		pc := new(ProviderConfig)
		api := new(APIConfig)
		err := api.Init(region, terraformVersion)
		if err != nil {
			return nil, err
		}

		pc.Api = api

		return &pc, nil
	}

	return p
}
