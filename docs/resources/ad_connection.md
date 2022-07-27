---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "oktapam_ad_connection Resource - terraform-provider-oktapam"
subcategory: ""
description: |-
  An Active Directory (AD) Connection to query AD Domain for available servers. For more information check out the documentation https://help.okta.com/asa/en-us/Content/Topics/Adv_Server_Access/docs/ad-connections.htm on AD Connection creation and usage.
---

# oktapam_ad_connection (Resource)

An Active Directory (AD) Connection to query AD Domain for available servers. For more information check out the [documentation](https://help.okta.com/asa/en-us/Content/Topics/Adv_Server_Access/docs/ad-connections.htm) on AD Connection creation and usage.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `domain` (String) The domain against which to query.
- `gateway_id` (String) The UUID of the Gateway with which this AD Connection is associated.
- `name` (String) The human-readable name of the resource. Values are case-sensitive.
- `service_account_password` (String) The password of the service account that can be used to query the domain.
- `service_account_username` (String) The username of the service account that can be used to query the domain.

### Optional

- `certificate_id` (String) Certificate ID used for password less access method.
- `domain_controllers` (Set of String) A comma-separated list of the specific domain controller(s) that should be used to query the domain. Can be specified as a hostname or IP.
- `use_passwordless` (Boolean) if 'true', Users will not need password to login.

### Read-Only

- `deleted_at` (String) The UTC time of resource creation. Format is `2022-01-01 00:00:00 +0000 UTC`.
- `id` (String) The ID of this resource.

