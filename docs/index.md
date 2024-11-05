---
page_title: "Provider: okta_pam"
description: |-
---

## oktapam Provider

The Okta PAM Terraform Provider is used to interact with Okta's Privileged Access product.  This provider supports common OPA use cases such as automating the creation of OPA Resource Groups and Projects, obtaining enrollment tokens to enroll servers in OPA Projects as part of server builds, creating Secret Folders, and assigning access via OPA Security policies.

This provider can also be utilized for the Advanced Server Access product.  Not all resources are applicable to both OPA and ASA.  The provider is intended as a full replacement for the [classic ASA Terraform provider](https://registry.terraform.io/providers/oktadeveloper/oktaasa/1.0.1). Users of the oktaasa Terraform provider are encouraged to migrate to the Okta PAM Terraform Provider.

### Key differences between oktaasa provider and Okta PAM Terraform Provider
- Project resources now support additional parameters for configuring server access behaviors.
- Team Roles can now be assigned to ASA Groups.
- ASA Users can be assigned to ASA Groups.
- Introduces support for Gateway Setup Tokens as a resource type, enabling Provider users to automate the creation and deployment of ASA Gateways.
- Adds support for configuring Active Directory connections for supporting AD Joined user authentication.
- Data sources are now available for AD connections, Gateway Setup and Server Enrollment tokens and Projects/Project Groups.

Additionally, this provider adds support for specific Beta features. Users should not use resources marked as 'Beta' in production environments and before confirming that the Beta feature has been enabled for your OPA/ASA  team.

## Authentication
The Okta PAM Providers requires a Service User account that is granted the 'Admin' role be created. This Service User account will be used by the Provider to authenticate to OPA/ASA.


1) Follow the [guide](https://help.okta.com/en-us/content/topics/privileged-access/pam-service-users.htm) to create a Service User account with Admin permissions
- Use your OPA/ASA Team name (same team that the service account in step 1 resides in) as the value for the 'oktapam_team' value.
- Use the Service Account ID for the 'oktapam_key' value
- Use the Service Account key for the 'oktapam_secret' value.

## Schema

### Required

- `oktapam_key` (String) Okta PAM API Key
- `oktapam_secret` (String) Okta PAM API Secret
- `oktapam_team` (String) Okta PAM Team

### Optional

- `oktapam_api_host` (String) Okta PAM API Host.  Note that this will be required when utilizing an OPA account
