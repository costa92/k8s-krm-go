#!/usr/bin/env bash

function krm::util::telnet()
{
  (
    set +o errexit
    set +o pipefail
    echo | telnet "$1" "$2" 2>&1|grep refused &>/dev/null
    if [ $? -eq 0 ]; then
      return 1
    fi
    return 0
  )
}