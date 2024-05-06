#!/usr/bin/env bash

KRM_ROOT=$(dirname "${BASH_SOURCE[0]}")/../..
[[ -z ${COMMON_SOURCED} ]] && source ${KRM_ROOT}/scripts/installation/common.sh


KRM_REDIS_HOST=${KRM_REDIS_PORT:-127.0.0.1}
KRM_REDIS_PORT=${KRM_REDIS_PORT:-6379}
KRM_PASSWORD=${KRM_PASSWORD:-krm(#)666}
KRM_REDIS_DATA_DIR=${HOME}/Data
# If common.sh has already been sourced, it will not be sourced again here.
source "${KRM_ROOT}/scripts/lib/init.sh"


krm::redis::docker::install()
{
    krm::common::network
    docker run -d --name krm-redis \
      --network krm \
      -v ${KRM_REDIS_DATA_DIR}/redis:/data \
      -p ${KRM_REDIS_HOST}:${KRM_REDIS_PORT}:6379 \
      redis:7.2.3 \
      redis-server \
      --appendonly yes \
      --save 60 1 \
      --protected-mode no \
      --requirepass ${KRM_PASSWORD} \
      --loglevel debug

      krm::log::info "install redis successfully"
}


# Uninstall the docker container.
krm::redis::docker::uninstall()
{
  docker rm -f krm-redis &>/dev/null
  rm -rf ${KRM_REDIS_DATA_DIR}/redis
  krm::log::info "uninstall redis successfully"
}

# check the status of the redis
krm::redis::status()
{
  krm::log::info "start check status of redis"
  krm::util::telnet ${KRM_REDIS_HOST} ${KRM_REDIS_PORT} || return 1
  redis-cli --no-auth-warning -h ${KRM_REDIS_HOST} -p ${KRM_REDIS_PORT} -a "${KRM_PASSWORD}" --hotkeys || {
    krm::log::error "can not login with redis maybe not initialized properly."
    return 1
  }
}

# Check if the input contains "krm::redis::"
if [[ "$*" =~ krm::redis:: ]]; then
  echo "Input contains $*"
  # Execute the input as a command
  eval "$@"
fi