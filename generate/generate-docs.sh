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

# removed "Kafka-0.7.2" "Memcached" "SearchAnalytics" "Sensei" 


awsparas='\* `aws_access_key` - (optional) if not set then reads from env AWS_ACCESS_KEY_ID.\n\* `aws_secret_key` - (optionl) is not present set from env AWS_SECRET_ACCESS_KEY\n\* `aws_fetch_frequency` - (required) one of MINUTE|FIVE_MINUTES|FIFTEEN_MINUTES.\n\* `aws_region` - (optional) if not present withh set from env AWS_REGION.'

templatefile='./provider_index.md.template'
targetfile='../docs/index.md'

for apptype in "${apptypes[@]}"
do

    echo "Writing Terraform Resource Docs for $apptype"

    stripped=${apptype//[^[:alnum:]]/}
    lowercase=${stripped,,}
    titlecase=${lowercase^}

    resourcetemplatefile="./resource_app.md.template"
    classname=${titlecase}
    resourcename="sematext_app_${lowercase}"
    targetfile="../docs/resources/${resourcename}.md"
    resourcelinks+="\* [$apptype](..\/docs\/resources\/${resourcename}.md)\n"

    sed -e "s/<<APP_TYPE>>/${apptype}/g" -e "s/<<RESOURCE_NAME>>/${resourcename}/g" $templatefile > $targetfile

    case "$apptype" in  "AWS EBS"|"AWS EC2"|"AWS ELB") :
        sed -i -e "s/<<AWS_ARGSREF>>/${awsparas}/g" $targetfile
        ;;
    esac

    sed -i -e "s/<<AWS_ARGSREF>>//g" $targetfile


done

echo ${resourcelinks}
echo $templatefile
echo $targetfile

sed -e "s/<<RESOURCE_LINKS>>/${resourcelinks}/g" $templatefile > $targetfile


cd $dir
