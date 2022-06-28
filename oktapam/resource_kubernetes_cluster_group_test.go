package oktapam

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"text/template"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/kylelemons/godebug/pretty"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
)

func TestAccKubernetesClusterGroup(t *testing.T) {
	resourceName := "oktapam_kubernetes_cluster_group.test_group"

	groupName := "everyone"
	clusterSelector1 := "select=everything"
	clusterSelector2 := "select=nothing"

	clusterGroup1 := client.KubernetesClusterGroup{
		GroupName:       &groupName,
		ClusterSelector: &clusterSelector1,
		Claims: map[string][]string{
			"claim1": {"val1", "val2"},
			"claim2": {"val3"},
		},
	}

	clusterGroup2 := client.KubernetesClusterGroup{
		GroupName:       &groupName,
		ClusterSelector: &clusterSelector2,
		Claims: map[string][]string{
			"claim3": {"val4"},
			"claim4": {"val5", "val6"},
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccClusterGroupCheckDestroy(clusterGroup1),
		Steps: []resource.TestStep{
			{
				// ensure create works
				Config: createTestAccKubernetesClusterGroupCreateConfig(clusterGroup1ResourceTemplate, clusterGroup1),
				Check:  testAccKubernetesClusterGroupExists(resourceName, clusterGroup1),
			},
			{
				// test importing resource
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// ensure updating cluster selector and claims works
				Config: createTestAccKubernetesClusterGroupCreateConfig(clusterGroup2ResourceTemplate, clusterGroup2),
				Check:  testAccKubernetesClusterGroupExists(resourceName, clusterGroup2),
			},
		},
	})
}

// testAccKubernetesClusterGroupExists ensures the resource created by terraform has the values we expect.
func testAccKubernetesClusterGroupExists(resourceName string, expectedClusterGroup client.KubernetesClusterGroup) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		res, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource %q not found", resourceName)
		}

		resourceID := res.Primary.ID

		pamClient := testAccProvider.Meta().(client.OktaPAMClient)

		clusterGroup, err := pamClient.GetKubernetesClusterGroup(context.Background(), resourceID)
		if err != nil {
			return fmt.Errorf("error getting kubernetes cluster group %q: %w", resourceID, err)
		}

		if clusterGroup == nil {
			return fmt.Errorf("missing kubernetes cluster group %q", resourceID)
		}

		if *expectedClusterGroup.ClusterSelector != *clusterGroup.ClusterSelector {
			return fmt.Errorf("expected cluster selector %q, got %q on cluster group %q", *expectedClusterGroup.ClusterSelector, *clusterGroup.ClusterSelector, resourceID)
		}
		if *expectedClusterGroup.GroupName != *clusterGroup.GroupName {
			return fmt.Errorf("expected cluster group name %q, got %q on cluster group %q", *expectedClusterGroup.GroupName, *expectedClusterGroup.GroupName, resourceID)
		}
		claimDifferences := pretty.Compare(expectedClusterGroup.Claims, clusterGroup.Claims)
		if claimDifferences != "" {
			return fmt.Errorf("cluster group %q claims are different than expected: %q", resourceID, claimDifferences)
		}

		return nil
	}
}

func createTestAccKubernetesClusterGroupCreateConfig(body string, clusterGroup client.KubernetesClusterGroup) string {
	tpl, err := template.New("clusterGroup").Parse(body)
	if err != nil {
		panic(err)
	}

	var output strings.Builder
	if err := tpl.Execute(&output, clusterGroup); err != nil {
		panic(err)
	}
	return output.String()
}

func testAccClusterGroupCheckDestroy(expectedClusterGroup client.KubernetesClusterGroup) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		pamClient := testAccProvider.Meta().(client.OktaPAMClient)
		clusterGroups, err := pamClient.ListKubernetesClusterGroups(context.Background())
		if err != nil {
			return fmt.Errorf("error getting cluster groups: %w", err)
		}
		for _, group := range clusterGroups {
			if *expectedClusterGroup.GroupName == *group.GroupName {
				return fmt.Errorf("cluster group for group %q still exists", *expectedClusterGroup.GroupName)
			}
		}

		return nil
	}
}

// language=Terraform
const clusterGroup1ResourceTemplate = `resource "oktapam_kubernetes_cluster_group" "test_group" {
  group_name       = "{{ .GroupName }}"
  cluster_selector = "{{ .ClusterSelector }}"

  claims = {
    claim1 = "val1,val2"
    claim2 = "val3"
  }
}
`

// language=Terraform
const clusterGroup2ResourceTemplate = `resource "oktapam_kubernetes_cluster_group" "test_group" {
  group_name       = "{{ .GroupName }}"
  cluster_selector = "{{ .ClusterSelector }}"

  claims = {
	claim3 = "val4"
	claim4 = "val5,val6"
  }
}
`
