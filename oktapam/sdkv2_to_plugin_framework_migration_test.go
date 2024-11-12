package oktapam

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

// TestAccResourceGroupUpgradeFromSdkv2 tests the resource group resource could be created using the old SDKv2 provider
// and then verified that the new combined provider could handle the same configuration without any issues. It should not have any plan changes.
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
				ProtoV6ProviderFactories: testAccV6ProviderFactories(),
				Config:                   createTestAccGroupCreateConfig(groupName),
				// assert plan is a no-op prior to, terraform apply phase
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

// TestAccDataSourceResourceGroupProjectUpgradeFromSdkv2 tests the datasource resource group project could be created using the old SDKv2 provider
// and then verified that the new combined provider could handle the same configuration without any issues. It should not have any plan changes.
// Also, it verifies that the datasource could fetch the created resource group project with the combined provider.
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
				ProtoV6ProviderFactories: testAccV6ProviderFactories(),
				Config:                   initConfig,
				PlanOnly:                 true,
			},
			{
				ProtoV6ProviderFactories: testAccV6ProviderFactories(),
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
