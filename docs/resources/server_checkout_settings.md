---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "oktapam_server_checkout_settings Resource - terraform-provider-oktapam"
subcategory: ""
description: |-
  The settings for limitting access to vaulted (shared) accounts for a single user and control the maximum amount of time they are allowed to access resources. For details, see Server Checkout https://help.okta.com/okta_help.htm?type=oie&id=csh-pam-checkout-configure.
---

# oktapam_server_checkout_settings (Resource)

The settings for limitting access to vaulted (shared) accounts for a single user and control the maximum amount of time they are allowed to access resources. For details, see [Server Checkout](https://help.okta.com/okta_help.htm?type=oie&id=csh-pam-checkout-configure).



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `checkout_duration_in_seconds` (Number) The duration in seconds for the checkout. If the checkout is enabled, the duration is the maximum time a user can access the resource before the checkout expires.
- `checkout_required` (Boolean) Indicates whether a checkout is mandatory for accessing resources within the project. If `true`, checkout is enforced for all applicable resources by default. If `false`, checkout is not required, and resources are accessible without it.
- `project` (String) The UUID of a Project.
- `resource_group` (String) The UUID of a OktaPA Resource Group.

### Optional

- `exclude_list` (List of String) If provided, only the account identifiers listed are excluded from the checkout requirement. This list is only considered if `checkout_required` is set to `true`. Only one of `include_list` and `exclude_list` can be specified in a request since they are mutually exclusive.
- `include_list` (List of String) If provided, only the account identifiers listed are required to perform a checkout to access the resource. This list is only considered if `checkout_required` is set to `true`. Only one of `include_list` and `exclude_list` can be specified in a request since they are mutually exclusive.

### Read-Only

- `id` (String) The ID of this resource.