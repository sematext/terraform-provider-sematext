# Sematext Provider

The Sematext provider is used to interact with Sematext Cloud related resources.


## Example Usage

```hcl
# Configure the Sematext Provider
provider "sematext" {
  sematext_region = "US"
}

# Create a monitoring application
resource "sematext_monitor_nodejs" "mymonitor" {
  name = "mymonitor"
  billing_plan_id = 6
  discount_code = "<discount code>"
}
```

## Argument Reference

The following arguments are supported:

* `sematext_region` - (Required) desired Sematext Cloud Region  "US" or "EU";


## Authentication

There are two authentication tokens

* Sematext Cloud API access token - [available from your account](https://apps.sematext.com/ui/account/api)

When the Sematext Terraform Provider connects to Sematext Cloud it authenticates with an access token.
You'll need to gain this via your login to your Sematext Cloud account. TODO
Note this is different from the app key


## Enviropnment Variables

The following environment variables are required:

* SEMATEXT_REGION="US"
* SEMATEXT_API_TOKEN="&lt;Sematext-Cloud-Token&gt;"

If working with AWS Cloudwatch the following environment vars should be set:

* AWS_ACCESS_KEY_ID
* AWS_SECRET_ACCESS_KEY
* AWS_REGION


## Resources 

The following terraform resources are supported.

* [Akka](./resources/sematext_monitor_akka.md)
* [Apache](./resources/sematext_monitor_apache.md)
* [AWS EBS](./resources/sematext_monitor_awsebs.md)
* [AWS EC2](./resources/sematext_monitor_awsec2.md)
* [AWS ELB](./resources/sematext_monitor_awselb.md)
* [Cassandra](./resources/sematext_monitor_cassandra.md)
* [ClickHouse](./resources/sematext_monitor_clickhouse.md)
* [Docker](./resources/sematext_monitor_docker.md)
* [Elastic Search](./resources/sematext_monitor_elasticsearch.md)
* [Hadoop-MRv1](./resources/sematext_monitor_hadoopmrv1.md)
* [Hadoop-YARN](./resources/sematext_monitor_hadoopyarn.md)
* [HAProxy](./resources/sematext_monitor_haproxy.md)
* [HBase](./resources/sematext_monitor_hbase.md)
* [JVM](./resources/sematext_monitor_jvm.md)
* [Kafka](./resources/sematext_monitor_kafka.md)
* [Logsene](./resources/sematext_monitor_logsene.md)
* [MongoDB](./resources/sematext_monitor_mongodb.md)
* [MySQL](./resources/sematext_monitor_mysql.md)
* [Nginx](./resources/sematext_monitor_nginx.md)
* [Nginx-Plus](./resources/sematext_monitor_nginxplus.md)
* [Node.js](./resources/sematext_monitor_nodejs.md)
* [Redis](./resources/sematext_monitor_redis.md)
* [Solr](./resources/sematext_monitor_solr.md)
* [SolrCloud](./resources/sematext_monitor_solrcloud.md)
* [Spark](./resources/sematext_monitor_spark.md)
* [Storm](./resources/sematext_monitor_storm.md)
* [Tomcat](./resources/sematext_monitor_tomcat.md)
* [ZooKeeper](./resources/sematext_monitor_zookeeper.md)


