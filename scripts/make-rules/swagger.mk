# ==============================================================================
# Makefile helper functions for swagger
#

.PHONY: swagger.run # Generate swagger API docs
swagger.run: tools.verify.swagger
	@echo "===========> Generating swagger API docs"
	#@swagger generate spec --scan-models -w $(ONEX_ROOT)/cmd/gen-swagger-type-docs -o $(ONEX_ROOT)/api/swagger/kubernetes.yaml
	@swagger mixin `find $(KRM_ROOT)/api/openapi -name "*.swagger.json"` \
		-q                                                    \
		--keep-spec-order                                     \
		--format=yaml                                         \
		--ignore-conflicts                                    \
		-o $(KRM_ROOT)/api/swagger/swagger.yaml
	@echo "Generated at: $(KRM_ROOT)/api/swagger/swagger.yaml"

.PHONY: swagger.serve
swagger.serve: tools.verify.swagger
	@swagger serve -F=redoc --no-open --port 65534 $(KRM_ROOT)/api/swagger/swagger.yaml
