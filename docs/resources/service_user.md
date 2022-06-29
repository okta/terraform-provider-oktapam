---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "oktapam_service_user Resource - terraform-provider-oktapam"
subcategory: ""
description: |-
  An ASA User of type service, responsible for automation.
---

# oktapam_service_user (Resource)

An ASA User of type `service`, responsible for automation.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The human-readable name of the resource. Values are case-sensitive.

### Optional

- `status` (String) The status of the ASA User. Valid statuses are `ACTIVE`, `DISABLED`, and `DELETED`.

### Read-Only

- `deleted_at` (String) The UTC time of resource creation. Format is `2022-01-01 00:00:00 +0000 UTC`.
- `id` (String) The ID of this resource.
- `server_user_name` (String) The name of the corresponding ASA Server User.
- `team_name` (String) The human-readable name of the ASA Team that owns the resource. Values are lower-case.
- `user_type` (String) The user type. Valid types are `human` and `service`.

