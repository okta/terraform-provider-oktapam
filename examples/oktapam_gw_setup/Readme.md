# Deploying a Gateway with Terraform

This is an example of creating a Gateway Setup Token and using it to deploy the Gateway in GCP. Once the Gateway has been deployed, this example creates a Project configured to use the Gateway for traffic forwarding.

Refer to [ASA Gateways](https://help.okta.com/asa/en-us/Content/Topics/Adv_Server_Access/docs/gateways.htm) for more information about gateways.


To run the example:

#### Prerequisites
* Create service user required to use the provider [here](../../README.md#using-the-provider).
* This example deploys a virtual machine on Google Cloud Platform using the GCP Terraform provider. Note that running this example creates real resources in your GCP account so be sure to destroy the resources if they are no longer needed.
* Update the variables.tf OR provide overrides.
