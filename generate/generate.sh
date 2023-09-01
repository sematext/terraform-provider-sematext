#/bin/bash

# This file is used to genrate the various provider files from a common template.

dir=`pwd`
cd "$(dirname "${BASH_SOURCE[0]}")"

apptypes=(
    "Akka"
    "Apache"
    "AWS EBS"
    "AWS EC2"
    "AWS ELB"
    "Cassandra"
    "ClickHouse"
    "Elastic Search"
    "Hadoop-MRv1"
    "Hadoop-YARN"
    "HAProxy"
    "HBase"
    "Infra"
    "JVM"
    "Kafka"
    "Logsene"
    "MongoDB"
    "MySQL"
    "Nginx"
    "Nginx-Plus"
    "Node.js"
    "Redis"
    "Solr"
    "SolrCloud"
    "Spark"
    "Storm"
    "Tomcat"
    "ZooKeeper"
    "postgresql"
    "rabbitmq"
    "mobile-logs"
)

# removed "Kafka-0.7.2" "Memcached" "SearchAnalytics" "Sensei" "Infra"

resourcelist=""
resourceimportlist=""


for apptype in "${apptypes[@]}"
do

    echo "Writing Terraform Resource and test file for $apptype"

    stripped=${apptype//[^[:alnum:]]/}
    lowercase=${stripped,,}
    titlecase=${lowercase^}
    classname=${titlecase}
    package=${lowercase}
    resourcename="sematext_app_${lowercase}"

    mkdir -p ../internal/resource/$package
    rm -rf ../internal/resources/$package/*

    sed -e "s/<<PACKAGE_NAME>>/${package}/g" -e "s/<<CLASS_NAME>>/${classname}/g" -e "s/<<APP_TYPE>>/${apptype}/g" -e "s/<<RESOURCE_NAME>>/${resourcename}/g" ./templates/resource_app.go.template > "../internal/resource/$package/resource.go"
    sed -e "s/<<PACKAGE_NAME>>/${package}/g" -e "s/<<CLASS_NAME>>/${classname}/g" -e "s/<<APP_TYPE>>/${apptype}/g" -e "s/<<RESOURCE_NAME>>/${resourcename}/g" ./templates/resource_app_test.go.template > "../internal/resource/$package/resource_test.go"

    resourceimportlist+="    \"github.com\\/sematext\\/terraform-provider-sematext\\/internal\\/resource\\/${package}\"\n"
    resourcelist+="        ${package}.NewResource,\\n"

done

echo "Rewriting Terraform Provider and Test"
sed -e "s/<<RESOURCE_IMPORTS>>/${resourceimportlist}/g" -e "s/<<RESOURCE_LIST>>/${resourcelist}/g" ./templates/provider.go.template > "../internal/provider/provider.go"
sed -e "s/<<RESOURCE_IMPORTS>>/${resourceimportlist}/g" -e "s/<<RESOURCE_LIST>>/${resourcelist}/g" ./templates/provider_test.go.template > "../internal/provider_test/provider_test.go"

cd ..
make fmt

cd $dir
