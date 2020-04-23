#/bin/bash

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
)

# removed "Kafka-0.7.2" "Memcached" "SearchAnalytics" "Sensei" "Infra"


awsparas='\* `aws_access_key` - (optional) if not set then reads from env TODO\n\* `aws_secret_key` - List attributes that this resource exports\n\* `aws_fetch_frequency` - (required) List attributes that this resource exports.\n\* `aws_region` - List attributes that this resource exports.'
echo $awsparas



for apptype in "${apptypes[@]}"
do

    echo "Writing Terraform Resource Doc for $apptype"

    stripped=${apptype//[^[:alnum:]]/}
    lowercase=${stripped,,}
    titlecase=${lowercase^}

    resourcetemplatefile="./resource_sematext_resource.md.template"
    classname=${titlecase}
    resourcename="sematext_monitor_${lowercase}"
    resourcetargetfile="../docs/resources/${resourcename}.md"

    sed -e "s/<<APP_TYPE>>/${apptype}/g" -e "s/<<RESOURCE_NAME>>/${resourcename}/g" $resourcetemplatefile > $resourcetargetfile
    
    #sed -e "s/<<APP_TYPE>>/${apptype}/g" -e "s/<<RESOURCE_NAME>>/${resourcename}/g" $resourcetemplatefile > $resourcetargetfile
    
    case "$apptype" in  "AWS EBS"|"AWS EC2"|"AWS ELB") : 
        sed -i -e "s/<<AWS_ARGSREF>>/${awsparas}/g" $resourcetargetfile
        ;;
    esac
    
    sed -i -e "s/<<AWS_ARGSREF>>//g" $resourcetargetfile


done

cd $dir
