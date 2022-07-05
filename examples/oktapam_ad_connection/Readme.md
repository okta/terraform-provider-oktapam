# Creating AD Connection and Task Settings

This is an example of setting up AD Connection and task settings for AD-Joined functionality. AD-Joined provides Remote Desktop Protocol (RDP) access using existing Active Directory Accounts (AD) accounts. Refer [Creating AD Connection](https://help.okta.com/asa/en-us/Content/Topics/Adv_Server_Access/docs/ad-connections.htm) for more details.

To run the example:

#### Prerequisites
* Create service user required to use the provider [here](../../README.md#using-the-provider).
* Update the variables.tf OR provide overrides.
* Currently, OktaPAM terraform provider doesn't support creating gateway. Follow instructions at [Setup Gateway](https://help.okta.com/asa/en-us/Content/Topics/Adv_Server_Access/docs/ad-gateways.htm) to create one before creating AD Connections and Tasks.
