---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "oktapam_security_policy Data Source - terraform-provider-oktapam"
subcategory: ""
description: |-
  A policy which defines how users can gain access to resources. For details, see Security policy https://help.okta.com/okta_help.htm?type=oie&id=ext-pam-policy.
---

# oktapam_security_policy (Data Source)

A policy which defines how users can gain access to resources. For details, see [Security policy](https://help.okta.com/okta_help.htm?type=oie&id=ext-pam-policy).



<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `active` (Boolean) If true, indicates that the Security Policy is active.
- `description` (String) The human-readable description of the resource.
- `id` (String) The ID of this resource.
- `name` (String) The human-readable name of the resource. Values are case-sensitive.
- `principals` (List of Object) Defines the users bound to the Security Policy. (see [below for nested schema](#nestedatt--principals))
- `rule` (List of Object) Defines the privileges available to resources matched to the Security Policy. (see [below for nested schema](#nestedatt--rule))

<a id="nestedatt--principals"></a>
### Nested Schema for `principals`

Read-Only:

- `groups` (Set of String)


<a id="nestedatt--rule"></a>
### Nested Schema for `rule`

Read-Only:

- `conditions` (List of Object) (see [below for nested schema](#nestedobjatt--rule--conditions))
- `name` (String)
- `privileges` (List of Object) (see [below for nested schema](#nestedobjatt--rule--privileges))
- `resources` (List of Object) (see [below for nested schema](#nestedobjatt--rule--resources))

<a id="nestedobjatt--rule--conditions"></a>
### Nested Schema for `rule.conditions`

Read-Only:

- `access_request` (List of Object) (see [below for nested schema](#nestedobjatt--rule--conditions--access_request))
- `gateway` (List of Object) (see [below for nested schema](#nestedobjatt--rule--conditions--gateway))
- `mfa` (List of Object) (see [below for nested schema](#nestedobjatt--rule--conditions--mfa))

<a id="nestedobjatt--rule--conditions--access_request"></a>
### Nested Schema for `rule.conditions.access_request`

Read-Only:

- `expires_after_seconds` (Number)
- `request_type_id` (String)
- `request_type_name` (String)


<a id="nestedobjatt--rule--conditions--gateway"></a>
### Nested Schema for `rule.conditions.gateway`

Read-Only:

- `session_recording` (Boolean)
- `traffic_forwarding` (Boolean)


<a id="nestedobjatt--rule--conditions--mfa"></a>
### Nested Schema for `rule.conditions.mfa`

Read-Only:

- `acr_values` (String)
- `reauth_frequency_in_seconds` (Number)



<a id="nestedobjatt--rule--privileges"></a>
### Nested Schema for `rule.privileges`

Read-Only:

- `password_checkout_rdp` (List of Object) (see [below for nested schema](#nestedobjatt--rule--privileges--password_checkout_rdp))
- `password_checkout_ssh` (List of Object) (see [below for nested schema](#nestedobjatt--rule--privileges--password_checkout_ssh))
- `principal_account_rdp` (List of Object) (see [below for nested schema](#nestedobjatt--rule--privileges--principal_account_rdp))
- `principal_account_ssh` (List of Object) (see [below for nested schema](#nestedobjatt--rule--privileges--principal_account_ssh))
- `secret` (List of Object) (see [below for nested schema](#nestedobjatt--rule--privileges--secret))

<a id="nestedobjatt--rule--privileges--password_checkout_rdp"></a>
### Nested Schema for `rule.privileges.password_checkout_rdp`

Read-Only:

- `enabled` (Boolean)


<a id="nestedobjatt--rule--privileges--password_checkout_ssh"></a>
### Nested Schema for `rule.privileges.password_checkout_ssh`

Read-Only:

- `enabled` (Boolean)


<a id="nestedobjatt--rule--privileges--principal_account_rdp"></a>
### Nested Schema for `rule.privileges.principal_account_rdp`

Read-Only:

- `admin_level_permissions` (Boolean)
- `enabled` (Boolean)


<a id="nestedobjatt--rule--privileges--principal_account_ssh"></a>
### Nested Schema for `rule.privileges.principal_account_ssh`

Read-Only:

- `admin_level_permissions` (Boolean)
- `enabled` (Boolean)
- `sudo_command_bundles` (Set of String)
- `sudo_display_name` (String)


<a id="nestedobjatt--rule--privileges--secret"></a>
### Nested Schema for `rule.privileges.secret`

Read-Only:

- `folder_create` (Boolean)
- `folder_delete` (Boolean)
- `folder_update` (Boolean)
- `list` (Boolean)
- `secret_create` (Boolean)
- `secret_delete` (Boolean)
- `secret_reveal` (Boolean)
- `secret_update` (Boolean)



<a id="nestedobjatt--rule--resources"></a>
### Nested Schema for `rule.resources`

Read-Only:

- `secrets` (List of Object) (see [below for nested schema](#nestedobjatt--rule--resources--secrets))
- `servers` (List of Object) (see [below for nested schema](#nestedobjatt--rule--resources--servers))

<a id="nestedobjatt--rule--resources--secrets"></a>
### Nested Schema for `rule.resources.secrets`

Read-Only:

- `secret` (List of Object) (see [below for nested schema](#nestedobjatt--rule--resources--secrets--secret))
- `secret_folder` (List of Object) (see [below for nested schema](#nestedobjatt--rule--resources--secrets--secret_folder))

<a id="nestedobjatt--rule--resources--secrets--secret"></a>
### Nested Schema for `rule.resources.secrets.secret`

Read-Only:

- `secret_id` (String)


<a id="nestedobjatt--rule--resources--secrets--secret_folder"></a>
### Nested Schema for `rule.resources.secrets.secret_folder`

Read-Only:

- `secret_folder_id` (String)



<a id="nestedobjatt--rule--resources--servers"></a>
### Nested Schema for `rule.resources.servers`

Read-Only:

- `label_selectors` (List of Object) (see [below for nested schema](#nestedobjatt--rule--resources--servers--label_selectors))
- `server` (List of Object) (see [below for nested schema](#nestedobjatt--rule--resources--servers--server))
- `server_account` (List of Object) (see [below for nested schema](#nestedobjatt--rule--resources--servers--server_account))

<a id="nestedobjatt--rule--resources--servers--label_selectors"></a>
### Nested Schema for `rule.resources.servers.label_selectors`

Read-Only:

- `accounts` (List of String)
- `server_labels` (Map of String)


<a id="nestedobjatt--rule--resources--servers--server"></a>
### Nested Schema for `rule.resources.servers.server`

Read-Only:

- `server_id` (String)


<a id="nestedobjatt--rule--resources--servers--server_account"></a>
### Nested Schema for `rule.resources.servers.server_account`

Read-Only:

- `account` (String)
- `server_id` (String)
