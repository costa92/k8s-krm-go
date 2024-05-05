#!/usr/bin/env bash

# The root of the build/dist directory


# Log an error but keep going.  Don't dump the stack or exit.
function krm::log::error() {
  timestamp=$(date +"[%m%d %H:%M:%S]")
  echo "!!! ${timestamp} ${1-}" >&2
  shift
  for message; do
    echo "    ${message}" >&2
  done
}

function krm::log::info()
{
  local V="${V:-0}"
  if (( ONEX_VERBOSE < V )) || (( KUBE_VERBOSE < V )); then
    return
  fi

  for message; do
    echo "${message}"
  done
}
