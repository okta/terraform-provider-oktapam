package oktapam

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

func TestAccDatasourceDatabaseResourceFetch(t *testing.T) {
	checkTeamApplicable(t, true)

	prefix := "data.oktapam_database"
	groupName := fmt.Sprintf("test_acc_group_%s", randSeq())
	resourceGroupName := fmt.Sprintf("test_acc_resource_group_%s", randSeq())
	projectName := fmt.Sprintf("test_acc_project_%s", randSeq())

	resourceGroupKey := "test_acc_resource_group"
	projectKey := "test_acc_resource_group_project"
	databaseKey := "test_acc_database_resource"
	datasourceKey := "test_acc_database_resource_datasource"

	initConfig := createTestDatabaseCreateConfig(groupName, resourceGroupName, projectName)

	dataSourceName := fmt.Sprintf("%s.%s", prefix, datasourceKey)
	dataConfig := createTestAccDatasourceDatabaseResourceDataConfig(datasourceKey, databaseKey, resourceGroupKey, projectKey)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccV6ProviderFactories(),
		CheckDestroy:             testAccResourceGroupCheckDestroy(resourceGroupName),
		Steps: []resource.TestStep{
			{
				Config: initConfig,
			},
			{
				Config: fmt.Sprintf("%s\n%s", initConfig, dataConfig),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, fmt.Sprintf("%s.#", attributes.ManagementConnectionDetails), "1"),
					resource.TestCheckResourceAttr(dataSourceName, fmt.Sprintf("%s.0.%%", attributes.ManagementConnectionDetails), "1"),
					resource.TestCheckResourceAttr(dataSourceName,
						fmt.Sprintf("%s.0.%s.0.%s.0.%s",
							attributes.ManagementConnectionDetails, attributes.MySQL, attributes.BasicAuth, attributes.Username),
						"user"),
				),
			},
		},
	})
}

const testAccDatasourceDatabaseResourceDataConfigFormat = `
data "oktapam_database" "%s" {
	resource_group = oktapam_resource_group.%s.id
	project = oktapam_resource_group_project.%s.id
	id = oktapam_database.%s.id
}
`

func createTestAccDatasourceDatabaseResourceDataConfig(datasourceKey, databaseName, resourceGroupName, projectName string) string {
	return fmt.Sprintf(testAccDatasourceDatabaseResourceDataConfigFormat, datasourceKey, resourceGroupName, projectName, databaseName)
}
