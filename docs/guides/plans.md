# <img src="../assets/octi-blue.png" valign="bottom" width="60px"/>**&nbsp;&nbsp;Terraform Provider For Sematext Cloud**


The Sematext provider is used to interact with resources supplied by [Sematext Cloud](https://sematext.com/cloud/).

## Plan Reference

The following table containes valid plan id values to be used in resource definitions.
<br>

| PlanID | &nbsp;&nbsp;Resource                       |          Plan Name           |
| -----: | :----------------------------------------- | :--------------------------: |
|    121 | &nbsp;&nbsp;sematext_monitor_akka          |        Basic SPM Akka        |
|    122 | &nbsp;&nbsp;sematext_monitor_akka          |      Standard SPM Akka       |
|    123 | &nbsp;&nbsp;sematext_monitor_akka          |    Professional SPM Akka     |
|    124 | &nbsp;&nbsp;sematext_monitor_akka          |     Pro SPM Silver Akka      |
|     81 | &nbsp;&nbsp;sematext_monitor_apache        |       Basic SPM Apache       |
|     82 | &nbsp;&nbsp;sematext_monitor_apache        |     Standard SPM Apache      |
|     83 | &nbsp;&nbsp;sematext_monitor_apache        |        Pro SPM Apache        |
|     84 | &nbsp;&nbsp;sematext_monitor_apache        |    Pro Silver SPM Apache     |
|     73 | &nbsp;&nbsp;sematext_monitor_aswebs        |        Basic SPM AWS         |
|     74 | &nbsp;&nbsp;sematext_monitor_aswebs        |       Standard SPM AWS       |
|     75 | &nbsp;&nbsp;sematext_monitor_aswebs        |         Pro SPM AWS          |
|     76 | &nbsp;&nbsp;sematext_monitor_aswebs        |      Pro Silver SPM AWS      |
|     65 | &nbsp;&nbsp;sematext_monitor_awsec2        |        Basic SPM AWS         |
|     66 | &nbsp;&nbsp;sematext_monitor_awsec2        |       Standard SPM AWS       |
|     67 | &nbsp;&nbsp;sematext_monitor_awsec2        |         Pro SPM AWS          |
|     68 | &nbsp;&nbsp;sematext_monitor_awsec2        |      Pro Silver SPM AWS      |
|     69 | &nbsp;&nbsp;sematext_monitor_awselb        |        Basic SPM AWS         |
|     70 | &nbsp;&nbsp;sematext_monitor_awselb        |       Standard SPM AWS       |
|     71 | &nbsp;&nbsp;sematext_monitor_awselb        |         Pro SPM AWS          |
|     72 | &nbsp;&nbsp;sematext_monitor_awselb        |      Pro Silver SPM AWS      |
|     57 | &nbsp;&nbsp;sematext_monitor_cassandra     |     Basic SPM Cassandra      |
|     58 | &nbsp;&nbsp;sematext_monitor_cassandra     |    Standard SPM Cassandra    |
|     59 | &nbsp;&nbsp;sematext_monitor_cassandra     |      Pro SPM Cassandra       |
|     60 | &nbsp;&nbsp;sematext_monitor_cassandra     |   Pro Silver SPM Cassandra   |
|    130 | &nbsp;&nbsp;sematext_monitor_clickhouse    |     Basic SPM ClickHouse     |
|    131 | &nbsp;&nbsp;sematext_monitor_clickhouse    |   Standard SPM ClickHouse    |
|    132 | &nbsp;&nbsp;sematext_monitor_clickhouse    |      Pro SPM ClickHouse      |
|    133 | &nbsp;&nbsp;sematext_monitor_clickhouse    |  Pro Silver SPM ClickHouse   |
|    117 | &nbsp;&nbsp;sematext_monitor_docker        |       Basic SPM Docker       |
|    118 | &nbsp;&nbsp;sematext_monitor_docker        |     Standard SPM Docker      |
|    119 | &nbsp;&nbsp;sematext_monitor_docker        |        Pro SPM Docker        |
|    120 | &nbsp;&nbsp;sematext_monitor_docker        |    Pro Silver SPM Docker     |
|     12 | &nbsp;&nbsp;sematext_monitor_elasticsearch |   Basic SPM Elasticsearch    |
|     13 | &nbsp;&nbsp;sematext_monitor_elasticsearch |  Standard SPM Elasticsearch  |
|     14 | &nbsp;&nbsp;sematext_monitor_elasticsearch |    Pro SPM Elasticsearch     |
|     29 | &nbsp;&nbsp;sematext_monitor_elasticsearch | Pro Silver SPM Elasticsearch |
|     21 | &nbsp;&nbsp;sematext_monitor_hadoopmrv1    |       Basic SPM Hadoop       |
|     22 | &nbsp;&nbsp;sematext_monitor_hadoopmrv1    |     Standard SPM Hadoop      |
|     23 | &nbsp;&nbsp;sematext_monitor_hadoopmrv1    |        Pro SPM Hadoop        |
|     32 | &nbsp;&nbsp;sematext_monitor_hadoopmrv1    |    Pro Silver SPM Hadoop     |
|     24 | &nbsp;&nbsp;sematext_monitor_hadoopyarn    |       Basic SPM Hadoop       |
|     25 | &nbsp;&nbsp;sematext_monitor_hadoopyarn    |     Standard SPM Hadoop      |
|     26 | &nbsp;&nbsp;sematext_monitor_hadoopyarn    |        Pro SPM Hadoop        |
|     33 | &nbsp;&nbsp;sematext_monitor_hadoopyarn    |    Pro Silver SPM Hadoop     |
|    105 | &nbsp;&nbsp;sematext_monitor_haproxy       |      Basic SPM HAProxy       |
|    106 | &nbsp;&nbsp;sematext_monitor_haproxy       |     Standard SPM HAProxy     |
|    107 | &nbsp;&nbsp;sematext_monitor_haproxy       |       Pro SPM HAProxy        |
|    108 | &nbsp;&nbsp;sematext_monitor_haproxy       |    Pro Silver SPM HAProxy    |
|      4 | &nbsp;&nbsp;sematext_monitor_hbase         |       Basic SPM HBase        |
|      5 | &nbsp;&nbsp;sematext_monitor_hbase         |      Standard SPM HBase      |
|      6 | &nbsp;&nbsp;sematext_monitor_hbase         |        Pro SPM HBase         |
|     28 | &nbsp;&nbsp;sematext_monitor_hbase         |     Pro Silver SPM HBase     |
|     18 | &nbsp;&nbsp;sematext_monitor_jvm           |        Basic SPM JVM         |
|     19 | &nbsp;&nbsp;sematext_monitor_jvm           |       Standard SPM JVM       |
|     20 | &nbsp;&nbsp;sematext_monitor_jvm           |         Pro SPM JVM          |
|     31 | &nbsp;&nbsp;sematext_monitor_jvm           |      Pro Silver SPM JVM      |
|     93 | &nbsp;&nbsp;sematext_monitor_kafka         |       Basic SPM Kafka        |
|     94 | &nbsp;&nbsp;sematext_monitor_kafka         |      Standard SPM Kafka      |
|     95 | &nbsp;&nbsp;sematext_monitor_kafka         |        Pro SPM Kafka         |
|     96 | &nbsp;&nbsp;sematext_monitor_kafka         |     Pro Silver SPM Kafka     |
|  10100 | &nbsp;&nbsp;sematext_monitor_logsene       |  Basic, 7 days, 500 MB/day   |
|  10101 | &nbsp;&nbsp;sematext_monitor_logsene       |  Standard, 7 days, 1 GB/day  |
|  10102 | &nbsp;&nbsp;sematext_monitor_logsene       |  Standard, 7 days, 5 GB/day  |
|  10103 | &nbsp;&nbsp;sematext_monitor_logsene       | Standard, 7 days, 10 GB/day  |
|  10104 | &nbsp;&nbsp;sematext_monitor_logsene       | Standard, 15 days, 1 GB/day  |
|  10105 | &nbsp;&nbsp;sematext_monitor_logsene       | Standard, 15 days, 5 GB/day  |
|  10106 | &nbsp;&nbsp;sematext_monitor_logsene       | Standard, 15 days, 10 GB/day |
|  10107 | &nbsp;&nbsp;sematext_monitor_logsene       |    Pro, 7 days, 1 GB/day     |
|  10108 | &nbsp;&nbsp;sematext_monitor_logsene       |    Pro, 7 days, 5 GB/day     |
|  10109 | &nbsp;&nbsp;sematext_monitor_logsene       |    Pro, 7 days, 10 GB/day    |
|  10110 | &nbsp;&nbsp;sematext_monitor_logsene       |    Pro, 7 days, 15 GB/day    |
|  10111 | &nbsp;&nbsp;sematext_monitor_logsene       |    Pro, 7 days, 20 GB/day    |
|  10112 | &nbsp;&nbsp;sematext_monitor_logsene       |    Pro, 7 days, 35 GB/day    |
|  10113 | &nbsp;&nbsp;sematext_monitor_logsene       |    Pro, 7 days, 50 GB/day    |
|  10114 | &nbsp;&nbsp;sematext_monitor_logsene       |    Pro, 7 days, 75 GB/day    |
|  10115 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 7 days, 100 GB/day    |
|  10116 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 7 days, 150 GB/day    |
|  10117 | &nbsp;&nbsp;sematext_monitor_logsene       |    Pro, 15 days, 1 GB/day    |
|  10118 | &nbsp;&nbsp;sematext_monitor_logsene       |    Pro, 15 days, 5 GB/day    |
|  10119 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 15 days, 10 GB/day    |
|  10120 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 15 days, 15 GB/day    |
|  10121 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 15 days, 20 GB/day    |
|  10122 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 15 days, 35 GB/day    |
|  10123 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 15 days, 50 GB/day    |
|  10124 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 15 days, 75 GB/day    |
|  10125 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 15 days, 100 GB/day   |
|  10126 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 15 days, 150 GB/day   |
|  10127 | &nbsp;&nbsp;sematext_monitor_logsene       |    Pro, 30 days, 1 GB/day    |
|  10128 | &nbsp;&nbsp;sematext_monitor_logsene       |    Pro, 30 days, 5 GB/day    |
|  10129 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 30 days, 10 GB/day    |
|  10130 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 30 days, 15 GB/day    |
|  10131 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 30 days, 20 GB/day    |
|  10132 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 30 days, 35 GB/day    |
|  10133 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 30 days, 50 GB/day    |
|  10134 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 30 days, 75 GB/day    |
|  10135 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 30 days, 100 GB/day   |
|  10136 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 30 days, 150 GB/day   |
|  10137 | &nbsp;&nbsp;sematext_monitor_logsene       |    Pro, 60 days, 1 GB/day    |
|  10138 | &nbsp;&nbsp;sematext_monitor_logsene       |    Pro, 60 days, 5 GB/day    |
|  10139 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 60 days, 10 GB/day    |
|  10140 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 60 days, 15 GB/day    |
|  10141 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 60 days, 20 GB/day    |
|  10142 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 60 days, 35 GB/day    |
|  10143 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 60 days, 50 GB/day    |
|  10144 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 60 days, 75 GB/day    |
|  10145 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 60 days, 100 GB/day   |
|  10146 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 60 days, 150 GB/day   |
|  10147 | &nbsp;&nbsp;sematext_monitor_logsene       |    Pro, 90 days, 1 GB/day    |
|  10148 | &nbsp;&nbsp;sematext_monitor_logsene       |    Pro, 90 days, 5 GB/day    |
|  10149 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 90 days, 10 GB/day    |
|  10150 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 90 days, 15 GB/day    |
|  10151 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 90 days, 20 GB/day    |
|  10152 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 90 days, 35 GB/day    |
|  10153 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 90 days, 50 GB/day    |
|  10154 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 90 days, 75 GB/day    |
|  10155 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 90 days, 100 GB/day   |
|  10156 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 90 days, 150 GB/day   |
|  10157 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 180 days, 1 GB/day    |
|  10158 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 180 days, 5 GB/day    |
|  10159 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 180 days, 10 GB/day   |
|  10160 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 180 days, 15 GB/day   |
|  10161 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 180 days, 20 GB/day   |
|  10162 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 180 days, 35 GB/day   |
|  10163 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 180 days, 50 GB/day   |
|  10164 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 180 days, 75 GB/day   |
|  10165 | &nbsp;&nbsp;sematext_monitor_logsene       |  Pro, 180 days, 100 GB/day   |
|  10166 | &nbsp;&nbsp;sematext_monitor_logsene       |  Pro, 180 days, 150 GB/day   |
|  10167 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 365 days, 1 GB/day    |
|  10168 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 365 days, 5 GB/day    |
|  10169 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 365 days, 10 GB/day   |
|  10170 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 365 days, 15 GB/day   |
|  10171 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 365 days, 20 GB/day   |
|  10172 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 365 days, 35 GB/day   |
|  10173 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 365 days, 50 GB/day   |
|  10174 | &nbsp;&nbsp;sematext_monitor_logsene       |   Pro, 365 days, 75 GB/day   |
|  10175 | &nbsp;&nbsp;sematext_monitor_logsene       |  Pro, 365 days, 100 GB/day   |
|  10176 | &nbsp;&nbsp;sematext_monitor_logsene       |  Pro, 365 days, 150 GB/day   |
|    125 | &nbsp;&nbsp;sematext_monitor_mongodb       |      Basic SPM MongoDB       |
|    126 | &nbsp;&nbsp;sematext_monitor_mongodb       |     Standard SPM MongoDB     |
|    127 | &nbsp;&nbsp;sematext_monitor_mongodb       |       Pro SPM MongoDB        |
|    128 | &nbsp;&nbsp;sematext_monitor_mongodb       |    Pro Silver SPM MongoDB    |
|     77 | &nbsp;&nbsp;sematext_monitor_mysql         |       Basic SPM MySQL        |
|     78 | &nbsp;&nbsp;sematext_monitor_mysql         |      Standard SPM MySQL      |
|     79 | &nbsp;&nbsp;sematext_monitor_mysql         |        Pro SPM MySQL         |
|     80 | &nbsp;&nbsp;sematext_monitor_mysql         |     Pro Silver SPM MySQL     |
|     85 | &nbsp;&nbsp;sematext_monitor_nginx         |       Basic SPM Nginx        |
|     86 | &nbsp;&nbsp;sematext_monitor_nginx         |      Standard SPM Nginx      |
|     87 | &nbsp;&nbsp;sematext_monitor_nginx         |        Pro SPM Nginx         |
|     88 | &nbsp;&nbsp;sematext_monitor_nginx         |     Pro Silver SPM Nginx     |
|     89 | &nbsp;&nbsp;sematext_monitor_nginxplus     |     Basic SPM Nginx Plus     |
|     90 | &nbsp;&nbsp;sematext_monitor_nginxplus     |   Standard SPM Nginx Plus    |
|     91 | &nbsp;&nbsp;sematext_monitor_nginxplus     |      Pro SPM Nginx Plus      |
|     92 | &nbsp;&nbsp;sematext_monitor_nginxplus     |  Pro Silver SPM Nginx Plus   |
|    109 | &nbsp;&nbsp;sematext_monitor_nodejs        |      Basic SPM Node.js       |
|    110 | &nbsp;&nbsp;sematext_monitor_nodejs        |     Standard SPM Node.js     |
|    111 | &nbsp;&nbsp;sematext_monitor_nodejs        |       Pro SPM Node.js        |
|    112 | &nbsp;&nbsp;sematext_monitor_nodejs        |    Pro Silver SPM Node.js    |
|     49 | &nbsp;&nbsp;sematext_monitor_redis         |       Basic SPM Redis        |
|     50 | &nbsp;&nbsp;sematext_monitor_redis         |      Standard SPM Redis      |
|     51 | &nbsp;&nbsp;sematext_monitor_redis         |        Pro SPM Redis         |
|     52 | &nbsp;&nbsp;sematext_monitor_redis         |     Pro Silver SPM Redis     |
|      1 | &nbsp;&nbsp;sematext_monitor_solr          |        Basic SPM Solr        |
|      2 | &nbsp;&nbsp;sematext_monitor_solr          |      Standard SPM Solr       |
|      3 | &nbsp;&nbsp;sematext_monitor_solr          |         Pro SPM Solr         |
|     27 | &nbsp;&nbsp;sematext_monitor_solr          |     Pro Silver SPM Solr      |
|     45 | &nbsp;&nbsp;sematext_monitor_solrcloud     |     Basic SPM SolrCloud      |
|     46 | &nbsp;&nbsp;sematext_monitor_solrcloud     |    Standard SPM SolrCloud    |
|     47 | &nbsp;&nbsp;sematext_monitor_solrcloud     |      Pro SPM SolrCloud       |
|     48 | &nbsp;&nbsp;sematext_monitor_solrcloud     |   Pro Silver SPM SolrCloud   |
|     97 | &nbsp;&nbsp;sematext_monitor_spark         |       Basic SPM Spark        |
|     98 | &nbsp;&nbsp;sematext_monitor_spark         |      Standard SPM Spark      |
|     99 | &nbsp;&nbsp;sematext_monitor_spark         |        Pro SPM Spark         |
|    100 | &nbsp;&nbsp;sematext_monitor_spark         |     Pro Silver SPM Spark     |
|     53 | &nbsp;&nbsp;sematext_monitor_storm         |       Basic SPM Storm        |
|     54 | &nbsp;&nbsp;sematext_monitor_storm         |      Standard SPM Storm      |
|     55 | &nbsp;&nbsp;sematext_monitor_storm         |        Pro SPM Storm         |
|     56 | &nbsp;&nbsp;sematext_monitor_storm         |     Pro Silver SPM Storm     |
|    113 | &nbsp;&nbsp;sematext_monitor_tomcat        |       Basic SPM Tomcat       |
|    114 | &nbsp;&nbsp;sematext_monitor_tomcat        |     Standard SPM Tomcat      |
|    115 | &nbsp;&nbsp;sematext_monitor_tomcat        |        Pro SPM Tomcat        |
|    116 | &nbsp;&nbsp;sematext_monitor_tomcat        |    Pro Silver SPM Tomcat     |
|     41 | &nbsp;&nbsp;sematext_monitor_zookeeper     |     Basic SPM ZooKeeper      |
|     42 | &nbsp;&nbsp;sematext_monitor_zookeeper     |    Standard SPM ZooKeeper    |
|     43 | &nbsp;&nbsp;sematext_monitor_zookeeper     |      Pro SPM ZooKeeper       |
|     44 | &nbsp;&nbsp;sematext_monitor_zookeeper     |   Pro Silver SPM ZooKeeper   |
