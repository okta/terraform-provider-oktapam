# Terraform Provider for Okta's Privileged Access Management (Okta's PAM)

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) 0.13.x
-	[Go](https://golang.org/doc/install) 1.17+ (to build the provider plugin)

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

You will need to create a team a service user.  Then set the following environment variables prior to running: Okta's PAM API key, secret and team name. 

```
export OKTAPAM_SECRET=<secret here>
export OKTAPAM_KEY=<key here>
export OKTAPAM_TEAM=<team name>
```

If you are able to test against a non-production server, you can set the URL via: 

```
export OKTAPAM_API_HOST=<root url for host here>
```

If the non-production server is not within the `scaleft.com` or `okta.com` domains, you will need to set the `OKTAPAM_TRUSTED_DOMAIN_OVERRIDE` environment variable to the full domain being used, e.g.:

```
export OKTAPAM_API_HOST="https://my.testing.domain"
export OKTAPAM_TRUSTED_DOMAIN_OVERRIDE="my.testing.domain"
```

## Developing the Provider

### Building Provider
To compile the provider, run `make build`. This will build the provider and put in the project directory

```sh
$ make build
```

To install the provider on your local machine, run `make install`.  This will build the provider and install it within the `~/.terraform.d/plugins/hashicorp.com/okta/oktapam/[VERSION]/[OS_ARCH]/terraform-provider-oktapam` directory.

```sh
$ make install
```

If you are using terraform 0.12.x, you will need to run `make link_legacy`.  This will create a symlink in `~/.terraform.d/plugins/` to the binary that is created by `make install`.  You will only need to run `make link_legacy` once.  

```sh
$ make link_legacy
```

### Terraform Settings for Dev Environment

#### dev_overrides

While developing provider if you want to try a test configuration against development build of a provider then dev_overrides
setting comes handy. 

*Terraform init* command creates .terraform.lock.hcl file to store the version and checksums for the required providers.
Normally, Terraform verifies version and checksums to ensure operations are made with the intended version of a provider. 
These version and checksum rules are inconvenient while developing a provider, because every time you build a new binary checksum
is going to be different and require developer to delete the exiting .terraform.lock.hcl file and re-initialize it via init
command.

To avoid all this hassle, set the dev_overrides in terraform config file. By default, terraform look for .terraformrc file
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

After performing above steps, there is no need to run *terraform init* command. You can directly perform other operations
like terraform plan/apply etc.

[Terraform Documentation](https://developer.hashicorp.com/terraform/cli/config/config-file#development-overrides-for-provider-developers)
for more details.

**Note:** If the test configuration file has references to multiple providers and for some of them you don't want development
overrides then this may not be a preferred approach. Please refer [githhub open issue](https://github.com/hashicorp/terraform/issues/27459).

### Running Tests

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.  If you wish to test against a dev/test server, ensure that you have the `OKTAPAM_API_HOST` variable set.

```sh
$ make testacc
```

If you want to run specific acceptance tests then set TESTARGS variable. TestCaseFunctionName(t *testing.T) can be a regular 
expression too.

```sh
$ TESTARGS='-run TestcaseFunctionName' make testacc
```

# Releasing the Provider

1. Bump version in `Makefile`.
2. Add last version to `tag-checks.yml`.
3. Merge and make tag corresponding to the new version.
4. Make release corresponding to new tag.

# Warnings

- In the `oktapam_project` resource the public key algorithm for certificate signing and validation can be set. By default, projects use the `ssh-ed25519` algorithm, but admins can configure the project to use the `ssh-rsa` to support legacy servers. `ssh-rsa` has been [deprecated by OpenSSH](https://www.openssh.com/txt/release-8.3) and should not be used, if possible.