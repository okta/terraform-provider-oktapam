---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "oktapam_database_password_settings Data Source - terraform-provider-oktapam"
subcategory: ""
description: |-
  Returns an existing Database Password Policy for a PAM Project.
---

# oktapam_database_password_settings (Data Source)

Returns an existing Database Password Policy for a PAM Project.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `project` (String) The UUID of a Project.
- `resource_group` (String) The UUID of a OktaPA Resource Group.

### Optional

- `character_options` (Block List, Max: 1) The specific characters rules required by the Password Policy. (see [below for nested schema](#nestedblock--character_options))
- `enable_periodic_rotation` (Boolean) If `true`, requires passwords to be rotated after a period of time has passed. You must also set the `periodic_rotation_duration_in_seconds` param.
- `max_length` (Number) The maximum length allowed for the password.
- `min_length` (Number) The minimum length allowed for the password.
- `periodic_rotation_duration_in_seconds` (Number) If `periodic_rotation` is enabled, specifies how often passwords are rotated.

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--character_options"></a>
### Nested Schema for `character_options`

Optional:

- `digits` (Boolean) If `true`, passwords can include one or more numeric characters.
- `lower_case` (Boolean) If `true`, passwords can include one or more lowercase characters.
- `punctuation` (Boolean) If `true`, passwords can include one or more punctuation/symbol characters.
- `require_from_each_set` (Boolean) If `true`, passwords must include at least one character from the selected sets.
- `upper_case` (Boolean) If `true`, passwords can include one or more uppercase characters.

