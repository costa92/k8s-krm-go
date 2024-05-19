
.PHONY: gen.protoc
gen.protoc: ## Generate go source files from protobuf files.
	@protoc \
		--proto_path=$(APIROOT) \
		--proto_path=$(APISROOT) \
		--proto_path=$(KRM_ROOT)/third_party \
		--go_out=paths=source_relative:$(APIROOT) \
		--go-http_out=paths=source_relative:$(APIROOT) \
		--go-grpc_out=paths=source_relative:$(APIROOT) \
		--go-errors_out=paths=source_relative:$(APIROOT) \
        --go-errors-code_out=paths=source_relative:$(KRM_ROOT)/docs/guide/zh-CN/api/errors-code \
		--validate_out=paths=source_relative,lang=go:$(APIROOT) \
		--openapi_out=fq_schema_naming=true,default_response=false:$(KRM_ROOT)/api/openapi \
		--openapiv2_out=$(KRM_ROOT)/api/openapi \
		--openapiv2_opt=logtostderr=true \
		--openapiv2_opt=json_names_for_fields=false \
		$(shell find $(APIROOT) -name *.proto)

#	@protoc \
#		--proto_path=$(APIROOT) \
#		--proto_path=$(APISROOT) \
#		--proto_path=$(KRM_ROOT)/third_party \
#		--grpc-gateway_out=paths=source_relative:$(APIROOT) \
#		$(shell find $(APIROOT)/fakeserver -name *.proto)

.PHONY: gen.ca.%
gen.ca.%: ## Generate CA files.
	$(eval CA := $(word 1,$(subst ., ,$*)))
	@echo "===========> Generating CA files for $(CA) $(OUTPUT_DIR)"
	@${SCRIPTS_DIR}/gen-certs.sh generate-node-cert $(OUTPUT_DIR)/cert $(CA)

# make gen.ca CERTIFICATES=usercente
.PHONY: gen.ca
gen.ca: $(addprefix gen.ca., $(CERTIFICATES)) ## Generate all CA files.
