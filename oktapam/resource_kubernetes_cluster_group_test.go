package oktapam

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"text/template"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
)

func TestAccKubernetesClusterGroup(t *testing.T) {
	if _, ok := os.LookupEnv("SFT_KUBERNETES_BETA"); !ok {
		t.Skip("skipping Kubernetes tests")
	}

	resourceName := "oktapam_kubernetes_cluster_group.test_group"

	groupName := id.PrefixedUniqueId("cluster-group-test-group-")
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
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, attributes.ClusterSelector, clusterSelector1),
					resource.TestCheckResourceAttr(resourceName, attributes.GroupName, groupName),
					resource.TestCheckResourceAttr(resourceName, attributes.Claims+".%", "2"),
					resource.TestCheckResourceAttr(resourceName, attributes.Claims+".claim1", "val1,val2"),
					resource.TestCheckResourceAttr(resourceName, attributes.Claims+".claim2", "val3"),
				),
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
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, attributes.ClusterSelector, clusterSelector2),
					resource.TestCheckResourceAttr(resourceName, attributes.Claims+".%", "1"),
					resource.TestCheckResourceAttr(resourceName, attributes.Claims+".claim3", "val4"),
				),
			},
		},
	})
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
		pamClient := getLocalClientFromMetadata(testAccProvider.Meta())
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

const clusterGroup1ResourceTemplate = `
resource "oktapam_group" "test_group" {
	name = "{{ .GroupName }}"
	roles = ["access_user"]
}

resource "oktapam_kubernetes_cluster_group" "test_group" {
  group_name       = oktapam_group.test_group.name
  cluster_selector = "{{ .ClusterSelector }}"

  claims = {
    claim1 = "val1,val2"
    claim2 = "val3"
  }
}
`

const clusterGroup2ResourceTemplate = `
resource "oktapam_group" "test_group" {
	name = "{{ .GroupName }}"
	roles = ["access_user"]
}

resource "oktapam_kubernetes_cluster_group" "test_group" {
  group_name       = oktapam_group.test_group.name
  cluster_selector = "{{ .ClusterSelector }}"

  claims = {
	claim3 = "val4"
  }
}
`
