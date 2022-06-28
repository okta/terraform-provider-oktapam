# Creating AD Connection and Task Settings

This is an example of setting up an AD Connection and task settings for AD-Joined functionality [Creating AD Connection](https://help.okta.com/asa/en-us/Content/Topics/Adv_Server_Access/docs/ad-connections.htm). AD-Joined provides 
Remote Desktop Protocol (RDP) access using existing Active Directory Accounts (AD) accounts.

To run the example:

#### Prerequisites
* Create service user required to use the provider
* Update the variables.tf OR provide overrides
* OktaPAM terraform provider doesn't support creating gateway. Follow instructions to [Setup Gateway](https://help.okta.com/asa/en-us/Content/Topics/Adv_Server_Access/docs/ad-gateways.htm)
