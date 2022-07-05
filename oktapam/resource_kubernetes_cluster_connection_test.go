package oktapam

import (
	"encoding/pem"
	"strings"
	"testing"
	"text/template"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

func TestAccKubernetesClusterConnection(t *testing.T) {
	resourceName := "oktapam_kubernetes_cluster_connection.acctest"
	apiURL := "https://localhost:6443"

	publicCertificate := string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: make([]byte, 1000)}))
	publicCertificateToSend := strings.Replace(publicCertificate, "\n", "\\n", -1)

	clusterConnection := client.KubernetesClusterConnection{
		APIURL:            &apiURL,
		PublicCertificate: &publicCertificateToSend,
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: createTestAccKubernetesClusterConnectionConfig(resource.PrefixedUniqueId("cluster-key-"), clusterConnection),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, attributes.KubernetesAPIURL, apiURL),
					resource.TestCheckResourceAttr(resourceName, attributes.PublicCertificate, publicCertificate),
				),
			},
			{
				// make sure import works
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func createTestAccKubernetesClusterConnectionConfig(clusterKey string, clusterConnection client.KubernetesClusterConnection) string {
	tpl, err := template.New("cluster_connection").New("cluster").Parse(clusterConnectionResourceTemplate)
	if err != nil {
		panic(err)
	}

	type templateData struct {
		client.KubernetesClusterConnection
		ClusterKey string
	}

	var output strings.Builder
	if err := tpl.Execute(&output, templateData{
		KubernetesClusterConnection: clusterConnection,
		ClusterKey:                  clusterKey,
	}); err != nil {
		panic(err)
	}
	return output.String()
}

const clusterConnectionResourceTemplate = `
resource "oktapam_kubernetes_cluster" "acctest" {
	key = "{{.ClusterKey}}"
	auth_mechanism="NONE"
}

resource "oktapam_kubernetes_cluster_connection" "acctest" {
	cluster_id=oktapam_kubernetes_cluster.acctest.id
	api_url = "{{.APIURL}}"
	public_certificate = "{{.PublicCertificate}}"
}
`
