package oktapam

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/kylelemons/godebug/pretty"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

func TestAccSecretFolder(t *testing.T) {
	checkTeamApplicable(t, true)

	topLevelResourceName := "oktapam_secret_folder.test_acc_secret_folder_top_level"
	childResourceName := "oktapam_secret_folder.test_acc_secret_folder_child"

	resourceGroupName := fmt.Sprintf("test_acc_resource_group_%s", randSeq())
	delegatedAdminGroupName := fmt.Sprintf("test_acc_resource_group_dga_%s", randSeq())
	projectName := fmt.Sprintf("test_acc_resource_group_project_%s", randSeq())
	initialTopLevelFolderName := fmt.Sprintf("test_acc_secret_folder_%s", randSeq())
	initialTopLevelFolderDescription := "top-level folder for test"
	topLevelFolderSecurityPolicyName := fmt.Sprintf("test_acc_secret_folder_security_policy_%s", randSeq())
	childFolderName := fmt.Sprintf("test_acc_secret_folder_%s", randSeq())
	childFolderDescription := "child folder for test"

	initialTopLevelSecretFolder := &client.SecretFolder{
		Name:        &initialTopLevelFolderName,
		Description: &initialTopLevelFolderDescription,
	}
	initialChildSecretFolder := &client.SecretFolder{
		Name:        &childFolderName,
		Description: &childFolderDescription,
	}

	updatedTopLevelFolderName := fmt.Sprintf("test_acc_secret_folder_%s", randSeq())
	updatedTopLevelFolderDescription := "updated top-level folder for test"

	updatedTopLevelSecretFolder := &client.SecretFolder{
		Name:        &updatedTopLevelFolderName,
		Description: &updatedTopLevelFolderDescription,
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccV6ProviderFactories,
		CheckDestroy:             testAccResourceGroupCheckDestroy(resourceGroupName),
		Steps: []resource.TestStep{
			{
				Config: createTestAccSecretFolderCreateConfig(delegatedAdminGroupName, resourceGroupName, projectName, initialTopLevelFolderName, topLevelFolderSecurityPolicyName, childFolderName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccSecretFolderCheckExists(topLevelResourceName, initialTopLevelSecretFolder, false),
					testAccSecretFolderCheckExists(childResourceName, initialChildSecretFolder, false),
					resource.TestCheckResourceAttr(
						topLevelResourceName, attributes.Name, initialTopLevelFolderName,
					),
					resource.TestCheckResourceAttr(
						topLevelResourceName, attributes.Description, initialTopLevelFolderDescription,
					),
					resource.TestCheckResourceAttr(
						childResourceName, attributes.Name, childFolderName,
					),
					resource.TestCheckResourceAttr(
						childResourceName, attributes.Description, childFolderDescription,
					),
				),
			},
			{
				Config: createTestAccSecretFolderUpdateConfig(delegatedAdminGroupName, resourceGroupName, projectName, updatedTopLevelFolderName, topLevelFolderSecurityPolicyName, childFolderName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccSecretFolderCheckExists(topLevelResourceName, updatedTopLevelSecretFolder, false),
					testAccSecretFolderCheckExists(childResourceName, initialChildSecretFolder, false),
					resource.TestCheckResourceAttr(
						topLevelResourceName, attributes.Name, updatedTopLevelFolderName,
					),
					resource.TestCheckResourceAttr(
						topLevelResourceName, attributes.Description, updatedTopLevelFolderDescription,
					),
					resource.TestCheckResourceAttr(
						childResourceName, attributes.Name, childFolderName,
					),
					resource.TestCheckResourceAttr(
						childResourceName, attributes.Description, childFolderDescription,
					),
				),
			},
		},
	})
}

func testAccSecretFolderCheckExists(rn string, expectedSecretFolder *client.SecretFolder, ensureHasParent bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[rn]
		if !ok {
			return fmt.Errorf("resource not found: %s", rn)
		}

		resourceGroupID := rs.Primary.Attributes[attributes.ResourceGroup]
		projectID := rs.Primary.Attributes[attributes.Project]
		secretFolderID := rs.Primary.Attributes[attributes.ID]
		pamClient := getTestAccAPIClients().LocalClient
		secretFolder, err := pamClient.GetSecretFolder(context.Background(), resourceGroupID, projectID, secretFolderID)
		if err != nil {
			return fmt.Errorf("error getting secret folder: %w", err)
		} else if secretFolder == nil {
			return fmt.Errorf("secret folder does not exist")
		}

		err = insertComputedValuesForSecretFolder(expectedSecretFolder, secretFolder)
		if err != nil {
			return err
		}

		comparison := pretty.Compare(expectedSecretFolder, secretFolder)
		if comparison != "" {
			return fmt.Errorf("expected secret folder does not match returned secret folder.\n%s", comparison)
		}

		if ensureHasParent && secretFolder.ParentFolderID == nil {
			return fmt.Errorf("expected secret folder to have a parent folder")
		}

		return nil
	}
}

func insertComputedValuesForSecretFolder(expectedSecretFolder, actualSecretFolder *client.SecretFolder) error {
	expectedSecretFolder.ID = actualSecretFolder.ID
	expectedSecretFolder.ResourceGroupID = actualSecretFolder.ResourceGroupID
	expectedSecretFolder.ProjectID = actualSecretFolder.ProjectID
	expectedSecretFolder.ParentFolderID = actualSecretFolder.ParentFolderID

	return nil
}

