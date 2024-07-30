package oktapam

import (
	"context"
	"fmt"
	"testing"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/kylelemons/godebug/pretty"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
	"github.com/okta/terraform-provider-oktapam/oktapam/client/wrappers"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

func TestAccSecret(t *testing.T) {
	checkTeamApplicable(t, true)

	secretResourceName := "oktapam_secret.test_acc_secret_secret"

	resourceGroupName := fmt.Sprintf("test_acc_resource_group_%s", randSeq())
	delegatedAdminGroupName := fmt.Sprintf("test_acc_resource_group_dga_%s", randSeq())
	projectName := fmt.Sprintf("test_acc_resource_group_project_%s", randSeq())
	topLevelFolderName := fmt.Sprintf("test_acc_secret_%s", randSeq())
	topLevelFolderSecurityPolicyName := fmt.Sprintf("test_acc_secret_security_policy_%s", randSeq())
	initialSecretName := fmt.Sprintf("test_acc_secret_secret_%s", randSeq())
	initialSecretDescription := "secret for secret resource test"

	initialSecret := &wrappers.SecretWrapper{
		Secret: &pam.Secret{
			Name:        initialSecretName,
			Description: *pam.NewNullableString(&initialSecretDescription),
		},
		SecretContents: map[string]string{"test": "data"},
	}

	updatedSecretName := fmt.Sprintf("test_acc_secret_secret_%s", randSeq())
	updatedSecretDescription := "updated secret for secret resource test"

	updatedSecret := &wrappers.SecretWrapper{
		Secret: &pam.Secret{
			Name:        updatedSecretName,
			Description: *pam.NewNullableString(&updatedSecretDescription),
		},
		SecretContents: map[string]string{"test": "other_data"},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccV6ProviderFactories,
		CheckDestroy:             testAccResourceGroupCheckDestroy(resourceGroupName),
		Steps: []resource.TestStep{
			{
				Config: createTestAccSecretCreateConfig(delegatedAdminGroupName, resourceGroupName, projectName, topLevelFolderName, topLevelFolderSecurityPolicyName, initialSecretName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccSecretCheckExists(secretResourceName, initialSecret),
					resource.TestCheckResourceAttr(
						secretResourceName, attributes.Name, initialSecretName,
					),
					resource.TestCheckResourceAttr(
						secretResourceName, attributes.Description, initialSecretDescription,
					),
				),
			},
			{
				Config: createTestAccSecretUpdateConfig(delegatedAdminGroupName, resourceGroupName, projectName, topLevelFolderName, topLevelFolderSecurityPolicyName, updatedSecretName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccSecretCheckExists(secretResourceName, updatedSecret),
					resource.TestCheckResourceAttr(
						secretResourceName, attributes.Name, updatedSecretName,
					),
					resource.TestCheckResourceAttr(
						secretResourceName, attributes.Description, updatedSecretDescription,
					),
				),
			},
		},
	})
}

func testAccSecretCheckExists(rn string, expectedSecret *wrappers.SecretWrapper) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[rn]
		if !ok {
			return fmt.Errorf("resource not found: %s", rn)
		}

		resourceGroupID := rs.Primary.Attributes[attributes.ResourceGroup]
		projectID := rs.Primary.Attributes[attributes.Project]
		secretID := rs.Primary.Attributes[attributes.ID]

		pamClient := getTestAccAPIClients().SDKClient
		secret, err := client.RevealSecret(context.Background(), pamClient, resourceGroupID, projectID, secretID)
		if err != nil {
			return fmt.Errorf("error getting secret: %w", err)
		} else if secret == nil {
			return fmt.Errorf("secret does not exist")
		}

		err = insertComputedValuesForSecret(expectedSecret, secret)
		if err != nil {
			return err
		}

		comparison := pretty.Compare(expectedSecret, secret)
		if comparison != "" {
			return fmt.Errorf("expected secret does not match returned secret.\n%s", comparison)
		}

		return nil
	}
}

