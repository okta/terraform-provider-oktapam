PROVIDER_NAME=oktapam
OUTPUT_DIR=../oktapam/fwprovider
RESOURCE_NAME=resource_group_clone
GENERATOR_CONFIG_FILE=./generator_config.yaml
PROVIDER_CODE_SPEC_FILE=./provider-code-spec.json
OPENAPI_SPEC_FILE=./openapi.yaml
PACKAGE_NAME=fwprovider

scaffold-provider:
	tfplugingen-framework scaffold provider \
      --name ${PROVIDER_NAME} \
      --output-dir ${OUTPUT_DIR}

scaffold-resource:
	tfplugingen-framework scaffold resource \
    --name ${RESOURCE_NAME} \
    --output-dir ${OUTPUT_DIR} \
    --package ${PACKAGE_NAME}

generate-code-spec:
	tfplugingen-openapi generate \
      --config ${GENERATOR_CONFIG_FILE} \
      --output ${PROVIDER_CODE_SPEC_FILE} \
      ${OPENAPI_SPEC_FILE}

generate-resource-schema:
	tfplugingen-framework generate resources \
	  --input ${PROVIDER_CODE_SPEC_FILE} \
	  --output ${OUTPUT_DIR}