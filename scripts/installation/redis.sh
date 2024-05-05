#!/usr/bin/env bash

KRM_ROOT=$(dirname "${BASH_SOURCE[0]}")/../..
KRM_REDIS_HOST=${KRM_REDIS_PORT:-127.0.0.1}
KRM_REDIS_PORT=${KRM_REDIS_PORT:-6379}
KRM_PASSWORD=${KRM_PASSWORD:-krm(#)666}
KRM_REDIS_DATA_DIR=${HOME}/Data

source "${KRM_ROOT}/scripts/lib/init.sh"
# 确保 onex 容器网络存在。
# 在 uninstall 时，可不删除 onex 容器网络，可以作为一个无害的无用数据
krm::common::network()
{
  docker network ls |grep -q krm || docker network create krm
}

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