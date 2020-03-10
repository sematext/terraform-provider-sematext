#/bin/bash


cd "$(dirname "${BASH_SOURCE[0]}")"

monitors=(
    Akka
    Apache
    AWS
    Cassandra
    Clickhouse
    Docker
    Elasticsearch
    Hadoop
    Haproxy
    Hbase
    Infra
    Java
    Kafka
    Logging
    Mongodb
    Mysql
    Nginx
    Nginxplus
    Nodejs
    Redis
    Solr
    Solrcloud
    Spark
    Storm
    Tomcat
    Yarn
    Zookeeper
)

for monitor in "${monitors[@]}"
do

    sed -e "s/<<MONITOR_TYPE>>/${monitor}/g" ./resource_sematext_monitor.go.template > ../sematext/resource_sematext_monitor_${monitor,,}.go
    sed -e "s/<<MONITOR_TYPE>>/${monitor}/g" ./resource_sematext_monitor_test.go.template > ../sematext/resource_sematext_monitor_${monitor,,}_test.go
    
done

cd -





