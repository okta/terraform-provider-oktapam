Quickstart 
---------------------------
To install the provider on your local machine, run `make install`.  This will build the provider and install it within the `~/.terraform.d/plugins/hashicorp.com/okta/oktapam/[VERSION]/[OS_ARCH]/terraform-provider-oktapam` directory.

```sh
$ make install
```

This will add the provider binary to `~/.terraform.d/plugins/` for local usage. 

If you are using terraform 0.12.x, you will need to run `make link_legacy`.  This will create a symlink in `~/.terraform.d/plugins/` to the binary that is created by `make install`.  You will only need to run `make link_legacy` once.

```sh
$ make link_legacy
````

Log into your team on `app.scaleft.com`, create a `service user`. Add it to a group with admin-level permissions. 
Go to the user details and create a new secret. Add the key, secret, and team to `terraform.tfvars` for `oktapam_key`, `oktapam_secret`, and `oktapam_team`, respectively.

You're now ready to roll! 

First check that terraform properly initializes: 
```sh 
$ export TF_LOG=trace
$ terraform init
```

Then check the plan: 
```sh 
$ terraform plan
```

And apply! 
```sh 
$ terraform apply
```
