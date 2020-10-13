# <img src="https://sematext.com/wp-content/uploads/2020/09/just-octi-blue.png" valign="bottom" width="60px"/>**&nbsp;&nbsp;Terraform Provider For Sematext Cloud**

# Overview

The Sematext provider is used to interact with resources supplied by [Sematext Cloud](https://sematext.com/cloud/).


## Example Usage

```hcl
terraform {
  required_providers {
    sematext = {
      source = "sematext/sematext"
      version = ">=0.1.3"
    }
  }
}

provider "sematext" {
    sematext_region = "US"
}

resource "sematext_monitor_nodejs" "mymonitor" {
    name = "Node.js Monitor Example"
    billing_plan_id = 6
}
```

## Argument Reference

The following arguments are supported:

* `sematext_region` - (Required) desired Sematext Cloud Region  "US" or "EU";


## Authentication

There are two authentication tokens

* Sematext Cloud API access token - [available from your account](https://apps.sematext.com/ui/account/api);
* Sematext Cloud App access token - retrieved on resource creation - refer to examples on how to access this inside your Terrform scripting.


## Enviropnment Variables

The following environment variables are required:

* SEMATEXT_REGION="US"
* SEMATEXT_API_KEY="&lt;Sematext-Cloud-Token&gt;"

If working with AWS Cloudwatch the following environment vars should be set:

* AWS_ACCESS_KEY_ID
* AWS_SECRET_ACCESS_KEY
* AWS_REGION


## Resources

The following terraform resources are supported.

* [Akka](../docs/resources/sematext_monitor_akka.md)
* [Apache](../docs/resources/sematext_monitor_apache.md)
* [AWS EBS](../docs/resources/sematext_monitor_awsebs.md)
* [AWS EC2](../docs/resources/sematext_monitor_awsec2.md)
* [AWS ELB](../docs/resources/sematext_monitor_awselb.md)
* [Cassandra](../docs/resources/sematext_monitor_cassandra.md)
* [ClickHouse](../docs/resources/sematext_monitor_clickhouse.md)
* [Docker](../docs/resources/sematext_monitor_docker.md)
* [Elastic Search](../docs/resources/sematext_monitor_elasticsearch.md)
* [Hadoop-MRv1](../docs/resources/sematext_monitor_hadoopmrv1.md)
* [Hadoop-YARN](../docs/resources/sematext_monitor_hadoopyarn.md)
* [HAProxy](../docs/resources/sematext_monitor_haproxy.md)
* [HBase](../docs/resources/sematext_monitor_hbase.md)
* [JVM](../docs/resources/sematext_monitor_jvm.md)
* [Kafka](../docs/resources/sematext_monitor_kafka.md)
* [Logsene](../docs/resources/sematext_monitor_logsene.md)
* [MongoDB](../docs/resources/sematext_monitor_mongodb.md)
* [MySQL](../docs/resources/sematext_monitor_mysql.md)
* [Nginx](../docs/resources/sematext_monitor_nginx.md)
* [Nginx-Plus](../docs/resources/sematext_monitor_nginxplus.md)
* [Node.js](../docs/resources/sematext_monitor_nodejs.md)
* [Redis](../docs/resources/sematext_monitor_redis.md)
* [Solr](../docs/resources/sematext_monitor_solr.md)
* [SolrCloud](../docs/resources/sematext_monitor_solrcloud.md)
* [Spark](../docs/resources/sematext_monitor_spark.md)
* [Storm](../docs/resources/sematext_monitor_storm.md)
* [Tomcat](../docs/resources/sematext_monitor_tomcat.md)
* [ZooKeeper](../docs/resources/sematext_monitor_zookeeper.md)
* [postgresql](../docs/resources/sematext_monitor_postgresql.md)
* [rabbitmq](../docs/resources/sematext_monitor_rabbitmq.md)
* [mobile-logs](../docs/resources/sematext_monitor_mobilelogs.md)

