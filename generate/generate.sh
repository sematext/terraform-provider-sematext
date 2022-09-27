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


for apptype in "${apptypes[@]}"
do

    echo "Writing Terraform Resource and test file for $apptype"

    stripped=${apptype//[^[:alnum:]]/}
    lowercase=${stripped,,}
    titlecase=${lowercase^}
    classname=${titlecase}

    resourcename="sematext_app_${lowercase}"
    sed -e "s/<<CLASS_NAME>>/${classname}/g" -e "s/<<APP_TYPE>>/${apptype}/g" -e "s/<<RESOURCE_NAME>>/${resourcename}/g" ./resource_app.go.template > "../sematext/generated/resource_app_${lowercase}.go"
    sed -e "s/<<CLASS_NAME>>/${classname}/g" -e "s/<<APP_TYPE>>/${apptype}/g" -e "s/<<RESOURCE_NAME>>/${resourcename}/g" ./resource_app_test.go.template > "../sematext/generated/resource_app_${lowercase}_test.go"
    resourcelist+="\"${resourcename}\": resourceApp${classname}(),\\n"

    resourcename="sematext_alertrule_${lowercase}"
    sed -e "s/<<CLASS_NAME>>/${classname}/g" -e "s/<<APP_TYPE>>/${apptype}/g" -e "s/<<RESOURCE_NAME>>/${resourcename}/g" ./resource_alertrule.go.template > "../sematext/generated/resource_alertrule_${lowercase}.go"
    sed -e "s/<<CLASS_NAME>>/${classname}/g" -e "s/<<APP_TYPE>>/${apptype}/g" -e "s/<<RESOURCE_NAME>>/${resourcename}/g" ./resource_alertrule_test.go.template > "../sematext/generated/resource_alertrule_${lowercase}_test.go"
    resourcelist+="\"${resourcename}\": resourceAlertRule${classname}(),\\n"


done

echo "Rewriting Terraform Provider and test file"

sed -e "s/<<RESOURCE_LIST>>/${resourcelist}/g" ./provider.go.template > "../sematext/generated/provider.go"

cd ..
make fmt


cd $dir
