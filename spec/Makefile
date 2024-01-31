PROVIDER_NAME=oktapam
OUTPUT_DIR=../oktapam/fwprovider
RESOURCE_NAME=resource_server_checkout_settings
GENERATOR_CONFIG_FILE=./generator_config.yaml
PROVIDER_CODE_SPEC_FILE=./provider-code-spec.json
OPENAPI_SPEC_FILE=./openapi.yaml
PACKAGE_NAME=fwprovider

scaffold-provider: ## Scaffold Provider Code (One time)
	tfplugingen-framework scaffold provider \
      --name ${PROVIDER_NAME} \
      --output-dir ${OUTPUT_DIR}

scaffold-resource: ## Scaffold Resource Code
	tfplugingen-framework scaffold resource \
    --name ${RESOURCE_NAME} \
    --output-dir ${OUTPUT_DIR} \
    --package ${PACKAGE_NAME}

generate-code-spec: ## Generate provider code spec
	tfplugingen-openapi generate \
      --config ${GENERATOR_CONFIG_FILE} \
      --output ${PROVIDER_CODE_SPEC_FILE} \
      ${OPENAPI_SPEC_FILE}

help:
	@echo "Usage: make [TARGET]"
	@echo "Targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-30s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

generate-resource-schema: ## Generate resource schema
	tfplugingen-framework generate resources \
	  --input ${PROVIDER_CODE_SPEC_FILE} \
	  --output ${OUTPUT_DIR}