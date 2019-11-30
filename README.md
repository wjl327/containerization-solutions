elk:
    docker pull docker.elastic.co/elasticsearch/elasticsearch:7.4.2
    docker pull docker.elastic.co/beats/filebeat:7.4.2
    docker pull docker.elastic.co/logstash/logstash:7.4.2
    docker pull docker.elastic.co/kibana/kibana:7.4.2

    启动命令: docker-compose up -d
    停止命令：docker-compose down
