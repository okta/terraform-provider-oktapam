## Code Generation Steps

### Scaffold OktaPAM Provider (Just need to do one time)
```shell
tfplugingen-framework scaffold provider \
  --name oktapam_framework_provider \
  --output-dir ./oktapam/fwprovider
```

### Scaffolding Resource
```shell
tfplugingen-framework scaffold resource \
--name resource_group \
--output-dir ./oktapam/fwprovider
```

### Generating Provider Code Spec
```shell
tfplugingen-openapi generate \
  --config ./spec/generator_config.yaml \
  --output ./spec/provider-code-spec.json \
  ./spec/openapi.yaml
```

### Framework Code generate command
```shell
tfplugingen-framework generate resources \
  --input ./spec/provider-code-spec.json \
  --output ./oktapam/fwprovider
```
