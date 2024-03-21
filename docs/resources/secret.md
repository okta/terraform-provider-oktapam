---
page_title: "oktapam_secret Resource - terraform-provider-oktapam"
description: |-
  Secrets are encrypted key-value pairs that can be used to store sensitive information, like usernames, passwords, API tokens, keys, or any string value. For details, see [Secrets](https://help.okta.com/okta_help.htm?type=oie&id=ext-pam-secrets)
---

# oktapam_secret (Resource)

Secrets are encrypted key-value pairs that can be used to store sensitive information, like usernames, passwords, API tokens, keys, or any string value. For details, see [Secrets](https://help.okta.com/okta_help.htm?type=oie&id=ext-pam-secrets)

## Security Notice
The secret specified in this resource will be stored *unencrypted* in your Terraform state and plan files.  Please take precautions to store these in a secure location.

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The human-readable name of the resource. Values are case-sensitive.
- `parent_folder` (String) The UUID of the directory which contains this Secret/Secret Folder element.
- `project` (String) The UUID of a Project.
- `resource_group` (String) The UUID of a OktaPA Resource Group.
- `secret` (Map of String, Sensitive) Defines the key value pairs that are used to store sensitive information, like usernames, passwords, API tokens, keys, or any string value.

### Optional

- `description` (String) The human-readable description of the resource.

### Read-Only

- `id` (String) The ID of this resource.