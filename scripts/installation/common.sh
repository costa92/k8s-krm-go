#!/usr/bin/env bash


# 确保 onex 容器网络存在。
# 在 uninstall 时，可不删除 onex 容器网络，可以作为一个无害的无用数据
krm::common::network()
{
  docker network ls |grep -q krm || docker network create krm
}