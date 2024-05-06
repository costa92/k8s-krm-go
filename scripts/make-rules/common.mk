
# Define the root directory of the project

# ===================
# include the common make file
ifeq ($(origin KRM_ROOT), undefined)
KRM_ROOT := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
endif


SHELL := /usr/bin/env bash -o errexit -o pipefail -o nounset
.SHELLFLAGS = -ec

export SHELLOPTS := errexit


# ==============================
PRJ_SRC_PATH := github.com/costa92/k8s-krm-go

COMMA := ,
SPACE :=
SPACE +=

# ==============================
ifeq ($(origin OUTPUT_DIR), undefined)
OUTPUT_DIR := $(KRM_ROOT)/_output
$(shell mkdir -p $(OUTPUT_DIR))
endif

ifeq ($(origin BIN_DIR), undefined)
BIN_DIR := $(OUTPUT_DIR)/bin
$(shell mkdir -p $(BIN_DIR))
endif

ifeq ($(origin TOOLS_DIR), undefined)
TOOLS_DIR := $(OUTPUT_DIR)/tools
$(shell mkdir -p $(TOOLS_DIR))
endif

ifeq ($(origin TMP_DIR), undefined)
TMP_DIR := $(OUTPUT_DIR)/tmp
$(shell mkdir -p $(TMP_DIR))
endif

# set the version number. you should not need to do this
# for the majority of scenarios.
ifeq ($(origin VERSION), undefined)
# Current version of the project.
# 判断是否在git仓库中
  ifeq (, $(wildcard .git))
	VERSION := v0.0.0
  else
    VERSION := $(shell git describe --tags --always --match='v*')
    ifneq (,$(shell git status --porcelain 2>/dev/null))
      VERSION := $(VERSION)-dirty
    endif
  endif
endif


# Minimum test coverage
ifeq ($(origin COVERAGE),undefined)
COVERAGE := 60
endif

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
GOPATH ?= $(shell go env GOPATH)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# Set the default platform to the current system
PLATFORMS ?= linux_amd64 linux_arm64

# Set the default platform to the current system
ifeq ($(origin PLATFORM), undefined)
	ifeq ($(origin GOOS), undefined)
		GOOS := $(shell go env GOOS)
	endif
	ifeq ($(origin GOARCH), undefined)
		GOARCH := $(shell go env GOARCH)
	endif
	PLATFORM := $(GOOS)_$(GOARCH)
	# Use linux as the default OS when building images
	IMAGE_PLAT := linux_$(GOARCH)
else
	GOOS := $(word 1, $(subst _, ,$(PLATFORM)))
	GOARCH := $(word 2, $(subst _, ,$(PLATFORM)))
	IMAGE_PLAT := $(PLATFORM)
endif



# Specify components which need certificate
ifeq ($(origin CERTIFICATES),undefined)
CERTIFICATES=onex-apiserver admin
endif

MANIFESTS_DIR=$(KRM_ROOT)/manifests
SCRIPTS_DIR=$(KRM_ROOT)/scripts


APIROOT ?= $(KRM_ROOT)/pkg/api
APISROOT ?= $(KRM_ROOT)/pkg/apis
