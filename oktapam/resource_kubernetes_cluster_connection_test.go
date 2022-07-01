package oktapam

import (
	"encoding/pem"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"strings"
	"testing"
	"text/template"
)

func TestAccKubernetesClusterConnection(t *testing.T) {
	resourceName := "oktapam_kubernetes_cluster_connection.acctest"

	apiURL := "https://localhost:6443"

	publicCertificate := string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: make([]byte, 1000)}))
	publicCertificate = strings.Replace(publicCertificate, "\n", "\\n", -1)

	clusterConnection := client.KubernetesClusterConnection{
		APIURL:            &apiURL,
		PublicCertificate: &publicCertificate,
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: createTestAccKubernetesClusterConnectionConfig(clusterConnection),
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

func createTestAccKubernetesClusterConnectionConfig(clusterConnection client.KubernetesClusterConnection) string {
	tpl, err := template.New("cluster_connection").New("cluster").Parse(clusterConnectionResourceTemplate)
	if err != nil {
		panic(err)
	}

	var output strings.Builder
	if err := tpl.Execute(&output, clusterConnection); err != nil {
		panic(err)
	}
	return output.String()
}

const clusterConnectionResourceTemplate = `
resource "oktapam_kubernetes_cluster" "acctest" {
	key = "acctest"
	auth_mechanism="NONE"
}

resource "oktapam_kubernetes_cluster_connection" "acctest" {
	cluster_id=oktapam_kubernetes_cluster.acctest.id
	api_url = "{{.APIURL}}"
	public_certificate = "{{.PublicCertificate}}"
}
`
