package oktapam

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

func TestAccDatasourceSecretFetch(t *testing.T) {
	checkTeamApplicable(t, true)

	resourceGroupName := fmt.Sprintf("test_acc_resource_group_%s", randSeq())
	delegatedAdminGroupName := fmt.Sprintf("test_acc_resource_group_dga_%s", randSeq())
	projectName := fmt.Sprintf("test_acc_resource_group_project_%s", randSeq())
	initialTopLevelFolderName := fmt.Sprintf("test_acc_secret_folder_%s", randSeq())
	topLevelFolderSecurityPolicyName := fmt.Sprintf("test_acc_secret_folder_security_policy_%s", randSeq())
	secret1Name := fmt.Sprintf("test_acc_secret_%s", randSeq())
	secret2Name := fmt.Sprintf("test_acc_secret_%s", randSeq())

	initConfig := createTestAccDatasourceSecretCreateConfig(delegatedAdminGroupName, resourceGroupName, projectName, initialTopLevelFolderName, topLevelFolderSecurityPolicyName, secret1Name, secret2Name)

	prefix := "data.oktapam_secret"
	dataSource1Name := fmt.Sprintf("%s.%s", prefix, "data1")
	data1Config := createTestAccDatasourceSecretDataConfig("data1", "oktapam_secret.test_acc_secret_1.id")
	dataSource2Name := fmt.Sprintf("%s.%s", prefix, "data2")
	data2Config := createTestAccDatasourceSecretDataConfig("data2", "oktapam_secret.test_acc_secret_2.id")

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccResourceGroupCheckDestroy(resourceGroupName),
		Steps: []resource.TestStep{
			{
				Config: initConfig,
			},
			{
				Config: fmt.Sprintf("%s\n%s", initConfig, data1Config),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(dataSource1Name, attributes.Name, secret1Name),
					resource.TestCheckResourceAttr(dataSource1Name, attributes.Description, "secret for test 1"),
					resource.TestCheckResourceAttr(dataSource1Name, fmt.Sprintf("%s.data", attributes.Secret), "test secret 1"),
				),
			},
			{
				Config: fmt.Sprintf("%s\n%s", initConfig, data2Config),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(dataSource2Name, attributes.Name, secret2Name),
					resource.TestCheckResourceAttr(dataSource2Name, attributes.Description, "secret for test 2"),
					resource.TestCheckResourceAttr(dataSource2Name, fmt.Sprintf("%s.data", attributes.Secret), "test secret 2"),
				),
			},
		},
	})
}

const testAccDatasourceSecretDataConfigFormat = `
data "oktapam_secret" "%s" {
	id = %s
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	project = oktapam_resource_group_project.test_acc_resource_group_project.id
}
`

func createTestAccDatasourceSecretDataConfig(resourceName string, secretID string) string {
	return fmt.Sprintf(testAccDatasourceSecretDataConfigFormat, resourceName, secretID)
}

const testAccDatasourceSecretCreateConfigFormat = `
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
	description = "top-level folder for secrets datasource test"
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	project = oktapam_resource_group_project.test_acc_resource_group_project.id
}
resource "oktapam_security_policy" "test_acc_secret_folder_security_policy" {
	name = "%s"
	description = "security policy for top-level folder created in secrets data source test"
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
			  	secret_create = true
			  	secret_delete = true
			  	secret_reveal = true
			  	secret_update = true
			}
		}
	}
}
resource "oktapam_secret" "test_acc_secret_1" {
	name = "%s"
	description = "secret for test 1"
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	project = oktapam_resource_group_project.test_acc_resource_group_project.id
	parent_folder = oktapam_secret_folder.test_acc_secret_folder_top_level.id
	secret = {
		data = "test secret 1"
	}
	depends_on = [oktapam_security_policy.test_acc_secret_folder_security_policy, oktapam_user_group_attachment.test_resource_group_dga_group_attachment]
}
resource "oktapam_secret" "test_acc_secret_2" {
	name = "%s"
	description = "secret for test 2"
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	project = oktapam_resource_group_project.test_acc_resource_group_project.id
	parent_folder = oktapam_secret_folder.test_acc_secret_folder_top_level.id
	secret = {
		data = "test secret 2"
	}
	depends_on = [oktapam_security_policy.test_acc_secret_folder_security_policy, oktapam_user_group_attachment.test_resource_group_dga_group_attachment]
}
`

func createTestAccDatasourceSecretCreateConfig(dgaName, resourceGroupName, projectName, topLevelFolderName, topLevelFolderSecurityPolicyName, secret1Name, secret2Name string) string {
	return fmt.Sprintf(testAccDatasourceSecretCreateConfigFormat, dgaName, resourceGroupName, projectName, topLevelFolderName, topLevelFolderSecurityPolicyName, secret1Name, secret2Name)
}
