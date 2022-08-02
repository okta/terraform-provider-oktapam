Terraform Provider for Okta's Privileged Access Management (Okta's PAM)
=========================

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.13.x
-	[Go](https://golang.org/doc/install) 1.17+ (to build the provider plugin)

Building The Provider
---------------------

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

Using the provider
----------------------
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

Developing the Provider
---------------------------
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

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.  If you wish to test against a dev/test server, ensure that you have the `OKTAPAM_API_HOST` variable set.

```sh
$ make testacc
```

Releasing the Provider
---------------------------
1. Bump version in `Makefile`.
2. Add last version to `tag-checks.yml`.
3. Merge and make tag corresponding to the new version.
4. Make release corresponding to new tag.