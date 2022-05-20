---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "oktapam_project_group Data Source - terraform-provider-oktapam"
subcategory: ""
description: |-
  
---

# oktapam_project_group (Data Source)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `project_name` (String)

### Optional

- `create_server_group` (Boolean)
- `has_no_selectors` (Boolean)
- `has_selectors` (Boolean)
- `include_removed` (Boolean)
- `offline_enabled` (Boolean)

### Read-Only

- `id` (String) The ID of this resource.
- `project_groups` (List of Object) (see [below for nested schema](#nestedatt--project_groups))

<a id="nestedatt--project_groups"></a>
### Nested Schema for `project_groups`

Read-Only:

- `create_server_group` (Boolean)
- `deleted_at` (String)
- `group_id` (String)
- `group_name` (String)
- `project_name` (String)
- `removed_at` (String)
- `server_access` (Boolean)
- `server_admin` (Boolean)
- `servers_selector` (Map of String)

