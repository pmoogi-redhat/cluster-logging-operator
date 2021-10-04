package kafka

const (
	initKafkaScript = `
      #!/bin/bash
      set -e
      cp /etc/kafka-configmap/log4j.properties /etc/kafka/

      KAFKA_BROKER_ID=${HOSTNAME##*-}
      SEDS=("s/#init#broker.id=#init#/broker.id=$KAFKA_BROKER_ID/")
      LABELS="kafka-broker-id=$KAFKA_BROKER_ID"
      ANNOTATIONS=""

      hash kubectl 2>/dev/null || {
        SEDS+=("s/#init#broker.rack=#init#/#init#broker.rack=# kubectl not found in path/")
      } && {
        ZONE=$(kubectl get node "$NODE_NAME" -o=go-template='{{index .metadata.labels "failure-domain.beta.kubernetes.io/zone"}}')
        if [ "x$ZONE" == "x<no value>" ]; then
          SEDS+=("s/#init#broker.rack=#init#/#init#broker.rack=# zone label not found for node $NODE_NAME/")
        else
          SEDS+=("s/#init#broker.rack=#init#/broker.rack=$ZONE/")
          LABELS="$LABELS kafka-broker-rack=$ZONE"
        fi

        [ -z "$ADVERTISE_ADDR" ] && echo "ADVERTISE_ADDR is empty, will advertise detected DNS name"
        SEDS+=("s|#init#advertised.listeners=PLAINTEXT://#init#|advertised.listeners=PLAINTEXT://${ADVERTISE_ADDR}:9092,SSL://${ADVERTISE_ADDR}:9093|")

        if [ ! -z "$LABELS" ]; then
          kubectl -n $POD_NAMESPACE label pod $POD_NAME $LABELS || echo "Failed to label $POD_NAMESPACE.$POD_NAME - RBAC issue?"
        fi
        if [ ! -z "$ANNOTATIONS" ]; then
          kubectl -n $POD_NAMESPACE annotate pod $POD_NAME $ANNOTATIONS || echo "Failed to annotate $POD_NAMESPACE.$POD_NAME - RBAC issue?"
        fi
      }
      printf '%s\n' "${SEDS[@]}" | sed -f - /etc/kafka-configmap/server.properties > /etc/kafka/server.properties.tmp
      [ $? -eq 0 ] && mv /etc/kafka/server.properties.tmp /etc/kafka/server.properties

      rm -rf /var/lib/kafka/data/*
    `

	clientProperties = `
      security.protocol=SSL
      ssl.truststore.location=/etc/kafka-certs/ca-bundle.jks
      ssl.truststore.type=JKS
      ssl.truststore.password=ca-bundle
    `

	serverProperties = `
      ############################# Log Basics #############################

      # A comma separated list of directories under which to store log files
      # Overrides log.dir
      log.dirs=/var/lib/kafka/data/topics

      # The default number of log partitions per topic. More partitions allow greater
      # parallelism for consumption, but this will also result in more files across
      # the brokers.
      num.partitions=12

      default.replication.factor=1

      min.insync.replicas=1

      auto.create.topics.enable=false

      # Max Messages in Bytes set to 10M > fluentd buffer chunk_limit_size config
      message.max.bytes=10000000

      # The number of threads per data directory to be used for log recovery at startup and flushing at shutdown.
      # This value is recommended to be increased for installations with data dirs located in RAID array.
      #num.recovery.threads.per.data.dir=1

      ############################# Server Basics #############################

      # The id of the broker. This must be set to a unique integer for each broker.
      #init#broker.id=#init#

      #init#broker.rack=#init#

      ############################# Socket Server Settings #############################

      # The address the socket server listens on. It will get the value returned from
      # java.net.InetAddress.getCanonicalHostName() if not configured.
      #   FORMAT:
      #     listeners = listener_name://host_name:port
      #   EXAMPLE:
      #     listeners = PLAINTEXT://your.host.name:9092
      #listeners=PLAINTEXT://:9092
      listeners=PLAINTEXT://:9092,SSL://:9093
      ssl.keystore.type=JKS
      ssl.keystore.location=/etc/kafka-certs/server.jks
      ssl.keystore.password=server

      # Hostname and port the broker will advertise to producers and consumers. If not set,
      # it uses the value for "listeners" if configured.  Otherwise, it will use the value
      # returned from java.net.InetAddress.getCanonicalHostName().
      #advertised.listeners=PLAINTEXT://your.host.name:9092
      #init#advertised.listeners=PLAINTEXT://#init#

      # Maps listener names to security protocols, the default is for them to be the same. See the config documentation for more details
      #listener.security.protocol.map=PLAINTEXT:PLAINTEXT,SSL:SSL,SASL_PLAINTEXT:SASL_PLAINTEXT,SASL_SSL:SASL_SSL
      listener.security.protocol.map=PLAINTEXT:PLAINTEXT,SSL:SSL,SASL_PLAINTEXT:SASL_PLAINTEXT,SASL_SSL:SASL_SSL,OUTSIDE:PLAINTEXT
      inter.broker.listener.name=PLAINTEXT

      # The number of threads that the server uses for receiving requests from the network and sending responses to the network
      #num.network.threads=3

      # The number of threads that the server uses for processing requests, which may include disk I/O
      #num.io.threads=8

      # The send buffer (SO_SNDBUF) used by the socket server
      #socket.send.buffer.bytes=102400

      # The receive buffer (SO_RCVBUF) used by the socket server
      #socket.receive.buffer.bytes=102400

      # The maximum size of a request that the socket server will accept (protection against OOM)
      #socket.request.max.bytes=104857600

      ############################# Internal Topic Settings  #############################
      # The replication factor for the group metadata internal topics "__consumer_offsets" and "__transaction_state"
      # For anything other than development testing, a value greater than 1 is recommended for to ensure availability such as 3.
      offsets.topic.replication.factor=1
      transaction.state.log.replication.factor=1
      transaction.state.log.min.isr=1

      ############################# Log Flush Policy #############################

      # Messages are immediately written to the filesystem but by default we only fsync() to sync
      # the OS cache lazily. The following configurations control the flush of data to disk.
      # There are a few important trade-offs here:
      #    1. Durability: Unflushed data may be lost if you are not using replication.
      #    2. Latency: Very large flush intervals may lead to latency spikes when the flush does occur as there will be a lot of data to flush.
      #    3. Throughput: The flush is generally the most expensive operation, and a small flush interval may lead to excessive seeks.
      # The settings below allow one to configure the flush policy to flush data after a period of time or
      # every N messages (or both). This can be done globally and overridden on a per-topic basis.

      # The number of messages to accept before forcing a flush of data to disk
      #log.flush.interval.messages=10000

      # The maximum amount of time a message can sit in a log before we force a flush
      #log.flush.interval.ms=1000

      ############################# Log Retention Policy #############################

      # The following configurations control the disposal of log segments. The policy can
      # be set to delete segments after a period of time, or after a given size has accumulated.
      # A segment will be deleted whenever *either* of these criteria are met. Deletion always happens
      # from the end of the log.

      # https://cwiki.apache.org/confluence/display/KAFKA/KIP-186%3A+Increase+offsets+retention+default+to+7+days
      offsets.retention.minutes=10080

      # The minimum age of a log file to be eligible for deletion due to age
      log.retention.hours=-1

      # A size-based retention policy for logs. Segments are pruned from the log unless the remaining
      # segments drop below log.retention.bytes. Functions independently of log.retention.hours.
      #log.retention.bytes=1073741824

      # The maximum size of a log segment file. When this size is reached a new log segment will be created.
      #log.segment.bytes=1073741824

      # The interval at which log segments are checked to see if they can be deleted according
      # to the retention policies
      #log.retention.check.interval.ms=300000

      ############################# Zookeeper #############################

      # Zookeeper connection string (see zookeeper docs for details).
      # This is a comma separated host:port pairs, each corresponding to a zk
      # server. e.g. "127.0.0.1:3000,127.0.0.1:3001,127.0.0.1:3002".
      # You can also append an optional chroot string to the urls to specify the
      # root directory for all kafka znodes.
      zookeeper.connect=zookeeper.openshift-logging.svc.cluster.local:2181

      # Timeout in ms for connecting to zookeeper
      #zookeeper.connection.timeout.ms=6000


      ############################# Group Coordinator Settings #############################

      # The following configuration specifies the time, in milliseconds, that the GroupCoordinator will delay the initial consumer rebalance.
      # The rebalance will be further delayed by the value of group.initial.rebalance.delay.ms as new members join the group, up to a maximum of max.poll.interval.ms.
      # The default value for this is 3 seconds.
      # We override this to 0 here as it makes for a better out-of-the-box experience for development and testing.
      # However, in production environments the default value of 3 seconds is more suitable as this will help to avoid unnecessary, and potentially expensive, rebalances during application startup.
      #group.initial.rebalance.delay.ms=0
    `

	log4jProperties = `
      # Unspecified loggers and loggers with additivity=true output to server.log and stdout
      # Note that INFO only applies to unspecified loggers, the log level of the child logger is used otherwise
      log4j.rootLogger=INFO, stdout

      log4j.appender.stdout=org.apache.log4j.ConsoleAppender
      log4j.appender.stdout.layout=org.apache.log4j.PatternLayout
      log4j.appender.stdout.layout.ConversionPattern=[%d] %p %m (%c)%n

      log4j.appender.kafkaAppender=org.apache.log4j.DailyRollingFileAppender
      log4j.appender.kafkaAppender.DatePattern='.'yyyy-MM-dd-HH
      log4j.appender.kafkaAppender.File=${kafka.logs.dir}/server.log
      log4j.appender.kafkaAppender.layout=org.apache.log4j.PatternLayout
      log4j.appender.kafkaAppender.layout.ConversionPattern=[%d] %p %m (%c)%n

      log4j.appender.stateChangeAppender=org.apache.log4j.DailyRollingFileAppender
      log4j.appender.stateChangeAppender.DatePattern='.'yyyy-MM-dd-HH
      log4j.appender.stateChangeAppender.File=${kafka.logs.dir}/state-change.log
      log4j.appender.stateChangeAppender.layout=org.apache.log4j.PatternLayout
      log4j.appender.stateChangeAppender.layout.ConversionPattern=[%d] %p %m (%c)%n

      log4j.appender.requestAppender=org.apache.log4j.DailyRollingFileAppender
      log4j.appender.requestAppender.DatePattern='.'yyyy-MM-dd-HH
      log4j.appender.requestAppender.File=${kafka.logs.dir}/kafka-request.log
      log4j.appender.requestAppender.layout=org.apache.log4j.PatternLayout
      log4j.appender.requestAppender.layout.ConversionPattern=[%d] %p %m (%c)%n

      log4j.appender.cleanerAppender=org.apache.log4j.DailyRollingFileAppender
      log4j.appender.cleanerAppender.DatePattern='.'yyyy-MM-dd-HH
      log4j.appender.cleanerAppender.File=${kafka.logs.dir}/log-cleaner.log
      log4j.appender.cleanerAppender.layout=org.apache.log4j.PatternLayout
      log4j.appender.cleanerAppender.layout.ConversionPattern=[%d] %p %m (%c)%n

      log4j.appender.controllerAppender=org.apache.log4j.DailyRollingFileAppender
      log4j.appender.controllerAppender.DatePattern='.'yyyy-MM-dd-HH
      log4j.appender.controllerAppender.File=${kafka.logs.dir}/controller.log
      log4j.appender.controllerAppender.layout=org.apache.log4j.PatternLayout
      log4j.appender.controllerAppender.layout.ConversionPattern=[%d] %p %m (%c)%n

      log4j.appender.authorizerAppender=org.apache.log4j.DailyRollingFileAppender
      log4j.appender.authorizerAppender.DatePattern='.'yyyy-MM-dd-HH
      log4j.appender.authorizerAppender.File=${kafka.logs.dir}/kafka-authorizer.log
      log4j.appender.authorizerAppender.layout=org.apache.log4j.PatternLayout
      log4j.appender.authorizerAppender.layout.ConversionPattern=[%d] %p %m (%c)%n

      # Change the two lines below to adjust ZK client logging
      log4j.logger.org.I0Itec.zkclient.ZkClient=INFO
      log4j.logger.org.apache.zookeeper=INFO

      # Change the two lines below to adjust the general broker logging level (output to server.log and stdout)
      log4j.logger.kafka=INFO
      log4j.logger.org.apache.kafka=INFO

      # Change to DEBUG or TRACE to enable request logging
      log4j.logger.kafka.request.logger=WARN, requestAppender
      log4j.additivity.kafka.request.logger=false

      # Uncomment the lines below and change log4j.logger.kafka.network.RequestChannel$ to TRACE for additional output
      # related to the handling of requests
      #log4j.logger.kafka.network.Processor=TRACE, requestAppender
      #log4j.logger.kafka.server.KafkaApis=TRACE, requestAppender
      #log4j.additivity.kafka.server.KafkaApis=false
      log4j.logger.kafka.network.RequestChannel$=WARN, requestAppender
      log4j.additivity.kafka.network.RequestChannel$=false

      log4j.logger.kafka.controller=TRACE, controllerAppender
      log4j.additivity.kafka.controller=false

      log4j.logger.kafka.log.LogCleaner=INFO, cleanerAppender
      log4j.additivity.kafka.log.LogCleaner=false

      log4j.logger.state.change.logger=TRACE, stateChangeAppender
      log4j.additivity.state.change.logger=false

      # Change to DEBUG to enable audit log for the authorizer
      log4j.logger.kafka.authorizer.logger=WARN, authorizerAppender
      log4j.additivity.kafka.authorizer.logger=false
    `

	initZookeeperScript = `
      #!/bin/bash
      set -e

      [ -d /var/lib/zookeeper/data ] || mkdir /var/lib/zookeeper/data
      [ -z "$ID_OFFSET" ] && ID_OFFSET=1
      export ZOOKEEPER_SERVER_ID=$((${HOSTNAME##*-} + $ID_OFFSET))
      echo "${ZOOKEEPER_SERVER_ID:-1}" | tee /var/lib/zookeeper/data/myid
      cp -Lur /etc/kafka-configmap/* /etc/kafka/
    `

	zookeeperProperties = `
      4lw.commands.whitelist=ruok
      tickTime=2000
      dataDir=/var/lib/zookeeper/data
      dataLogDir=/var/lib/zookeeper/log
      clientPort=2181
    `
	zookeeperLog4JProperties = `
      log4j.rootLogger=INFO, stdout
      log4j.appender.stdout=org.apache.log4j.ConsoleAppender
      log4j.appender.stdout.layout=org.apache.log4j.PatternLayout
      log4j.appender.stdout.layout.ConversionPattern=[%d] %p %m (%c)%n

      # Suppress connection log messages, three lines per livenessProbe execution
      log4j.logger.org.apache.zookeeper.server.NIOServerCnxnFactory=WARN
      log4j.logger.org.apache.zookeeper.server.NIOServerCnxn=WARN
    `
	functionalPodinitKafkaScript = `
      #!/bin/bash
      set -e
      cp /etc/kafka-configmap/log4j.properties /etc/kafka/

      KAFKA_BROKER_ID=${HOSTNAME##*-}
      #SEDS=("s/#init#broker.id=#init#/broker.id=$KAFKA_BROKER_ID/")
      LABELS="kafka-broker-id=$KAFKA_BROKER_ID"
      ANNOTATIONS=""

    `

	functionalPodclientProperties = `
      security.protocol=SSL
      ssl.truststore.location=/etc/kafka-certs/ca-bundle.jks
      ssl.truststore.type=JKS
      ssl.truststore.password=ca-bundle
    `

	functionalPodserverProperties = `
      ############################# Log Basics #############################

      # A comma separated list of directories under which to store log files
      # Overrides log.dir
      log.dirs=/var/lib/kafka/data/topics

      # The default number of log partitions per topic. More partitions allow greater
      # parallelism for consumption, but this will also result in more files across
      # the brokers.
      num.partitions=1

      default.replication.factor=1

      min.insync.replicas=1

      auto.create.topics.enable=false

      # Max Messages in Bytes set to 10M > fluentd buffer chunk_limit_size config
      message.max.bytes=10000000

      # The number of threads per data directory to be used for log recovery at startup and flushing at shutdown.
      # This value is recommended to be increased for installations with data dirs located in RAID array.
      #num.recovery.threads.per.data.dir=1

      ############################# Server Basics #############################

      # The id of the broker. This must be set to a unique integer for each broker.
      broker.id=0

      #init#broker.rack=#init#

      ############################# Socket Server Settings #############################

      # The address the socket server listens on. It will get the value returned from
      # java.net.InetAddress.getCanonicalHostName() if not configured.
      #   FORMAT:
      #     listeners = listener_name://host_name:port
      #   EXAMPLE:
      #     listeners = PLAINTEXT://your.host.name:9092

      listeners=PLAINTEXT://:9092,SSL://:9093
      ssl.keystore.type=JKS
      ssl.keystore.location=/etc/kafka-certs/server.jks
      ssl.keystore.password=server

      # Hostname and port the broker will advertise to producers and consumers. If not set,
      # it uses the value for "listeners" if configured.  Otherwise, it will use the value
      # returned from java.net.InetAddress.getCanonicalHostName().
      advertised.listeners=PLAINTEXT://localhost:9092,SSL://localhost:9093
      #init#advertised.listeners=PLAINTEXT://#init#

      # Maps listener names to security protocols, the default is for them to be the same. See the config documentation for more details
      #listener.security.protocol.map=PLAINTEXT:PLAINTEXT,SSL:SSL,SASL_PLAINTEXT:SASL_PLAINTEXT,SASL_SSL:SASL_SSL
      listener.security.protocol.map=PLAINTEXT:PLAINTEXT,SSL:SSL,SASL_PLAINTEXT:SASL_PLAINTEXT,SASL_SSL:SASL_SSL,OUTSIDE:PLAINTEXT
      inter.broker.listener.name=PLAINTEXT
      

      # The number of threads that the server uses for receiving requests from the network and sending responses to the network
      #num.network.threads=3

      # The number of threads that the server uses for processing requests, which may include disk I/O
      #num.io.threads=8

      # The send buffer (SO_SNDBUF) used by the socket server
      #socket.send.buffer.bytes=102400

      # The receive buffer (SO_RCVBUF) used by the socket server
      #socket.receive.buffer.bytes=102400

      # The maximum size of a request that the socket server will accept (protection against OOM)
      #socket.request.max.bytes=104857600

      ############################# Internal Topic Settings  #############################
      # The replication factor for the group metadata internal topics "__consumer_offsets" and "__transaction_state"
      # For anything other than development testing, a value greater than 1 is recommended for to ensure availability such as 3.
      offsets.topic.replication.factor=1
      transaction.state.log.replication.factor=1
      transaction.state.log.min.isr=1

      ############################# Log Flush Policy #############################

      # Messages are immediately written to the filesystem but by default we only fsync() to sync
      # the OS cache lazily. The following configurations control the flush of data to disk.
      # There are a few important trade-offs here:
      #    1. Durability: Unflushed data may be lost if you are not using replication.
      #    2. Latency: Very large flush intervals may lead to latency spikes when the flush does occur as there will be a lot of data to flush.
      #    3. Throughput: The flush is generally the most expensive operation, and a small flush interval may lead to excessive seeks.
      # The settings below allow one to configure the flush policy to flush data after a period of time or
      # every N messages (or both). This can be done globally and overridden on a per-topic basis.

      # The number of messages to accept before forcing a flush of data to disk
      #log.flush.interval.messages=10000

      # The maximum amount of time a message can sit in a log before we force a flush
      #log.flush.interval.ms=1000

      ############################# Log Retention Policy #############################

      # The following configurations control the disposal of log segments. The policy can
      # be set to delete segments after a period of time, or after a given size has accumulated.
      # A segment will be deleted whenever *either* of these criteria are met. Deletion always happens
      # from the end of the log.

      # https://cwiki.apache.org/confluence/display/KAFKA/KIP-186%3A+Increase+offsets+retention+default+to+7+days
      offsets.retention.minutes=10080

      # The minimum age of a log file to be eligible for deletion due to age
      log.retention.hours=-1

      # A size-based retention policy for logs. Segments are pruned from the log unless the remaining
      # segments drop below log.retention.bytes. Functions independently of log.retention.hours.
      #log.retention.bytes=1073741824

      # The maximum size of a log segment file. When this size is reached a new log segment will be created.
      #log.segment.bytes=1073741824

      # The interval at which log segments are checked to see if they can be deleted according
      # to the retention policies
      #log.retention.check.interval.ms=300000

      ############################# Zookeeper #############################

      # Zookeeper connection string (see zookeeper docs for details).
      # This is a comma separated host:port pairs, each corresponding to a zk
      # server. e.g. "127.0.0.1:3000,127.0.0.1:3001,127.0.0.1:3002".
      # You can also append an optional chroot string to the urls to specify the
      # root directory for all kafka znodes.
      zookeeper.connect=localhost:2181

      # Timeout in ms for connecting to zookeeper
      zookeeper.connection.timeout.ms=60000


      ############################# Group Coordinator Settings #############################

      # The following configuration specifies the time, in milliseconds, that the GroupCoordinator will delay the initial consumer rebalance.
      # The rebalance will be further delayed by the value of group.initial.rebalance.delay.ms as new members join the group, up to a maximum of max.poll.interval.ms.
      # The default value for this is 3 seconds.
      # We override this to 0 here as it makes for a better out-of-the-box experience for development and testing.
      # However, in production environments the default value of 3 seconds is more suitable as this will help to avoid unnecessary, and potentially expensive, rebalances during application startup.
      #group.initial.rebalance.delay.ms=0
    `

	functionalPodlog4jProperties = `
      # Unspecified loggers and loggers with additivity=true output to server.log and stdout
      # Note that INFO only applies to unspecified loggers, the log level of the child logger is used otherwise
      log4j.rootLogger=INFO, stdout

      log4j.appender.stdout=org.apache.log4j.ConsoleAppender
      log4j.appender.stdout.layout=org.apache.log4j.PatternLayout
      log4j.appender.stdout.layout.ConversionPattern=[%d] %p %m (%c)%n

      log4j.appender.kafkaAppender=org.apache.log4j.DailyRollingFileAppender
      log4j.appender.kafkaAppender.DatePattern='.'yyyy-MM-dd-HH
      log4j.appender.kafkaAppender.File=${kafka.logs.dir}/server.log
      log4j.appender.kafkaAppender.layout=org.apache.log4j.PatternLayout
      log4j.appender.kafkaAppender.layout.ConversionPattern=[%d] %p %m (%c)%n

      log4j.appender.stateChangeAppender=org.apache.log4j.DailyRollingFileAppender
      log4j.appender.stateChangeAppender.DatePattern='.'yyyy-MM-dd-HH
      log4j.appender.stateChangeAppender.File=${kafka.logs.dir}/state-change.log
      log4j.appender.stateChangeAppender.layout=org.apache.log4j.PatternLayout
      log4j.appender.stateChangeAppender.layout.ConversionPattern=[%d] %p %m (%c)%n

      log4j.appender.requestAppender=org.apache.log4j.DailyRollingFileAppender
      log4j.appender.requestAppender.DatePattern='.'yyyy-MM-dd-HH
      log4j.appender.requestAppender.File=${kafka.logs.dir}/kafka-request.log
      log4j.appender.requestAppender.layout=org.apache.log4j.PatternLayout
      log4j.appender.requestAppender.layout.ConversionPattern=[%d] %p %m (%c)%n

      log4j.appender.cleanerAppender=org.apache.log4j.DailyRollingFileAppender
      log4j.appender.cleanerAppender.DatePattern='.'yyyy-MM-dd-HH
      log4j.appender.cleanerAppender.File=${kafka.logs.dir}/log-cleaner.log
      log4j.appender.cleanerAppender.layout=org.apache.log4j.PatternLayout
      log4j.appender.cleanerAppender.layout.ConversionPattern=[%d] %p %m (%c)%n

      log4j.appender.controllerAppender=org.apache.log4j.DailyRollingFileAppender
      log4j.appender.controllerAppender.DatePattern='.'yyyy-MM-dd-HH
      log4j.appender.controllerAppender.File=${kafka.logs.dir}/controller.log
      log4j.appender.controllerAppender.layout=org.apache.log4j.PatternLayout
      log4j.appender.controllerAppender.layout.ConversionPattern=[%d] %p %m (%c)%n

      log4j.appender.authorizerAppender=org.apache.log4j.DailyRollingFileAppender
      log4j.appender.authorizerAppender.DatePattern='.'yyyy-MM-dd-HH
      log4j.appender.authorizerAppender.File=${kafka.logs.dir}/kafka-authorizer.log
      log4j.appender.authorizerAppender.layout=org.apache.log4j.PatternLayout
      log4j.appender.authorizerAppender.layout.ConversionPattern=[%d] %p %m (%c)%n

      # Change the two lines below to adjust ZK client logging
      log4j.logger.org.I0Itec.zkclient.ZkClient=INFO
      log4j.logger.org.apache.zookeeper=INFO

      # Change the two lines below to adjust the general broker logging level (output to server.log and stdout)
      log4j.logger.kafka=INFO
      log4j.logger.org.apache.kafka=INFO

      # Change to DEBUG or TRACE to enable request logging
      log4j.logger.kafka.request.logger=WARN, requestAppender
      log4j.additivity.kafka.request.logger=false

      # Uncomment the lines below and change log4j.logger.kafka.network.RequestChannel$ to TRACE for additional output
      # related to the handling of requests
      #log4j.logger.kafka.network.Processor=TRACE, requestAppender
      #log4j.logger.kafka.server.KafkaApis=TRACE, requestAppender
      #log4j.additivity.kafka.server.KafkaApis=false
      log4j.logger.kafka.network.RequestChannel$=WARN, requestAppender
      log4j.additivity.kafka.network.RequestChannel$=false

      log4j.logger.kafka.controller=TRACE, controllerAppender
      log4j.additivity.kafka.controller=false

      log4j.logger.kafka.log.LogCleaner=INFO, cleanerAppender
      log4j.additivity.kafka.log.LogCleaner=false

      log4j.logger.state.change.logger=TRACE, stateChangeAppender
      log4j.additivity.state.change.logger=false

      # Change to DEBUG to enable audit log for the authorizer
      log4j.logger.kafka.authorizer.logger=WARN, authorizerAppender
      log4j.additivity.kafka.authorizer.logger=false
    `

	functionalPodinitZookeeperScript = `
      #!/bin/bash
      set -e

      [ -d /var/lib/zookeeper/data ] || mkdir /var/lib/zookeeper/data
      [ -z "$ID_OFFSET" ] && ID_OFFSET=1
      export ZOOKEEPER_SERVER_ID=$((${HOSTNAME##*-} + $ID_OFFSET))
      echo "${ZOOKEEPER_SERVER_ID:-1}" | tee /var/lib/zookeeper/data/myid
      cp -Lur /etc/kafka-configmap/* /etc/kafka/
    `

	functionalPodzookeeperProperties = `
      4lw.commands.whitelist=ruok
      tickTime=2000
      dataDir=/var/lib/zookeeper/data
      dataLogDir=/var/lib/zookeeper/log
      clientPort=2181
    `
	functionalPodzookeeperLog4JProperties = `
      log4j.rootLogger=INFO, stdout
      log4j.appender.stdout=org.apache.log4j.ConsoleAppender
      log4j.appender.stdout.layout=org.apache.log4j.PatternLayout
      log4j.appender.stdout.layout.ConversionPattern=[%d] %p %m (%c)%n

      # Suppress connection log messages, three lines per livenessProbe execution
      log4j.logger.org.apache.zookeeper.server.NIOServerCnxnFactory=WARN
      log4j.logger.org.apache.zookeeper.server.NIOServerCnxn=WARN
    `

	FluentKafkaConf = `
## CLO GENERATED CONFIGURATION ###
# This file is a copy of the fluentd configuration entrypoint
# which should normally be supplied in a configmap.

<system>
  log_level "#{ENV['LOG_LEVEL'] || 'warn'}"
</system>

# Prometheus Monitoring
<source>
  @type prometheus
  bind "#{ENV['POD_IP']}"
  <ssl>
    enable true
    certificate_path "#{ENV['METRICS_CERT'] || '/etc/fluent/metrics/tls.crt'}"
    private_key_path "#{ENV['METRICS_KEY'] || '/etc/fluent/metrics/tls.key'}"
  </ssl>
</source>

<source>
  @type prometheus_monitor
  <labels>
    hostname ${hostname}
  </labels>
</source>

# excluding prometheus_tail_monitor
# since it leaks namespace/pod info
# via file paths

# tail_monitor plugin which publishes log_collected_bytes_total
<source>
  @type collected_tail_monitor
  <labels>
    hostname ${hostname}
  </labels>
</source>

# This is considered experimental by the repo
<source>
  @type prometheus_output_monitor
  <labels>
    hostname ${hostname}
  </labels>
</source>

# Logs from containers (including openshift containers)
<source>
  @type tail
  @id container-input
  path "/var/log/containers/*.log"
  exclude_path ["/var/log/containers/fluentd-*_openshift-logging_*.log", "/var/log/containers/elasticsearch-*_openshift-logging_*.log", "/var/log/containers/kibana-*_openshift-logging_*.log"]
  pos_file "/var/lib/fluentd/pos/es-containers.log.pos"
  refresh_interval 5
  rotate_wait 5
  tag kubernetes.*
  read_from_head "true"
  skip_refresh_on_startup true
  @label @MEASURE
  <parse>
    @type multi_format
    <pattern>
      format json
      time_format '%Y-%m-%dT%H:%M:%S.%N%Z'
      keep_time_key true
    </pattern>
    <pattern>
      format regexp
      expression /^(?<time>[^\s]+) (?<stream>stdout|stderr)( (?<logtag>.))? (?<log>.*)$/
      time_format '%Y-%m-%dT%H:%M:%S.%N%:z'
      keep_time_key true
    </pattern>
  </parse>
</source>

# Increment Prometheus metrics
<label @MEASURE>
  <filter **>
    @type record_transformer
    enable_ruby
    <record>
      msg_size ${record.to_s.length}
    </record>
  </filter>
  
  <filter **>
    @type prometheus
    <metric>
      name cluster_logging_collector_input_record_total
      type counter
      desc The total number of incoming records
      <labels>
        tag ${tag}
        hostname ${hostname}
      </labels>
    </metric>
  </filter>
  
  <filter **>
    @type prometheus
    <metric>
      name cluster_logging_collector_input_record_bytes
      type counter
      desc The total bytes of incoming records
      key msg_size
      <labels>
        tag ${tag}
        hostname ${hostname}
      </labels>
    </metric>
  </filter>
  
  <filter **>
    @type record_transformer
    remove_keys msg_size
  </filter>
  
  # Journal Logs go to INGRESS pipeline
  <match journal>
    @type relabel
    @label @INGRESS
  </match>
  
  # Audit Logs go to INGRESS pipeline
  <match *audit.log>
    @type relabel
    @label @INGRESS
  </match>
  
  # Kubernetes Logs go to CONCAT pipeline
  <match kubernetes.**>
    @type relabel
    @label @CONCAT
  </match>
</label>

# Concat log lines of container logs, and send to INGRESS pipeline
<label @CONCAT>
  <filter kubernetes.**>
    @type concat
    key log
    partial_key logtag
    partial_value P
    separator ''
  </filter>
  
  <match kubernetes.**>
    @type relabel
    @label @_MULITLINE_DETECT
  </match>
</label>

<label @_MULITLINE_DETECT>
  <match kubernetes.**>
    @id multiline-detect-except
    @type detect_exceptions
    remove_tag_prefix 'kubernetes'
    message log
    force_line_breaks true
    multiline_flush_interval .2
  </match>
  <match **>
    @type relabel
    @label @INGRESS
  </match>
</label>

# Ingress pipeline
<label @INGRESS>
  # Fix tag removed by multiline exception detection
  <match var.log.containers.**>
    @type rewrite_tag_filter
    <rule>
      key log
      pattern /.*/
      tag kubernetes.${tag}
    </rule>
  </match>
  
  # Filter out PRIORITY from journal logs
  <filter journal>
    @type grep
    <exclude>
      key PRIORITY
      pattern ^7$
    </exclude>
  </filter>
  
  # Process OVN logs
  <filter ovn-audit.log**>
    @type record_modifier
    <record>
      @timestamp ${DateTime.parse(record['message'].split('|')[0]).rfc3339(6)}
      level ${record['message'].split('|')[3].downcase}
    </record>
  </filter>
  
  # Retag Journal logs to specific tags
  <match journal>
    @type rewrite_tag_filter
    # skip to @INGRESS label section
    @label @INGRESS
  
    # see if this is a kibana container for special log handling
    # looks like this:
    # k8s_kibana.a67f366_logging-kibana-1-d90e3_logging_26c51a61-2835-11e6-ad29-fa163e4944d5_f0db49a2
    # we filter these logs through the kibana_transform.conf filter
    <rule>
      key CONTAINER_NAME
      pattern ^k8s_kibana\.
      tag kubernetes.journal.container.kibana
    </rule>
  
    <rule>
      key CONTAINER_NAME
      pattern ^k8s_[^_]+_logging-eventrouter-[^_]+_
      tag kubernetes.journal.container._default_.kubernetes-event
    </rule>
  
    # mark logs from default namespace for processing as k8s logs but stored as system logs
    <rule>
      key CONTAINER_NAME
      pattern ^k8s_[^_]+_[^_]+_default_
      tag kubernetes.journal.container._default_
    </rule>
  
    # mark logs from kube-* namespaces for processing as k8s logs but stored as system logs
    <rule>
      key CONTAINER_NAME
      pattern ^k8s_[^_]+_[^_]+_kube-(.+)_
      tag kubernetes.journal.container._kube-$1_
    </rule>
  
    # mark logs from openshift-* namespaces for processing as k8s logs but stored as system logs
    <rule>
      key CONTAINER_NAME
      pattern ^k8s_[^_]+_[^_]+_openshift-(.+)_
      tag kubernetes.journal.container._openshift-$1_
    </rule>
  
    # mark logs from openshift namespace for processing as k8s logs but stored as system logs
    <rule>
      key CONTAINER_NAME
      pattern ^k8s_[^_]+_[^_]+_openshift_
      tag kubernetes.journal.container._openshift_
    </rule>
  
    # mark fluentd container logs
    <rule>
      key CONTAINER_NAME
      pattern ^k8s_.*fluentd
      tag kubernetes.journal.container.fluentd
    </rule>
  
    # this is a kubernetes container
    <rule>
      key CONTAINER_NAME
      pattern ^k8s_
      tag kubernetes.journal.container
    </rule>
  
    # not kubernetes - assume a system log or system container log
    <rule>
      key _TRANSPORT
      pattern .+
      tag journal.system
    </rule>
  </match>
  
  # Invoke kubernetes apiserver to get kunbernetes metadata
  <filter kubernetes.**>
    @type kubernetes_metadata
    kubernetes_url 'https://kubernetes.default.svc'
    cache_size '1000'
    watch 'false'
    use_journal 'nil'
    ssl_partial_chain 'true'
  </filter>
  
  # Parse Json fields for container, journal and eventrouter logs
  <filter kubernetes.journal.**>
    @type parse_json_field
    merge_json_log 'false'
    preserve_json_log 'true'
    json_fields 'log,MESSAGE'
  </filter>
  
  <filter kubernetes.var.log.containers.**>
    @type parse_json_field
    merge_json_log 'false'
    preserve_json_log 'true'
    json_fields 'log,MESSAGE'
  </filter>
  
  <filter kubernetes.var.log.containers.eventrouter-** kubernetes.var.log.containers.cluster-logging-eventrouter-**>
    @type parse_json_field
    merge_json_log true
    preserve_json_log true
    json_fields 'log,MESSAGE'
  </filter>
  
  # Clean kibana log fields
  <filter **kibana**>
    @type record_transformer
    enable_ruby
    <record>
      log ${record['err'] || record['msg'] || record['MESSAGE'] || record['log']}
    </record>
    remove_keys req,res,msg,name,level,v,pid,err
  </filter>
  
  # Fix level field in audit logs
  <filter k8s-audit.log**>
    @type record_modifier
    <record>
      k8s_audit_level ${record['level']}
      level info
    </record>
  </filter>
  
  <filter openshift-audit.log**>
    @type record_modifier
    <record>
      openshift_audit_level ${record['level']}
      level info
    </record>
  </filter>
  
  # Viaq Data Model
  <filter **>
    @type viaq_data_model
    elasticsearch_index_prefix_field 'viaq_index_name'
    default_keep_fields CEE,time,@timestamp,aushape,ci_job,collectd,docker,fedora-ci,file,foreman,geoip,hostname,ipaddr4,ipaddr6,kubernetes,level,message,namespace_name,namespace_uuid,offset,openstack,ovirt,pid,pipeline_metadata,rsyslog,service,systemd,tags,testcase,tlog,viaq_msg_id
    extra_keep_fields ''
    keep_empty_fields 'message'
    use_undefined false
    undefined_name 'undefined'
    rename_time true
    rename_time_if_missing false
    src_time_name 'time'
    dest_time_name '@timestamp'
    pipeline_type 'collector'
    undefined_to_string 'false'
    undefined_dot_replace_char 'UNUSED'
    undefined_max_num_fields '-1'
    process_kubernetes_events 'false'
    <formatter>
      tag "system.var.log**"
      type sys_var_log
      remove_keys host,pid,ident
    </formatter>
    <formatter>
      tag "journal.system**"
      type sys_journal
      remove_keys log,stream,MESSAGE,_SOURCE_REALTIME_TIMESTAMP,__REALTIME_TIMESTAMP,CONTAINER_ID,CONTAINER_ID_FULL,CONTAINER_NAME,PRIORITY,_BOOT_ID,_CAP_EFFECTIVE,_CMDLINE,_COMM,_EXE,_GID,_HOSTNAME,_MACHINE_ID,_PID,_SELINUX_CONTEXT,_SYSTEMD_CGROUP,_SYSTEMD_SLICE,_SYSTEMD_UNIT,_TRANSPORT,_UID,_AUDIT_LOGINUID,_AUDIT_SESSION,_SYSTEMD_OWNER_UID,_SYSTEMD_SESSION,_SYSTEMD_USER_UNIT,CODE_FILE,CODE_FUNCTION,CODE_LINE,ERRNO,MESSAGE_ID,RESULT,UNIT,_KERNEL_DEVICE,_KERNEL_SUBSYSTEM,_UDEV_SYSNAME,_UDEV_DEVNODE,_UDEV_DEVLINK,SYSLOG_FACILITY,SYSLOG_IDENTIFIER,SYSLOG_PID
    </formatter>
    <formatter>
      tag "kubernetes.journal.container**"
      type k8s_journal
      remove_keys 'log,stream,MESSAGE,_SOURCE_REALTIME_TIMESTAMP,__REALTIME_TIMESTAMP,CONTAINER_ID,CONTAINER_ID_FULL,CONTAINER_NAME,PRIORITY,_BOOT_ID,_CAP_EFFECTIVE,_CMDLINE,_COMM,_EXE,_GID,_HOSTNAME,_MACHINE_ID,_PID,_SELINUX_CONTEXT,_SYSTEMD_CGROUP,_SYSTEMD_SLICE,_SYSTEMD_UNIT,_TRANSPORT,_UID,_AUDIT_LOGINUID,_AUDIT_SESSION,_SYSTEMD_OWNER_UID,_SYSTEMD_SESSION,_SYSTEMD_USER_UNIT,CODE_FILE,CODE_FUNCTION,CODE_LINE,ERRNO,MESSAGE_ID,RESULT,UNIT,_KERNEL_DEVICE,_KERNEL_SUBSYSTEM,_UDEV_SYSNAME,_UDEV_DEVNODE,_UDEV_DEVLINK,SYSLOG_FACILITY,SYSLOG_IDENTIFIER,SYSLOG_PID'
    </formatter>
    <formatter>
      tag "kubernetes.var.log.containers.eventrouter-** kubernetes.var.log.containers.cluster-logging-eventrouter-** k8s-audit.log** openshift-audit.log** ovn-audit.log**"
      type k8s_json_file
      remove_keys log,stream,CONTAINER_ID_FULL,CONTAINER_NAME
      process_kubernetes_events 'true'
    </formatter>
    <formatter>
      tag "kubernetes.var.log.containers**"
      type k8s_json_file
      remove_keys log,stream,CONTAINER_ID_FULL,CONTAINER_NAME
    </formatter>
    <elasticsearch_index_name>
      enabled 'true'
      tag "journal.system** system.var.log** **_default_** **_kube-*_** **_openshift-*_** **_openshift_**"
      name_type static
      static_index_name infra-write
    </elasticsearch_index_name>
    <elasticsearch_index_name>
      enabled 'true'
      tag "linux-audit.log** k8s-audit.log** openshift-audit.log** ovn-audit.log**"
      name_type static
      static_index_name audit-write
    </elasticsearch_index_name>
    <elasticsearch_index_name>
      enabled 'true'
      tag "**"
      name_type static
      static_index_name app-write
    </elasticsearch_index_name>
  </filter>
  
  # Generate elasticsearch id
  <filter **>
    @type elasticsearch_genid_ext
    hash_id_key viaq_msg_id
    alt_key kubernetes.event.metadata.uid
    alt_tags 'kubernetes.var.log.containers.logging-eventrouter-*.** kubernetes.var.log.containers.eventrouter-*.** kubernetes.var.log.containers.cluster-logging-eventrouter-*.** kubernetes.journal.container._default_.kubernetes-event'
  </filter>
  
  # Discard Infrastructure logs
  <match **_default_** **_kube-*_** **_openshift-*_** **_openshift_** journal.** system.var.log**>
    @type null
  </match>
  
  # Include Application logs
  <match kubernetes.**>
    @type relabel
    @label @_APPLICATION
  </match>
  
  # Discard Audit logs
  <match linux-audit.log** k8s-audit.log** openshift-audit.log** ovn-audit.log**>
    @type null
  </match>
  
  # Send any remaining unmatched tags to stdout
  <match **>
   @type stdout
  </match>
</label>

# Sending application source type to pipeline
<label @_APPLICATION>
  <filter **>
    @type record_modifier
    <record>
      log_type application
    </record>
  </filter>
  
  <match **>
    @type relabel
    @label @FORWARD_PIPELINE
  </match>
</label>

# Copying pipeline forward-pipeline to outputs
<label @FORWARD_PIPELINE>
  <match **>
    @type relabel
    @label @KAFKA
  </match>
</label>

# Ship logs to specific outputs
<label @KAFKA>
  <match **>
    @type kafka2
    @id kafka
    brokers localhost:9093
    default_topic clo-app-topic
    use_event_time true
    ssl_client_cert_key '/etc/kafka-certs/tls.key'
    ssl_client_cert '/etc/kafka-certs/tls.crt'
    ssl_ca_cert '/etc/kafka-certs/ca-bundle.crt'
    sasl_over_ssl false
    <format>
      @type json
    </format>
    <buffer clo-app-topic>
      @type file
      path '/var/lib/fluentd/kafka'
      flush_mode interval
      flush_interval 1s
      flush_thread_count 2
      retry_type exponential_backoff
      retry_wait 1s
      retry_max_interval 60s
      retry_timeout 60m
      queued_chunks_limit_size "#{ENV['BUFFER_QUEUE_LIMIT'] || '32'}"
      total_limit_size "#{ENV['TOTAL_LIMIT_SIZE_PER_BUFFER'] || '8589934592'}"
      chunk_limit_size "#{ENV['BUFFER_SIZE_LIMIT'] || '8m'}"
      overflow_action block
    </buffer>
  </match>
</label>
  `
)
