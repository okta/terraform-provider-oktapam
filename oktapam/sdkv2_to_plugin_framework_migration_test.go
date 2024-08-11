package oktapam

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestResource_Resource_Group_UpgradeFromSdkv2(t *testing.T) {
	groupName := fmt.Sprintf("test_acc_group_%s", randSeq())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccGroupCheckDestroy(groupName),
		Steps: []resource.TestStep{
			{

				ExternalProviders: map[string]resource.ExternalProvider{
					"oktapam": {
						VersionConstraint: "0.4.1",
						Source:            "okta.com/pam/oktapam",
					},
				},

				Config: createTestAccGroupCreateConfig(groupName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("oktapam_group.test_group", "name", groupName),
				),
			},
			{
				ProtoV6ProviderFactories: testAccV6ProviderFactories,
				Config:                   createTestAccGroupCreateConfig(groupName),
				// // ConfigPlanChecks is a terraform-plugin-testing feature.
				// // If acceptance testing is still using terraform-plugin-sdk/v2,
				// // use `PlanOnly: true` instead. When migrating to
				// // terraform-plugin-testing, switch to `ConfigPlanChecks` or you
				// // will likely experience test failures.
				// ConfigPlanChecks: resource.ConfigPlanChecks{
				// 	PreApply: []plancheck.PlanCheck{
				// 		&plancheck.ExpectEmptyPlan{},
				// 	},
				// },
				PlanOnly: true,
			},
		},
	})
}

func TestDataSource_Security_Policies_UpgradeFromSdkv2(t *testing.T) {
	identifier := randSeq()
	validServerID := getValidServerID()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccSecurityPoliciesCheckDestroy(identifier+"-1", identifier+"-2"),
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"oktapam": {
						VersionConstraint: "0.4.1",
						Source:            "okta.com/pam/oktapam",
					},
				},
				Config: createTestAccDatasourceSecurityPoliciesInitConfig(identifier, validServerID),
			},
			{
				ProtoV6ProviderFactories: testAccV6ProviderFactories,
				Config:                   createTestAccDatasourceSecurityPoliciesInitConfig(identifier, validServerID),
				PlanOnly:                 true,
			},
		},
	})
}
