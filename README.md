Terraform Provider for Okta's Advanced Server Access (Okta's ASA)
=========================

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.12.x
-	[Go](https://golang.org/doc/install) 1.16+ (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$PROJECT_DIR/terraform-provider-oktaasa`

where `$PROJECT_DIR` is a directory within your local machine.

```sh
$ git clone git@github.com:okta/terraform-provider-oktaasa $PROJECT_DIR/terraform-provider-oktaasa
```

Go to the provider directory and build the provider

```sh
$ cd $PROJECT_DIR/terraform-provider-oktaasa
$ make build
```

Using the provider
----------------------
You will need to create a team a service user.  Then set the following environment variables prior to running: Okta's ASA API key, secret and team name. 

```
export OKTAASA_KEY_SECRET=<secret here>
export OKTAASA_KEY=<key here>
export OKTAASA_TEAM=<team name>
```

If you are able to test against a non-production server, you can set the URL via: 

```
export OKTAASA_API_HOST=<root url for host here>
```


Developing the Provider
---------------------------
To compile the provider, run `make build`. This will build the provider and put in the project directory

```sh
$ make bin
```

To install the provider on your local machine, run `make install`.  This will build the provider and install it within the `~/.terraform.d/plugins/hashicorp.com/okta/oktaasa/[VERSION]/[OS_ARCH]/terraform-provider-oktaasa` directory.

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

*Note:* Acceptance tests create real resources, and often cost money to run.  If you wish to test against a dev/test server, ensure that you have the `OKTAASA_API_HOST` variable set.

```sh
$ make testacc
```
