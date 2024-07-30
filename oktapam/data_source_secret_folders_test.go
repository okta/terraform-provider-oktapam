package oktapam

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

func TestAccDatasourceSecretFoldersList(t *testing.T) {
	checkTeamApplicable(t, true)

	resourceGroupName := fmt.Sprintf("test_acc_resource_group_%s", randSeq())
	delegatedAdminGroupName := fmt.Sprintf("test_acc_resource_group_dga_%s", randSeq())
	projectName := fmt.Sprintf("test_acc_resource_group_project_%s", randSeq())
	initialTopLevelFolderName := fmt.Sprintf("test_acc_secret_folder_%s", randSeq())
	topLevelFolderSecurityPolicyName := fmt.Sprintf("test_acc_secret_folder_security_policy_%s", randSeq())
	child1FolderName := fmt.Sprintf("test_acc_secret_folder_%s", randSeq())
	child2FolderName := fmt.Sprintf("test_acc_secret_folder_%s", randSeq())

	initConfig := createTestAccDatasourceSecretFolderCreateConfig(delegatedAdminGroupName, resourceGroupName, projectName, initialTopLevelFolderName, topLevelFolderSecurityPolicyName, child1FolderName, child2FolderName)

	prefix := "data.oktapam_secret_folders"
	dataSource1Name := fmt.Sprintf("%s.%s", prefix, "data1")
	data1Config := createTestAccDatasourceSecretFolderDataConfig("data1", "/"+initialTopLevelFolderName, false)
	dataSource2Name := fmt.Sprintf("%s.%s", prefix, "data2")
	data2Config := createTestAccDatasourceSecretFolderDataConfig("data2", "/"+initialTopLevelFolderName, true)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccV6ProviderFactories,
		CheckDestroy:             testAccResourceGroupCheckDestroy(resourceGroupName),
		Steps: []resource.TestStep{
			{
				Config: initConfig,
			},
			{
				Config: fmt.Sprintf("%s\n%s", initConfig, data1Config),
				Check:  resource.TestCheckResourceAttr(dataSource1Name, fmt.Sprintf("%s.#", attributes.SecretFolders), "1"),
			},
			{
				Config: fmt.Sprintf("%s\n%s", initConfig, data2Config),
				Check:  resource.TestCheckResourceAttr(dataSource2Name, fmt.Sprintf("%s.#", attributes.SecretFolders), "2"),
			},
		},
	})
}

const testAccDatasourceSecretFolderDataConfigFormat = `
data "oktapam_secret_folders" "%s" {
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	project = oktapam_resource_group_project.test_acc_resource_group_project.id

	path = "%s"
	list_elements_under_path = %v
}
`

func createTestAccDatasourceSecretFolderDataConfig(resourceName string, path string, listElementsUnderPath bool) string {
	return fmt.Sprintf(testAccDatasourceSecretFolderDataConfigFormat, resourceName, path, listElementsUnderPath)
}

const testAccDatasourceSecretFolderCreateConfigFormat = `
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
}
resource "oktapam_secret_folder" "test_acc_secret_folder_child_1" {
	name = "%s"
	description = "child folder for test 1"
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	project = oktapam_resource_group_project.test_acc_resource_group_project.id
	parent_folder = oktapam_secret_folder.test_acc_secret_folder_top_level.id
	depends_on = [oktapam_security_policy.test_acc_secret_folder_security_policy]
}
resource "oktapam_secret_folder" "test_acc_secret_folder_child_2" {
	name = "%s"
	description = "child folder for test 2"
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	project = oktapam_resource_group_project.test_acc_resource_group_project.id
	parent_folder = oktapam_secret_folder.test_acc_secret_folder_top_level.id
	depends_on = [oktapam_security_policy.test_acc_secret_folder_security_policy]
}
`

func createTestAccDatasourceSecretFolderCreateConfig(dgaName, resourceGroupName, projectName, topLevelFolderName, topLevelFolderSecurityPolicyName, childFolder1Name, childFolder2Name string) string {
	return fmt.Sprintf(testAccDatasourceSecretFolderCreateConfigFormat, dgaName, resourceGroupName, projectName, topLevelFolderName, topLevelFolderSecurityPolicyName, childFolder1Name, childFolder2Name)
}
