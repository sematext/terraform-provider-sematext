package sematext

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/sematext/terraform-provider-sematext/sematext.com/api"
)

// Provider  - TODO Doc Comment
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
			"sematext_monitor_infra": resourceSematextMonitorInfra(),
			// "sematext_monitor_docker":        resourceSematextMonitorDocker(),
			// "sematext_monitor_solr":          resourceSematextMonitorSolr(),
			// "sematext_monitor_solrcloud":     resourceSematextMonitorSolrCloud(),
			// "sematext_monitor_elasticsearch": resourceSematextMonitorElasticsearch(),
			// "sematext_monitor_akka":          resourceSematextMonitorAkka(),
			// "sematext_monitor_java":          resourceSematextMonitorJava(),
			// "sematext_monitor_tomcat":        resourceSematextMonitorTomcat(),
			// "sematext_monitor_zookeeper":     resourceSematextMonitorZookeeper(),
			// "sematext_monitor_kafka":         resourceSematextMonitorKafka(),
			// "sematext_monitor_clickhouse":    resourceSematextMonitorClickhouse(),
			// "sematext_monitor_aws":           resourceSematextMonitorAws(),
			// "sematext_monitor_cassandra":     resourceSematextMonitorCassandra(),
			// "sematext_monitor_nginx":         resourceSematextMonitorNginx(),
			// "sematext_monitor_nginxplus":     resourceSematextMonitorNginxPlus(),
			// "sematext_monitor_mongodb":       resourceSematextMonitorMongoDb(),
			// "sematext_monitor_nodejs":        resourceSematextMonitorNodejs(),
			// "sematext_monitor_apache":        resourceSematextMonitorApache(),
			// "sematext_monitor_mysql":         resourceSematextMonitorMysql(),
			// "sematext_monitor_hbase":         resourceSematextMonitorHbase(),
			// "sematext_monitor_redis":         resourceSematextMonitorRedis(),
			// "sematext_monitor_storm":         resourceSematextMonitorStorm(),
			// "sematext_monitor_haproxy":       resourceSematextMonitorHaproxy(),
			// "sematext_monitor_spark":         resourceSematextMonitorSpark(),
			// "sematext_monitor_yarn":          resourceSematextMonitorYarn(),
			// "sematext_monitor_hadoop":        resourceSematextMonitorHadoop(),
			// TODO Logger
		},
	}

	p.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {

		/*
			Terraform 0.12 introduced this field to the protocol
			We can therefore assume that if it's missing it's 0.10 or 0.11
		*/
		terraformVersion := p.TerraformVersion
		if terraformVersion == "" {
			terraformVersion = "0.11+compatible"
		}

		/*
			TODO enforce region as mandatory
			TODO API-token storage in state and get from ENV
		*/

		region := d.Get("sematext_region").(string)

		client := new(api.Client)
		err := client.Init(region, terraformVersion)
		if err != nil {
			return nil, err
		}

		return &client, nil
	}

	return p
}
