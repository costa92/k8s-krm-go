#  ==============================
#  GoLang Makefile
#  ==============================
GO := go
GO_MINIMUM_VERSION ?= 1.22.2

GOPATH := $(shell go env GOPATH)
ifeq ($(origin GOBIN), undefined)
	GOBIN := $(GOPATH)/bin
endif

COMMANDS ?= $(filter-out %.md, $(wildcard ${KRM_ROOT}/cmd/*))
BINS ?= $(foreach cmd,${COMMANDS},$(notdir ${cmd}))

ifeq (${COMMANDS},)
  $(error Could not determine COMMANDS, set KRM_ROOT or run in source dir)
endif

ifeq (${BINS},)
  $(error Could not determine BINS, set ONEX_ROOT or run in source dir)
endif


# 确保每个目标后面都有一个以 ## 开头的解释。这是你修正后的 Makefile：
.PHONY: go.echo
go.echo: ## Echo GoLang environment variables.
	# 打印 go 的版本
	@echo "GOVERION: $(shell $(GO) version)"
	@echo "GOPATH: $(GOPATH)"
	@echo "GOBIN: $(GOBIN)"
	@echo "COMMANDS: $(COMMANDS)"
	@echo "KRM_ROOT: $(KRM_ROOT)"
	$(shell $(GO) version)":

.PHONY: go.version
go.version: ## Show GoLang version.
	@echo "GoLang version: $(shell go version)"


.PHONY: go.build.%
go.build.%:  ## Build specified applications with platform, os and arch.
	$(eval COMMAND := $(word 2,$(subst ., ,$*)))
	$(eval PLATFORM := $(word 1,$(subst ., ,$*)))
	$(eval OS := $(word 1,$(subst _, ,$(PLATFORM))))
	$(eval ARCH := $(word 2,$(subst _, ,$(PLATFORM))))
	@echo "{COMMAND}={COMMAND}"
	#@ONEX_GIT_VERSION=$(VERSION) $(SCRIPTS_DIR)/build.sh $(COMMAND) $(PLATFORM)
	@if grep -q "func main()" $(KRM_ROOT)/cmd/$(COMMAND)/*.go &>/dev/null; then \
		echo "===========> Building binary $(COMMAND) $(VERSION) for $(OS) $(ARCH)" ; \
		CGO_ENABLED=0 GOOS=$(OS) GOARCH=$(ARCH) $(GO) build $(GO_BUILD_FLAGS) \
		-o $(OUTPUT_DIR)/platforms/$(OS)/$(ARCH)/$(COMMAND)$(GO_OUT_EXT) $(PRJ_SRC_PATH)/cmd/$(COMMAND) ; \
		echo "===========> Done building binary BINS for $(BINS)" ; \
		echo "===========> Done building binary $(COMMAND) $(VERSION) for $(OS) $(ARCH)" ; \
	fi


.PHONY: go.build
go.build: $(addprefix go.build., $(addprefix $(PLATFORM)., $(BINS))) ## Build all applications.



.PHONY: go.build.multiarch
go.build.multiarch: $(foreach p,$(PLATFORMS),$(addprefix go.build., $(addprefix $(p)., $(BINS)))) ## Build all applications with all supported arch.


