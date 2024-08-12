package oktapam

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

func TestAccResourceGroupUpgradeFromSdkv2(t *testing.T) {
	checkTeamApplicable(t, true)
	groupName := fmt.Sprintf("test_acc_group_%s", randSeq())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccGroupCheckDestroy(groupName),
		Steps: []resource.TestStep{
			{

				ExternalProviders: map[string]resource.ExternalProvider{
					"oktapam": {
						VersionConstraint: "0.4.0",
						Source:            "okta/oktapam",
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

func TestAccDataSourceResourceGroupProjectUpgradeFromSdkv2(t *testing.T) {
	checkTeamApplicable(t, true)
	identifier := randSeq()
	// config to create the resources
	initConfig := createTestAccDatasourceResourceGroupProjectInitConfig(identifier)

	// config for the datasources
	fetchConfig := testAccDatasourceResourceGroupProjectConfig(identifier)

	resourceName := "data.oktapam_resource_group_project.rg_project"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccSecurityPoliciesCheckDestroy(identifier+"-1", identifier+"-2"),
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"oktapam": {
						VersionConstraint: "0.4.0",
						Source:            "okta/oktapam",
					},
				},
				Config: initConfig,
			},
			{
				ProtoV6ProviderFactories: testAccV6ProviderFactories,
				Config:                   initConfig,
				PlanOnly:                 true,
			},
			{
				ProtoV6ProviderFactories: testAccV6ProviderFactories,
				Config:                   fmt.Sprintf("%s\n%s", initConfig, fetchConfig),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						resourceName, attributes.Name, identifier,
					),
				),
			},
		},
	})
}
