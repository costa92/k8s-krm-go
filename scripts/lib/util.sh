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

function krm::util::host_arch() {
  local host_arch
  case "$(uname -m)" in
    x86_64*)
      host_arch=amd64
      ;;
    i?86_64*)
      host_arch=amd64
      ;;
    amd64*)
      host_arch=amd64
      ;;
    aarch64*)
      host_arch=arm64
      ;;
    arm64*)
      host_arch=arm64
      ;;
    arm*)
      host_arch=arm
      ;;
    i?86*)
      host_arch=x86
      ;;
    s390x*)
      host_arch=s390x
      ;;
    ppc64le*)
      host_arch=ppc64le
      ;;
    *)
      onex::log::error "Unsupported host arch. Must be x86_64, 386, arm, arm64, s390x or ppc64le."
      exit 1
      ;;
  esac
  echo "${host_arch}"
}

function krm::util::ensure-cfssl {
  # check if cfssl is installed
  if command -v cfssl &>/dev/null && command -v cfssljson &>/dev/null && command -v cfssl-certinfo &>/dev/null; then
    CFSSL_BIN=$(command -v cfssl)
    CFSSLJSON_BIN=$(command -v cfssljson)
    CFSSLCERTINFO_BIN=$(command -v cfssl-certinfo)
    return 0
  fi

    host_arch=$(krm::util::host_arch)
    krm::log::info "host_arch: ${host_arch}"
    if [[ "${host_arch}" == "arm64" ]]; then
    kernel=$(uname -s)
      if [[ "${kernel}" ==  "Darwin" ]]; then
          brew install cfssl
      else
          echo "Unknown, unsupported platform: ${kernel}." >&2
          echo "Supported platforms: Linux, Darwin." >&2
          exit 2
      fi
    elif [[ "${host_arch}" != "amd64" ]]; then
      echo "Cannot download cfssl on non-amd64 hosts and cfssl does not appear to be installed."
      echo "Please install cfssl, cfssljson and cfssl-certinfo and verify they are in \$PATH."
      echo "Hint: export PATH=\$PATH:\$GOPATH/bin; go get -u github.com/cloudflare/cfssl/cmd/..."
      exit 1
    else
        # Create a temp dir for cfssl if no directory was given
        local cfssldir=${1:-}
        if [[ -z "${cfssldir}" ]]; then
          cfssldir="$HOME/bin"
        fi

        mkdir -p "${cfssldir}"
        pushd "${cfssldir}" > /dev/null || return 1

        echo "Unable to successfully run 'cfssl' from ${PATH}; downloading instead..."
        kernel=$(uname -s)
        case "${kernel}" in
          Linux)
            curl --retry 10 -L -o cfssl https://pkg.cfssl.org/R1.2/cfssl_linux-amd64
            curl --retry 10 -L -o cfssljson https://pkg.cfssl.org/R1.2/cfssljson_linux-amd64
            curl --retry 10 -L -o cfssl-certinfo https://pkg.cfssl.org/R1.2/cfssl-certinfo_linux-amd64
            ;;
          Darwin)
            curl --retry 10 -L -o cfssl https://pkg.cfssl.org/R1.2/cfssl_darwin-amd64
            curl --retry 10 -L -o cfssljson https://pkg.cfssl.org/R1.2/cfssljson_darwin-amd64
            curl --retry 10 -L -o cfssl-certinfo https://pkg.cfssl.org/R1.2/cfssl-certinfo_darwin-amd64
            ;;
          *)
            echo "Unknown, unsupported platform: ${kernel}." >&2
            echo "Supported platforms: Linux, Darwin." >&2
            exit 2
        esac

        chmod +x cfssl || true
        chmod +x cfssljson || true
        chmod +x cfssl-certinfo || true

        CFSSL_BIN="${cfssldir}/cfssl"
        CFSSLJSON_BIN="${cfssldir}/cfssljson"
        CFSSLCERTINFO_BIN="${cfssldir}/cfssl-certinfo"
        if [[ ! -x ${CFSSL_BIN} || ! -x ${CFSSLJSON_BIN} || ! -x ${CFSSLCERTINFO_BIN} ]]; then
          echo "Failed to download 'cfssl'."
          echo "Please install cfssl, cfssljson and cfssl-certinfo and verify they are in \$PATH."
          echo "Hint: export PATH=\$PATH:\$GOPATH/bin; go get -u github.com/cloudflare/cfssl/cmd/..."
          exit 1
        fi
        popd > /dev/null || return 1
    fi
}