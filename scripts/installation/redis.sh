#!/usr/bin/env bash

KRM_REDIS_HOST=${KRM_REDIS_PORT:-127.0.0.1}
KRM_REDIS_PORT=${KRM_REDIS_PORT:-6379}
KRM_PASSWORD=${KRM_PASSWORD:-krm(#)666}
KRM_REDIS_DATA_DIR=${HOME}/Data

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


function krm::log::info() {
  local V="${V:-0}"
  if (( ONEX_VERBOSE < V )) || (( KUBE_VERBOSE < V )); then
    return
  fi

  for message; do
    echo "${message}"
  done
}

# Check if the input contains "krm::redis::"
if [[ "$*" =~ krm::redis:: ]]; then
  echo "Input contains 'krm::redis::'"
  # Execute the input as a command
  eval "$@"
fi