---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "oktapam_database Data Source - terraform-provider-oktapam"
subcategory: ""
description: |-
  Returns an existing Database. For details, see Databases https://help.okta.com/okta_help.htm?type=oie&id=ext-pam-databases
---

# oktapam_database (Data Source)

Returns an existing Database. For details, see [Databases](https://help.okta.com/okta_help.htm?type=oie&id=ext-pam-databases)



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `project` (String) The UUID of a Project.
- `resource_group` (String) The UUID of a OktaPA Resource Group.

### Optional

- `canonical_name` (String) The nickname or alias of the resource.
- `database_type` (String) Defines the type of database used and the feature set supported. A two-part string separated by a dot: '<db_engine>.level<level>'
- `management_connection_details` (Block List, Max: 1) A set of fields defining the database to connect to. (see [below for nested schema](#nestedblock--management_connection_details))
- `management_gateway_selector` (Map of String) A label selector to define which gateway(s) will be used to connect to the database.
- `recipe_book` (String) The ID of a recipe book which will override the db queries used.

### Read-Only

- `created_at` (String) The UTC time when the resource was created. Format is '2022-01-01 00:00:00 +0000 UTC'.
- `id` (String) The ID of this resource.
- `management_gateway_selector_id` (String) The ID of the selector.
- `updated_at` (String) The UTC time when the resource was last updated. Format is '2022-01-01 00:00:00 +0000 UTC'.

<a id="nestedblock--management_connection_details"></a>
### Nested Schema for `management_connection_details`

Required:

- `mysql` (Block List, Min: 1, Max: 1) A set of fields defining how to connect to a mysql database. (see [below for nested schema](#nestedblock--management_connection_details--mysql))

<a id="nestedblock--management_connection_details--mysql"></a>
### Nested Schema for `management_connection_details.mysql`

Required:

- `basic_auth` (Block List, Min: 1, Max: 1) A set of fields required to authenticate to the database. (see [below for nested schema](#nestedblock--management_connection_details--mysql--basic_auth))
- `hostname` (String) The hostname used to connect to the database
- `port` (String) The port used to connect to the database

<a id="nestedblock--management_connection_details--mysql--basic_auth"></a>
### Nested Schema for `management_connection_details.mysql.basic_auth`

Required:

- `username` (String) The human-readable name of the User.

Optional:

- `password` (String, Sensitive) The password used to authenticate to the database.

Read-Only:

- `secret` (String) The UUID of the secret corresponding to the stored database password.

