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
    "Docker"
    "Elastic Search"
    "Hadoop-MRv1"
    "Hadoop-YARN"
    "HAProxy"
    "HBase"
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


for apptype in "${apptypes[@]}"
do

    echo "Writing Terraform Resource and test file for $apptype"

    stripped=${apptype//[^[:alnum:]]/}
    lowercase=${stripped,,}
    titlecase=${lowercase^}

    monitorfile="../sematext/resource_sematext_monitor_${lowercase}.go"
    testfile="../sematext/resource_sematext_monitor_${lowercase}_test.go"
    classname=${titlecase}
    resourcename="sematext_monitor_${lowercase}"

    sed -e "s/<<CLASS_NAME>>/${classname}/g" -e "s/<<APP_TYPE>>/${apptype}/g" -e "s/<<RESOURCE_NAME>>/${resourcename}/g" ./resource_sematext_monitor.go.template > $monitorfile
    sed -e "s/<<CLASS_NAME>>/${classname}/g" -e "s/<<APP_TYPE>>/${apptype}/g" -e "s/<<RESOURCE_NAME>>/${resourcename}/g" ./resource_sematext_monitor_test.go.template > $testfile

    resourcelist+="\"${resourcename}\": resourceSematextMonitor${classname}(),\\n"

done

echo "Rewriting Terraform Provider and test file"

sed -e "s/<<RESOURCE_LIST>>/${resourcelist}/g" ./provider.go.template > "../sematext/provider.go"

cd ..
make fmt


cd $dir