func insertComputedValuesForSecret(expectedSecret, actualSecret *wrappers.SecretWrapper) error {
	if expectedSecret.Secret == nil {
		expectedSecret.Secret = &pam.Secret{}
	}
	expectedSecret.Secret.Id = actualSecret.Secret.Id
	expectedSecret.Secret.Path = actualSecret.Secret.Path
	expectedSecret.ResourceGroupID = actualSecret.ResourceGroupID
	expectedSecret.ProjectID = actualSecret.ProjectID
	expectedSecret.ParentFolderID = actualSecret.ParentFolderID

	return nil
}

const testAccSecretCreateConfigFormat = `
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
resource "oktapam_secret_folder" "test_acc_secret_top_level_folder" {
	name = "%s"
	description = "top-level folder for test"
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	project = oktapam_resource_group_project.test_acc_resource_group_project.id
}
resource "oktapam_security_policy" "test_acc_secret_security_policy" {
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
					secret_folder_id = oktapam_secret_folder.test_acc_secret_top_level_folder.id
				}
			}
		}
		privileges {
			secret {
				list = true
				folder_create = true
				folder_delete = true
				folder_update = true
			  	secret_create = true 
			  	secret_delete = true
			  	secret_reveal = true
			  	secret_update = true
			}
		}
	}
	depends_on = [oktapam_user_group_attachment.test_resource_group_dga_group_attachment]
}
resource "oktapam_secret" "test_acc_secret_secret" {
	name = "%s"
	description = "secret for secret resource test"
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	project = oktapam_resource_group_project.test_acc_resource_group_project.id
	parent_folder = oktapam_secret_folder.test_acc_secret_top_level_folder.id
	secret = {
		test = "data"
	}

	depends_on = [oktapam_security_policy.test_acc_secret_security_policy]
}
`

func createTestAccSecretCreateConfig(dgaName, resourceGroupName, projectName, topLevelFolderName, topLevelFolderSecurityPolicyName, secretName string) string {
	return fmt.Sprintf(testAccSecretCreateConfigFormat, dgaName, resourceGroupName, projectName, topLevelFolderName, topLevelFolderSecurityPolicyName, secretName)
}

const testAccSecretUpdateConfigFormat = `
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
resource "oktapam_secret_folder" "test_acc_secret_top_level_folder" {
	name = "%s"
	description = "updated top-level folder for test"
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	project = oktapam_resource_group_project.test_acc_resource_group_project.id
}
resource "oktapam_security_policy" "test_acc_secret_security_policy" {
	name = "%s"
	description = "security policy for top-level folder created in secret test"
	active = true
	principals {
		groups = [oktapam_group.test_resource_group_dga_group.id]
	}
	rule {
		name = "top-level folder rule"
		resources {
			secrets {
				secret_folder {
					secret_folder_id = oktapam_secret_folder.test_acc_secret_top_level_folder.id
				}
			}
		}
		privileges {
			secret {
				list = true
				folder_create = true
				folder_delete = true
				folder_update = true
			  	secret_create = true
			  	secret_delete = true
			  	secret_reveal = true
			  	secret_update = true
			}
		}
	}
	depends_on = [oktapam_user_group_attachment.test_resource_group_dga_group_attachment]
}
resource "oktapam_secret" "test_acc_secret_secret" {
	name = "%s"
	description = "updated secret for secret resource test"
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	project = oktapam_resource_group_project.test_acc_resource_group_project.id
	parent_folder = oktapam_secret_folder.test_acc_secret_top_level_folder.id
	secret = {
		test = "other_data"
	}

	depends_on = [oktapam_security_policy.test_acc_secret_security_policy]
}
`

func createTestAccSecretUpdateConfig(dgaName, resourceGroupName, projectName, topLevelFolderName, topLevelFolderSecurityPolicyName, secretName string) string {
	return fmt.Sprintf(testAccSecretUpdateConfigFormat, dgaName, resourceGroupName, projectName, topLevelFolderName, topLevelFolderSecurityPolicyName, secretName)
}
