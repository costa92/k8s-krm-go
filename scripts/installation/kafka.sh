#!/bin/bash



# The root of the build/dist directory.
KRM_ROOT=$(dirname "${BASH_SOURCE[0]}")/../..
# If common.sh has already been sourced, it will not be sourced again here.
[[ -z ${COMMON_SOURCED} ]] && source ${KRM_ROOT}/scripts/installation/common.sh
# Set some environment variables.
KRM_KAFKA_HOST=${KRM_KAFKA_HOST:-127.0.0.1}
KRM_KAFKA_PORT=${KRM_KAFKA_PORT:-4317}

source "${KRM_ROOT}/scripts/lib/init.sh"

# Install kafka using containerization.
# Refer to https://www.baeldung.com/ops/kafka-docker-setup
krm::kafka::docker::install()
{
  krm::common::network
  docker run -d --name krm-zookeeper --network krm -p 2181:2181 -t wurstmeister/zookeeper
  docker run -d --name krm-kafka --link krm-zookeeper:zookeeper \
    --network krm \
    --restart=always \
    -v /etc/localtime:/etc/localtime \
    -p ${KRM_KAFKA_HOST}:${KRM_KAFKA_PORT}:9092 \
    --env KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 \
    --env KAFKA_ADVERTISED_HOST_NAME=${KRM_KAFKA_HOST} \
    --env KAFKA_ADVERTISED_PORT=${KRM_KAFKA_PORT} \
    wurstmeister/kafka

  echo "Sleeping to wait for onex-kafka container to complete startup ..."
  sleep 5
 krm::kafka::status || return 1
 krm::kafka::info
 krm::log::info "install kafka successfully"
}

# Uninstall the docker container.
krm::kafka::docker::uninstall()
{
  docker rm -f onex-zookeeper &>/dev/null
  docker rm -f onex-kafka &>/dev/null
 krm::log::info "uninstall kafka successfully"
}

# Install the kafka step by step.
# sbs is the abbreviation for "step by step".
# Refer to https://kafka.apache.org/documentation/#quickstart
krm::kafka::sbs::install()
{
 krm::kafka::docker::install
 krm::log::info "install kafka successfully"
}

# Uninstall the kafka step by step.
krm::kafka::sbs::uninstall()
{
 krm::kafka::docker::uninstall
 krm::log::info "uninstall kafka successfully"
}

# Print necessary information after docker or sbs installation.
krm::kafka::info()
{
  echo -e ${C_GREEN}kafka has been installed, here are some useful information:${C_NORMAL}
  cat << EOF | sed 's/^/  /'
Kafka brokers is: ${KRM_KAFKA_HOST}:${KRM_KAFKA_PORT}
EOF
}

# Status check after docker or sbs installation.
krm::kafka::status()
{
 krm::util::telnet ${KRM_KAFKA_HOST} ${KRM_KAFKA_PORT} || return 1
}

if [[ $* =~ krm::kafka:: ]]; then
  eval $*
fi
