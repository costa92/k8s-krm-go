
.PHONY: _install.gentool
_install.gentool: ## Install gentool which is a tool used to generate gorm model and query code.
	@$(GO) install gorm.io/gen/tools/gentool@$(GEN_TOOL_VERSION)

.PHONY: _install.mockgen
_install.mockgen: ## Install mockgen which is a tool used to generate mock code.
	@$(GO) install github.com/golang/mock/mockgen@$(MOCKGEN_VERSION)

.PHONY: _install.cfssl
_install.cfssl: ## Install cfssl toolkit.
	@$(SCRIPTS_DIR)/install.sh krm::install::install_cfssl