const testAccSecretFolderCreateConfigFormat = `
data "oktapam_current_user" "current_user" { }
resource "oktapam_group" "test_resource_group_dga_group" {
	name = "%s"
}
resource "oktapam_user_group_attachment" "test_resource_group_dga_group_attachment" {
	username = data.oktapam_current_user.current_user.name
	group = oktapam_group.test_resource_group_dga_group.name
}
resource "oktapam_resource_group" "test_acc_resource_group" {
	name = "%s"
	description = "test resource group"
	delegated_resource_admin_groups = [oktapam_group.test_resource_group_dga_group.id]	
}
resource "oktapam_resource_group_project" "test_acc_resource_group_project" {
	name = "%s"
	resource_group = oktapam_resource_group.test_acc_resource_group.id
}
resource "oktapam_secret_folder" "test_acc_secret_folder_top_level" {
	name = "%s"
	description = "top-level folder for test"
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	project = oktapam_resource_group_project.test_acc_resource_group_project.id
}
resource "oktapam_security_policy" "test_acc_secret_folder_security_policy" {
	name = "%s"
	description = "security policy for top-level folder created in secret folder test"
	active = true
	principals {
		groups = [oktapam_group.test_resource_group_dga_group.id]
	}
	rule {
		name = "top-level folder rule"
		resources {
			secrets {
				secret_folder {
					secret_folder_id = oktapam_secret_folder.test_acc_secret_folder_top_level.id
				}
			}
		}
		privileges {
			secret {
				list = true
				folder_create = true
				folder_delete = true
				folder_update = true
			  	secret_create = false
			  	secret_delete = false
			  	secret_reveal = false
			  	secret_update = false
			}
		}
	}
	depends_on = [oktapam_user_group_attachment.test_resource_group_dga_group_attachment]
}
resource "oktapam_secret_folder" "test_acc_secret_folder_child" {
	name = "%s"
	description = "child folder for test"
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	project = oktapam_resource_group_project.test_acc_resource_group_project.id
	parent_folder = oktapam_secret_folder.test_acc_secret_folder_top_level.id
	depends_on = [oktapam_security_policy.test_acc_secret_folder_security_policy]
}
`

func createTestAccSecretFolderCreateConfig(dgaName, resourceGroupName, projectName, topLevelFolderName, topLevelFolderSecurityPolicyName, childFolderName string) string {
	return fmt.Sprintf(testAccSecretFolderCreateConfigFormat, dgaName, resourceGroupName, projectName, topLevelFolderName, topLevelFolderSecurityPolicyName, childFolderName)
}

const testAccSecretFolderUpdateConfigFormat = `
data "oktapam_current_user" "current_user" { }
resource "oktapam_group" "test_resource_group_dga_group" {
	name = "%s"
}
resource "oktapam_user_group_attachment" "test_resource_group_dga_group_attachment" {
	username = data.oktapam_current_user.current_user.name
	group = oktapam_group.test_resource_group_dga_group.name
}
resource "oktapam_resource_group" "test_acc_resource_group" {
	name = "%s"
	description = "test resource group"
	delegated_resource_admin_groups = [oktapam_group.test_resource_group_dga_group.id]	
}
resource "oktapam_resource_group_project" "test_acc_resource_group_project" {
	name = "%s"
	resource_group = oktapam_resource_group.test_acc_resource_group.id
}
resource "oktapam_secret_folder" "test_acc_secret_folder_top_level" {
	name = "%s"
	description = "updated top-level folder for test"
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	project = oktapam_resource_group_project.test_acc_resource_group_project.id
}
resource "oktapam_security_policy" "test_acc_secret_folder_security_policy" {
	name = "%s"
	description = "security policy for top-level folder created in secret folder test"
	active = true
	principals {
		groups = [oktapam_group.test_resource_group_dga_group.id]
	}
	rule {
		name = "top-level folder rule"
		resources {
			secrets {
				secret_folder {
					secret_folder_id = oktapam_secret_folder.test_acc_secret_folder_top_level.id
				}
			}
		}
		privileges {
			secret {
				list = true
				folder_create = true
				folder_delete = true
				folder_update = true
			  	secret_create = false
			  	secret_delete = false
			  	secret_reveal = false
			  	secret_update = false
			}
		}
	}
	depends_on = [oktapam_user_group_attachment.test_resource_group_dga_group_attachment]
}
resource "oktapam_secret_folder" "test_acc_secret_folder_child" {
	name = "%s"
	description = "child folder for test"
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	project = oktapam_resource_group_project.test_acc_resource_group_project.id
	parent_folder = oktapam_secret_folder.test_acc_secret_folder_top_level.id
	depends_on = [oktapam_security_policy.test_acc_secret_folder_security_policy]
}
`

func createTestAccSecretFolderUpdateConfig(dgaName, resourceGroupName, projectName, topLevelFolderName, topLevelFolderSecurityPolicyName, childFolderName string) string {
	return fmt.Sprintf(testAccSecretFolderUpdateConfigFormat, dgaName, resourceGroupName, projectName, topLevelFolderName, topLevelFolderSecurityPolicyName, childFolderName)
}
