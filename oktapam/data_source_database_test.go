package oktapam

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
	"testing"
)

func TestAccDatasourceDatabaseResourceFetch(t *testing.T) {
	checkTeamApplicable(t, true)

	prefix := "data.oktapam_database"
	groupName := "test_acc_resource_group_dga_group"
	resourceGroupName := "test_acc_resource_group"
	projectName := "test_acc_resource_group_project"
	databaseName := "test_acc_database_resource"

	initConfig := createTestDatabaseCreateConfig(groupName, resourceGroupName, projectName)

	dataSourceName := fmt.Sprintf("%s.%s", prefix, databaseName)
	dataConfig := createTestAccDatasourceDatabaseResourceDataConfig(databaseName, resourceGroupName, projectName)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccResourceGroupCheckDestroy(resourceGroupName),
		Steps: []resource.TestStep{
			{
				Config: initConfig,
			},
			{
				Config: fmt.Sprintf("%s\n%s", initConfig, dataConfig),
				Check:  resource.TestCheckResourceAttr(dataSourceName, attributes.ManagementConnectionDetailsType, "mysql.basic_auth"),
			},
		},
	})
}

const testAccDatasourceDatabaseResourceDataConfigFormat = `
data "oktapam_database" "%[1]s" {
	resource_group = oktapam_resource_group.%s.id
	project = oktapam_resource_group_project.%s.id
	id = oktapam_database.%[1]s.id
}
`

func createTestAccDatasourceDatabaseResourceDataConfig(databaseName, resourceGroupName, projectName string) string {
	return fmt.Sprintf(testAccDatasourceDatabaseResourceDataConfigFormat, databaseName, resourceGroupName, projectName)
}
