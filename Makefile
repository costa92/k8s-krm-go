# Build all by default, even if it's not first
.DEFAULT_GOAL := help
GO_MOD_NAME := "github.com/costa92/k8s-krm-go"
GO_MOD_DOMAIN := $(shell echo $(GO_MOD_NAME) | awk -F '/' '{print $$1}')


KRM_ROOT := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

# ==============================
include scripts/make-rules/common.mk
include scripts/make-rules/all.mk

#  ==============================
# Usage
define USAGE_OPTIONS

\033[35mOptions:\033[0m
	DBG              Whether to generate debug symbols. Default is 0.
	VERBOSE          Whether to output verbose logs. Default is 0.
endef
export USAGE_OPTIONS


##@ Test
.PHONY: test
test: ## Run tests
	@echo "Running tests..."
	@$(GO) test -v ./...

##@ Build
.PHONY: build
build: tidy ## Build the project.
	$(MAKE) go.build

.PHONY: tidy
tidy: ## Tidy the project.
	@$(GO) mod tidy

.PHONY: targets
targets: Makefile ## Show all Sub-makefile targets.
	@for mk in `echo $(MAKEFILE_LIST) | sed 's/Makefile //g'`; do echo -e \\n\\033[35m$$mk\\033[0m; awk -F':.*##' '/^[0-9A-Za-z._-]+:.*?##/ { printf "  \033[36m%-45s\033[0m %s\n", $$1, $$2 } /^\$$\([0-9A-Za-z_-]+\):.*?##/ { gsub("_","-", $$1); printf "  \033[36m%-45s\033[0m %s\n", tolower(substr($$1, 3, length($$1)-7)), $$2 }' $$mk;done;


.PHONY: swagger
#swagger: gen.protoc
swagger: ## Generate and aggregate swagger document.
	@$(MAKE) swagger.run

.PHONY: swagger.serve
serve-swagger: ## Serve swagger spec and docs at 65534.
	@$(MAKE) swagger.serve


##@ Help
.PHONY: help
help: Makefile ## Display this help info.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<TARGETS> <OPTIONS>\033[0m\n\n\033[35mTargets:\033[0m\n"} /^[0-9A-Za-z._-]+:.*?##/ { printf "  \033[36m%-45s\033[0m %s\n", $$1, $$2 } /^\$$\([0-9A-Za-z_-]+\):.*?##/ { gsub("_","-", $$1); printf "  \033[36m%-45s\033[0m %s\n", tolower(substr($$1, 3, length($$1)-7)), $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' Makefile #$(MAKEFILE_LIST)
	@echo -e "$$USAGE_OPTIONS"