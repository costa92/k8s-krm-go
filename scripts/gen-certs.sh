#!/usr/bin/env bash

# The root of the build/dist directory
KRM_ROOT=$(dirname "${BASH_SOURCE[0]}")/..

source "${KRM_ROOT}/scripts/lib/init.sh"

# The output directory for the generated certificates
readonly LOCAL_OUTPUT_ROOT="${KRM_ROOT}/${OUT_DIR:-_output}"
readonly LOCAL_OUTPUT_CAPATH="${LOCAL_OUTPUT_ROOT}/cert"
readonly ONEX_DOMAIN="krmx.io"

CFSSL_BIN="/opt/homebrew/bin/cfssl"
CFSSLJSON_BIN="/opt/homebrew/bin/cfssljson"

# Hostname for the cert
#readonly CERT_HOSTNAME="${CERT_HOSTNAME:-onex-apiserver},127.0.0.1,localhost,"
readonly CERT_HOSTNAME="${CERT_HOSTNAME:-krm-apiserver},127.0.0.1,localhost,10.37.83.200"

function generate-node-cert() {
    local cert_dir=${1}  # 证书文件保存的目录
    local prefix=${2:-}  # 证书文件名的前缀
    local expiry=${3:-876000h} # 证书有效期

    mkdir -p "${cert_dir}"

    pushd "${cert_dir}" >/dev/null 2>&1  # 切换到证书文件保存的目录

#    krm::util::ensure-cfssl
  if [ ! -r "ca-config.json" ]; then
    cat >ca-config.json <<EOF
{
  "signing": {
    "default": {
      "expiry": "${expiry}"
    },
    "profiles": {
      "node": {
        "usages": [
          "signing",
          "key encipherment",
          "server auth",
          "client auth"
        ],
        "expiry": "${expiry}"
      }
  }
}
}
EOF
  fi
  if [ ! -r "ca-csr.json" ]; then
    cat >ca-csr.json <<EOF
{
  "CN": "onex",
  "key": {
    "algo": "rsa",
    "size": 2048
  },
  "names": [
    {
      "C": "CN",
      "ST": "Shenzhen",
      "L": "Shenzhen",
      "O": "onex",
      "OU": "System"
    }
  ],
  "ca": {
    "expiry": "${expiry}"
  }
}
EOF
  fi

  if [[ ! -r "ca.pem" || ! -r "ca-key.pem" ]]; then
    ${CFSSL_BIN} gencert -initca ca-csr.json | ${CFSSLJSON_BIN} -bare ca -
  fi

  if [[ -z "${prefix}" ]];then
    return 0
  fi

    #echo "Generate "${prefix}" certificates..."
    echo '{"CN":"'"${prefix}"'","hosts":[],"key":{"algo":"rsa","size":2048},"names":[{"C":"CN","ST":"Shenzhen","L":"Shenzhen","O":"tencent","OU":"'"${prefix}"'"}]}' \
      | ${CFSSL_BIN} gencert -hostname="${CERT_HOSTNAME},${prefix/-/.}.${ONEX_DOMAIN}" -ca=ca.pem -ca-key=ca-key.pem \
      -config=ca-config.json -profile=node - | ${CFSSLJSON_BIN} -bare "${prefix}"

    # the popd will access `directory stack`, no `real` parameters is actually needed
    # shellcheck disable=SC2119
     popd >/dev/null 2>&1
}

generate-node-cert "${LOCAL_OUTPUT_CAPATH}" "krm-apiserver" "876000h"