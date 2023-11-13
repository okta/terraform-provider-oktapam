
### Scaffolding Resource
```shell
tfplugingen-framework scaffold resource \
--name resource_group \
--output-dir ./fwkProvider
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
  --output ./oktapam/fwkProvider 
```
