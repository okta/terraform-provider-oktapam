# Terraform Provider for Okta's Privileged Access Management (Okta's PAM)

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 1.x
- [Go](https://golang.org/doc/install) 1.21+ (to build the provider plugin)

## Building The Provider

Clone repository to: `$PROJECT_DIR/terraform-provider-oktapam`

where `$PROJECT_DIR` is a directory within your local machine.

```sh
$ git clone git@github.com:okta/terraform-provider-oktapam $PROJECT_DIR/terraform-provider-oktapam
```

Go to the provider directory and build the provider

```sh
$ cd $PROJECT_DIR/terraform-provider-oktapam
$ make build
```

## Using the provider

1. You will need to create a service user
2. Keep a copy of the new service user key and secret, you will be prompted with those after creating the service user
3. Make sure the new service user is added to the appropriate PAM Groups and have the needed capabilities
4. Set the following environment variables prior to running:

```
export OKTAPAM_KEY=<service user key>
export OKTAPAM_SECRET=<service user secret>
export OKTAPAM_TEAM=<team name>
```

If you want to test against a non-production server, you need to set the URL via:

```
export OKTAPAM_API_HOST=<root url for host here>
```

If the non-production server is not within the `scaleft.com` or `okta.com` domains, you will need to set the
`OKTAPAM_TRUSTED_DOMAIN_OVERRIDE` environment variable to the full domain being used, e.g.:

```
export OKTAPAM_API_HOST="https://my.testing.domain"
export OKTAPAM_TRUSTED_DOMAIN_OVERRIDE="my.testing.domain"
```

## Security Policy Resources

We are in the process of transitioning security policy resources, so expect some churn. The existing
"oktapam_security_policy" resource will continue to exist for the foreseeable future. It will not, however,
be getting any new features. If it suits your needs as-is, continue to use it.

Its replacement is the intaptly-named "oktpam_security_policy_v2". This resource more closely aligns with
OPA's internal APIs and thus can be fleshed out to include all current and future functionality at the cost
of some ergonomics. We plan to continue to develop and iterate on this security policy resource.

Note: The "oktapam_security_policy_v2" resource is still under active development and breaking changes may
occur. Please report any issues you encounter.

## Developing the Provider

### Building Provider

To compile the provider, run `make build`. This will build the provider and put in the project directory

```sh
$ make build
```

To install the provider on your local machine, run `make install`. This will build the provider and install it within
the `~/.terraform.d/plugins/hashicorp.com/okta/oktapam/[VERSION]/[OS_ARCH]/terraform-provider-oktapam` directory.

```sh
$ make install
```

If you are using terraform 0.12.x, you will need to run `make link_legacy`. This will create a symlink in
`~/.terraform.d/plugins/` to the binary that is created by `make install`. You will only need to run `make link_legacy`
once.

```sh
$ make link_legacy
```

### Developing with Terraform Provider Framework

To add a new resource or data source using the new Terraform Provider Framework:

1. Create a new file in the `oktapam` directory with the name `resource_<resource_name>.go` or
   `data_source_<data_source_name>.go`.

2. Define a new struct that implements the `resource.Resource` or `resource.DataSource` interface.
   reference: [terraform provider framework resource](https://developer.hashicorp.com/terraform/plugin/framework/resources)

3. Implement the required methods:
    - `Metadata()`: Returns the resource type name.
    - `Schema()`: Defines the schema for the resource or data source.
    - `Create()`, `Read()`, `Update()`, and `Delete()` for resources.
    - `Read()` for data sources.

4. Register the new resource or data source in the `Resources`/`DataSources` method in `provider_pluginframework.go`.

5. Add any necessary helper functions or API client methods in `oktapam/client/`.

6. Create corresponding test files (`resource_<resource_name>_test.go` or `data_source_<data_source_name>_test.go`) and
   implement acceptance tests.

7. Create documentation in `docs/resources/<resource_name>.md` or `docs/data-sources/<data_source_name>.md`.
   Docs should be generated by [terraform-plugin-docs](https://github.com/hashicorp/terraform-plugin-docs) by running
   the following command:

  ```
  make generate
  ```

8. Run `go fmt ./...` to format your code if needed.

9. Run `make test` to ensure all tests pass or you can run specific tests by setting the `TESTARGS` environment
   variable. Check out
   below [section](https://github.com/okta/terraform-provider-oktapam/blob/9e410b3fd019e21467c9e22e93d00046556f171b/README.md#L149)
   for more details on running tests.

> Remember to follow the existing patterns and coding standards in the project when adding new resources or data
> sources. take a look at existing resources for examples.(resource_server_checkout_settings.go)

### Terraform Settings for Dev Environment

#### dev_overrides

While developing provider if you want to try a test configuration against development build of a provider then
dev_overrides
setting comes handy.

*Terraform init* command creates .terraform.lock.hcl file to store the version and checksums for the required providers.
Normally, Terraform verifies version and checksums to ensure operations are made with the intended version of a
provider.
These version and checksum rules are inconvenient while developing a provider, because every time you build a new binary
checksum
is going to be different and require developer to delete the exiting .terraform.lock.hcl file and re-initialize it via
init
command.

To avoid all this hassle, set the dev_overrides in terraform config file. By default, terraform look for .terraformrc
file
under ${HOME} directory. If you want to change it then set TF_CLI_CONFIG_FILE env variable.

```sh
export TF_CLI_CONFIG_FILE=/home/developer/tmp/dev.tfrc
```

Terraform config file content:

```
provider_installation {
  dev_overrides {
    "okta.com/pam/oktapam" = "<path to terraform-provider-oktapam locally built binary>" 
  }
  
  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```

After performing above steps, there is no need to run *terraform init* command. You can directly perform other
operations
like terraform plan/apply etc.

[Terraform Documentation](https://developer.hashicorp.com/terraform/cli/config/config-file#development-overrides-for-provider-developers)
for more details.

**Note:** If the test configuration file has references to multiple providers and for some of them you don't want
development
overrides then this may not be a preferred approach. Please
refer [githhub open issue](https://github.com/hashicorp/terraform/issues/27459).

### Running Tests

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

We have two sets of acceptance tests, one for an Advanced Server Access team, and one for an Okta Privileged Access
team. There are some tests that are shared between the two suites.

*Note:* Acceptance tests create real resources, and often cost money to run. If you wish to test against a dev/test
server, ensure that you have the `OKTAPAM_API_HOST` variable set.

Some tests rely on a server already being enrolled for your team. The ID of this server should be assigned to the
`TF_ACC_VALID_SERVER_ID` environment variable.

```
 export TF_ACC_VALID_SERVER_ID=<server id>
```

To run the full test suite of Acceptance tests for ASA, run `make testacc`.:

```sh
$ make testacc
```

To run the full test suite of Acceptance tests for Okta PA, run `make testaccpam`.:

```sh
$ make testaccpam
```

If you want to run specific acceptance tests then set TESTARGS variable. TestCaseFunctionName(t *testing.T) can be a
regular
expression too.

```sh
$ TESTARGS='-run TestcaseFunctionName' make testacc
```

Note that you'll need to sub `testacc` for `testaccpam` if the tests is an Okta PA test.

# Releasing the Provider

1. Bump version in `Makefile`.
2. Add last version to `tag-checks.yml`.
3. Merge and make tag corresponding to the new version.
4. Make release corresponding to new tag.

# Warnings

- In the `oktapam_project` resource the public key algorithm for certificate signing and validation can be set. By
  default, projects use the `ssh-ed25519` algorithm, but admins can configure the project to use the `ssh-rsa` to
  support legacy servers. `ssh-rsa` has been [deprecated by OpenSSH](https://www.openssh.com/txt/release-8.3) and should
  not be used, if possible.
