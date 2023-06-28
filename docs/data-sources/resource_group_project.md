---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "oktapam_resource_group_project Data Source - terraform-provider-oktapam"
subcategory: ""
description: |-
  Returns a previously created PAM project associated with a specific PAM resource group. For details, see Projects https://help.okta.com/en/programs/opa-pam/Content/Topics/privileged-access/pam-projects.htm.
---

# oktapam_resource_group_project (Data Source)

Returns a previously created PAM project associated with a specific PAM resource group. For details, see [Projects](https://help.okta.com/en/programs/opa-pam/Content/Topics/privileged-access/pam-projects.htm).



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `resource_group` (String) The UUID of a PAM Resource Group.

### Read-Only

- `account_discovery` (Boolean) If `true`, will enable the discovery of local accounts on servers within the project.
- `deleted_at` (String) The UTC time when the resource was deleted. Format is '2022-01-01 00:00:00 +0000 UTC'.
- `gateway_selector` (String) Assigns Gateways with labels matching all selectors. At least one selector is required to forward traffic through a Gateway.
- `id` (String) The ID of this resource.
- `name` (String) The human-readable name of the resource. Values are case-sensitive.
- `ssh_certificate_type` (String) The SSH certificate type used by access requests. Options include: [`CERT_TYPE_ED25519_01`, `CERT_TYPE_ECDSA_521_01`, `CERT_TYPE_ECDSA_384_01`, `CERT_TYPE_ECDSA_256_01`, `CERT_TYPE_RSA_01`]. 'CERT_TYPE_RSA_01' is a deprecated key algorithm type. This option should only be used to connect to legacy systems that cannot use newer SSH versions. If you do need to use 'CERT_TYPE_RSA_01', it is recommended to connect via a gateway with traffic forwarding. Otherwise, please use a more current key algorithm. If left unspecified, 'CERT_TYPE_ED25519_01' is used by default.
- `team` (String) The human-readable name of the ASA Team that owns the resource. Values are lower-case.

