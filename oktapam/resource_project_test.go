package oktapam

import (
	"context"
	"fmt"
	"testing"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/kylelemons/godebug/pretty"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/utils"
)

func TestAccProject(t *testing.T) {
	checkTeamApplicable(t, false)
	resourceName := "oktapam_project.test_project"
	projectName := fmt.Sprintf("test_acc_project_%s", randSeq())
	initialProject := client.Project{
		Name:                   &projectName,
		NextUnixUID:            utils.AsIntPtr(60120),
		NextUnixGID:            utils.AsIntPtr(63020),
		CreateServerUsers:      utils.AsBoolPtrZero(true, true),
		ForwardTraffic:         utils.AsBoolPtrZero(false, true),
		RequirePreAuthForCreds: utils.AsBoolPtrZero(false, true),
		RDPSessionRecording:    utils.AsBoolPtrZero(false, true),
		SSHSessionRecording:    utils.AsBoolPtrZero(false, true),
		SSHCertificateType:     utils.AsStringPtr("CERT_TYPE_ED25519_01"),
		UserOnDemandPeriod:     utils.AsIntPtr(1),
	}
	updatedProject := client.Project{
		Name:                   &projectName,
		NextUnixUID:            utils.AsIntPtr(61200),
		NextUnixGID:            utils.AsIntPtr(63400),
		CreateServerUsers:      utils.AsBoolPtrZero(true, true),
		ForwardTraffic:         utils.AsBoolPtrZero(true, true),
		RequirePreAuthForCreds: utils.AsBoolPtrZero(false, true),
		RDPSessionRecording:    utils.AsBoolPtrZero(true, true),
		SSHSessionRecording:    utils.AsBoolPtrZero(true, true),
		GatewaySelector:        utils.AsStringPtr("env=test"),
		SSHCertificateType:     utils.AsStringPtr("CERT_TYPE_RSA_01"),
		UserOnDemandPeriod:     utils.AsIntPtr(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccV6ProviderFactories,
		CheckDestroy:             testAccProjectCheckDestroy(projectName),
		Steps: []resource.TestStep{
			{
				Config: createTestAccProjectCreateConfig(projectName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectCheckExists(resourceName, initialProject),
					resource.TestCheckResourceAttr(
						resourceName, attributes.Name, projectName,
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.NextUnixUID, "60120",
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.NextUnixGID, "63020",
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.SSHCertificateType, "CERT_TYPE_ED25519_01",
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.UserOnDemandPeriod, "1",
					),
				),
			},
			{
				Config: createTestAccProjectUpdateConfig(projectName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccProjectCheckExists(resourceName, updatedProject),
					resource.TestCheckResourceAttr(
						resourceName, attributes.Name, projectName,
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.NextUnixUID, "61200",
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.NextUnixGID, "63400",
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.SSHCertificateType, "CERT_TYPE_RSA_01",
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.UserOnDemandPeriod, "10",
					),
					resource.TestCheckResourceAttr(
						resourceName, attributes.RDPSessionRecording, "true",
					),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccProjectImportStateId(resourceName),
			},
		},
	})
}

func testAccProjectCheckExists(rn string, expectedProject client.Project) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[rn]
		if !ok {
			return fmt.Errorf("resource not found: %s", rn)
		}

		resourceID := rs.Primary.ID
		if resourceID == "" {
			return fmt.Errorf("resource id not set")
		}

		client := getTestAccAPIClients().LocalClient
		proj, err := client.GetProject(context.Background(), *expectedProject.Name, false)
		if err != nil {
			return fmt.Errorf("error getting project :%w", err)
		}
		if proj == nil || utils.IsBlank(proj.ID) {
			return fmt.Errorf("project %s does not exist", *expectedProject.Name)
		}
		expectedProject.ID = proj.ID
		expectedProject.Team = &client.Team
		comparison := pretty.Compare(proj, expectedProject)
		if comparison != "" {
			return fmt.Errorf("expected project does not match returned project.\n%s", comparison)
		}
		return nil
	}
}

func testAccProjectCheckDestroy(projectName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := getTestAccAPIClients().LocalClient
		proj, err := client.GetProject(context.Background(), projectName, false)
		if err != nil {
			return fmt.Errorf("error getting project: %w", err)
		}

		if proj != nil && proj.Exists() {
			return fmt.Errorf("project still exists")
		}

		return nil
	}
}

const testAccProjectCreateConfigFormat = `
resource "oktapam_project" "test_project" {
	name                  = "%s"
	next_unix_uid         = 60120
	next_unix_gid         = 63020
	ssh_certificate_type  = "CERT_TYPE_ED25519_01"
	user_on_demand_period = 1
}`

func createTestAccProjectCreateConfig(projectName string) string {
	return fmt.Sprintf(testAccProjectCreateConfigFormat, projectName)
}

const testAccProjectUpdateConfigFormat = `
resource "oktapam_project" "test_project" {
    name                      = "%s"
  	next_unix_uid             = 61200
  	next_unix_gid             = 63400
	create_server_users       = true
	forward_traffic           = true
	rdp_session_recording     = true
	ssh_session_recording     = true
	gateway_selector          = "env=test"
	ssh_certificate_type      = "CERT_TYPE_RSA_01"
	user_on_demand_period     = 10
}`

func createTestAccProjectUpdateConfig(projectName string) string {
	return fmt.Sprintf(testAccProjectUpdateConfigFormat, projectName)
}

func testAccProjectImportStateId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("Not found: %s", resourceName)
		}
		return rs.Primary.Attributes[attributes.Name], nil
	}
}
