

# Specify tools category.
CODE_GENERATOR_TOOLS= client-gen lister-gen informer-gen defaulter-gen deepcopy-gen prerelease-lifecycle-gen conversion-gen openapi-gen
# code-generator is a makefile target not a real tool.
CI_WORKFLOW_TOOLS := code-generator golangci-lint goimports wire

.PHONY: tools.print-manual-tool
tools.print-manual-tool:
	@echo "===========> The following tools may need to be installed manually:"
	@echo $(MANUAL_INSTALL_TOOLS) | awk 'BEGIN{RS=" "} {printf("%15s%s\n","- ",$$0)}'

.PHONY: tools.install.%
tools.install.%: ## Install a specified tool.
	@echo "===========> Installing $*"
	@$(MAKE) _install.$*

.PHONY: _install.gentool
_install.gentool: ## Install gentool which is a tool used to generate gorm model and query code.
	@$(GO) install gorm.io/gen/tools/gentool@$(GEN_TOOL_VERSION)

.PHONY: _install.mockgen
_install.mockgen: ## Install mockgen which is a tool used to generate mock code.
	@$(GO) install github.com/golang/mock/mockgen@$(MOCKGEN_VERSION)

.PHONY: _install.cfssl
_install.cfssl: ## Install cfssl toolkit.
	@$(SCRIPTS_DIR)/install.sh krm::install::install_cfssl


.PHONY: _install.grpc
_install.grpc: ## Install grpc toolkit, includes multiple protoc plugins.
	@$(GO) install google.golang.org/protobuf/cmd/protoc-gen-go@$(PROTOC_GEN_GO_VERSION)
	@$(GO) install google.golang.org/grpc/cmd/protoc-gen-go-grpc@$(PROTOC_GEN_GO_GRPC_VERSION)
	@$(GO) install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@$(GRPC_GATEWAY_VERSION)
	@$(GO) install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@$(GRPC_GATEWAY_VERSION)
	@$(SCRIPTS_DIR)/install-protoc.sh

.PHONY: _install.kratos
_install.kratos: _install.grpc ## Install kratos toolkit, includes multiple protoc plugins.
	@$(GO) install github.com/joelanford/go-apidiff@$(GO_APIDIFF_VERSION)
	@$(GO) install github.com/envoyproxy/protoc-gen-validate@$(PROTOC_GEN_VALIDATE_VERSION)
	@$(GO) install github.com/google/gnostic/cmd/protoc-gen-openapi@$(PROTOC_GEN_OPENAPI_VERSION)
	@$(GO) install github.com/go-kratos/kratos/cmd/kratos/v2@$(KRATOS_VERSION)
	@$(GO) install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@$(KRATOS_VERSION)
	@$(GO) install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@$(KRATOS_VERSION)
	@$(SCRIPTS_DIR)/add-completion.sh kratos bash

.PHONY: _install.grpcurl
_install.grpcurl:
	@$(GO) install github.com/fullstorydev/grpcurl/cmd/grpcurl@$(GRPCURL_VERSION)

.PHONY: _install.swagger
_install.swagger:
	@$(GO) install github.com/go-swagger/go-swagger/cmd/swagger@$(GO_SWAGGER_VERSION)


.PHONY: tools.verify.%
tools.verify.%: ## Verify a specified tool.
	@if ! which $* &>/dev/null; then $(MAKE) tools.install.$*; fi


.PHONY: tools.verify.code-generator
tools.verify.code-generator: $(addprefix _verify.code-generator., $(CODE_GENERATOR_TOOLS)) ## Verify a specified tool.

.PHONY: _verify.code-generator.%
_verify.code-generator.%:
	@if ! which $* &>/dev/null; then $(MAKE) tools.install.code-generator.$*; fi


# Code generator tools
.PHONY: _install.code-generator
_install.code-generator: $(addprefix tools.install.code-generator., $(CODE_GENERATOR_TOOLS)) ## Install all necessary code-generator tools.

.PHONY: _install.code-generator.%
_install.code-generator.%: ## Install specified code-generator tool.
	@$(GO) install k8s.io/code-generator/cmd/$*@$(CODE_GENERATOR_VERSION)
	#@$(GO) install github.com/colin404/code-generator/cmd/$*@$(CODE_GENERATOR_VERSION)
