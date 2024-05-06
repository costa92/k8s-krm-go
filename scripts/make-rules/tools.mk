
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
_install.grpc:
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