package oktapam

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"text/template"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

func TestAccKubernetesCluster(t *testing.T) {
	if _, ok := os.LookupEnv("SFT_KUBERNETES_BETA"); !ok {
		t.Skip("skipping Kubernetes tests")
	}

	resourceName := "oktapam_kubernetes_cluster.acctest_cluster"

	clusterKey := id.PrefixedUniqueId("cluster-key-")
	authMechanism := "NONE"

	labels1 := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	labels2 := map[string]string{
		"key3": "value3",
	}

	cluster1 := client.KubernetesCluster{
		Key:    &clusterKey,
		Auth:   &authMechanism,
		Labels: labels1,
	}

	cluster2 := client.KubernetesCluster{
		Key:    &clusterKey,
		Auth:   &authMechanism,
		Labels: labels2,
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccClusterCheckDestroy(cluster1),
		Steps: []resource.TestStep{
			{
				// create a normal cluster with some labels.
				Config: createTestAccKubernetesClusterConfig(cluster1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, attributes.KubernetesClusterKey, clusterKey),
					resource.TestCheckResourceAttr(resourceName, attributes.KubernetesAuthMechanism, authMechanism),
					resource.TestCheckResourceAttr(resourceName, attributes.Labels+".%", "2"),
					resource.TestCheckResourceAttr(resourceName, attributes.Labels+".key1", "value1"),
					resource.TestCheckResourceAttr(resourceName, attributes.Labels+".key2", "value2"),
				),
			},
			{
				// make sure import works
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// update the cluster labels (the only attribute that doesn't force a new resource)
				Config: createTestAccKubernetesClusterConfig(cluster2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, attributes.Labels+".%", "1"),
					resource.TestCheckResourceAttr(resourceName, attributes.Labels+".key3", "value3"),
				),
			},
		},
	})
}

func createTestAccKubernetesClusterConfig(cluster client.KubernetesCluster) string {
	tpl, err := template.New("cluster").New("cluster").Parse(clusterResourceTemplate)
	if err != nil {
		panic(err)
	}

	var output strings.Builder
	if err := tpl.Execute(&output, cluster); err != nil {
		panic(err)
	}
	return output.String()
}

func testAccClusterCheckDestroy(expectedCluster client.KubernetesCluster) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		pamClient := testAccProvider.Meta().(client.OktaPAMClient)
		clusters, err := pamClient.ListKubernetesClusters(context.Background())
		if err != nil {
			return fmt.Errorf("error getting cluster groups: %w", err)
		}
		for _, cluster := range clusters {
			if *expectedCluster.Key == *cluster.Key {
				return fmt.Errorf("cluster %q still exists", *expectedCluster.Key)
			}
		}
		return nil
	}
}

const clusterResourceTemplate = `
resource "oktapam_kubernetes_cluster" "acctest_cluster" {
	key = "{{.Key}}"
    auth_mechanism = "{{.Auth}}"
    labels = {
{{- range $key, $value := .Labels }}
      {{ $key }} = "{{$value}}"
{{- end }}
   }
}
`
