package oktapam

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"sync"
	"testing"

	"github.com/atko-pam/pam-sdk-go/client/pam"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/jarcoal/httpmock"
	"github.com/okta/terraform-provider-oktapam/oktapam/constants/attributes"
)

var _ plancheck.PlanCheck = debugPlan{}

type debugPlan struct{}

func (e debugPlan) CheckPlan(ctx context.Context, req plancheck.CheckPlanRequest, resp *plancheck.CheckPlanResponse) {
	rd, err := json.Marshal(req.Plan)
	if err != nil {
		fmt.Println("error marshalling machine-readable plan output:", err)
	}
	fmt.Printf("req.Plan - %s\n", string(rd))
}

func DebugPlan() plancheck.PlanCheck {
	return debugPlan{}
}

const testAccOktaUDCheckoutSettingsBaseConfigFormat = `
resource "oktapam_group" "test_resource_group_dga_group" {
	name = "%s"
}
resource "oktapam_resource_group" "test_acc_resource_group" {
	name = "%s"
	description = "test resource group"
	delegated_resource_admin_groups = [oktapam_group.test_resource_group_dga_group.id]
}
resource "oktapam_resource_group_project" "test_acc_resource_group_project" {
	name = "%s"
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	ssh_certificate_type  = "CERT_TYPE_ED25519_01"
	account_discovery     = true
}
`

const testAccOktaUDCheckoutSettingsCreateConfigFormat = `
resource "oktapam_okta_universal_directory_checkout_settings" "test_acc_okta_universal_directory_checkout_settings" {
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	project = oktapam_resource_group_project.test_acc_resource_group_project.id
	checkout_required = true
	checkout_duration_in_seconds = 900
}
`

const testAccOktaUDCheckoutSettingsUpdateWithIncludeListConfigFormat = `
resource "oktapam_okta_universal_directory_checkout_settings" "test_acc_okta_universal_directory_checkout_settings" {
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	project = oktapam_resource_group_project.test_acc_resource_group_project.id
	checkout_required = true
	checkout_duration_in_seconds = 3600
	include_list = [
		{
			id = "account1",
			service_account_user_name = "user1",
			saas_app_instance_name = "app1"
		},
	]
}
`

const testAccOktaUDCheckoutSettingsUpdateWithExcludeListConfigFormat = `
resource "oktapam_okta_universal_directory_checkout_settings" "test_acc_okta_universal_directory_checkout_settings" {
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	project = oktapam_resource_group_project.test_acc_resource_group_project.id
	checkout_required = true
	checkout_duration_in_seconds = 3600
	exclude_list = [
		{
			id = "account3",
			service_account_user_name = "user3",
			saas_app_instance_name = "app3"
		}
	]
}
`

const testAccOktaUDCheckoutSettingsUpdateWithBothListsConfigFormat = `
resource "oktapam_okta_universal_directory_checkout_settings" "test_acc_okta_universal_directory_checkout_settings" {
	resource_group = oktapam_resource_group.test_acc_resource_group.id
	project = oktapam_resource_group_project.test_acc_resource_group_project.id
	checkout_required = true
	checkout_duration_in_seconds = 7200
	include_list = [
		{
			id = "account1",
			service_account_user_name = "user1",
			saas_app_instance_name = "app1"
		}
	]
	exclude_list = [
		{
			id = "account3",
			service_account_user_name = "user3",
			saas_app_instance_name = "app3"
		}
	]
}
`

func TestAccOktaUDCheckoutSettings(t *testing.T) {
	checkTeamApplicable(t, true)
	resourceName := "oktapam_okta_universal_directory_checkout_settings.test_acc_okta_universal_directory_checkout_settings"
	resourceGroupName := fmt.Sprintf("test_acc_resource_group_%s", randSeq())
	projectName := fmt.Sprintf("test_acc_resource_group_project_%s", randSeq())
	delegatedAdminGroupName := fmt.Sprintf("test_acc_resource_group_dga_%s", randSeq())
	defaultCheckoutDuration := int32(900)

	initialSettings := &pam.APIServiceAccountCheckoutSettings{
		CheckoutRequired:          true,
		CheckoutDurationInSeconds: defaultCheckoutDuration,
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: createOktaUDCheckoutSettingsCreateConfig(delegatedAdminGroupName, resourceGroupName, projectName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccOktaUDCheckoutSettingsCheckExists(resourceName, initialSettings),
				),
			},
			{
				Config:      createOktaUDCheckoutSettingsUpdateWithBothListsConfig(delegatedAdminGroupName, resourceGroupName, projectName),
				ExpectError: regexp.MustCompile(`Only one of 'IncludeList' or 'ExcludeList' can be specified`),
			},
			// Delete Okta UD checkout settings resource
			{
				Config: createOktaUDCheckoutSettingsBaseConfig(delegatedAdminGroupName, resourceGroupName, projectName),
				Check:  testAccOktaUDCheckoutSettingsCheckDeleted(resourceName),
			},
		},
	})
}

// TestAccOktaUDCheckoutSettingsWithMockHTTPClient is a test that uses httpmock to mock the HTTP client
// and test the Okta UD checkout settings resource marshalling and unmarshalling correctly.
func TestAccOktaUDCheckoutSettingsWithMockHTTPClient(t *testing.T) {
	// Use fixed names and IDs for consistency
	resourceName := "oktapam_okta_universal_directory_checkout_settings.test_acc_okta_universal_directory_checkout_settings"
	resourceGroupName := fmt.Sprintf("test_acc_mock_resource_group_%s", randSeq())
	projectName := fmt.Sprintf("test_acc_mock_project_%s", randSeq())
	delegatedAdminGroupName := fmt.Sprintf("test_acc_mock_dga_%s", randSeq())

	// Setup httpmock
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Setup mock responders with fixed IDs
	groupID := uuid.New().String()
	resourceGroupID := uuid.New().String()
	projectID := uuid.New().String()

	// Create a map to store entities with mutex for thread safety
	var entitiesLock sync.RWMutex
	entities := make(map[string]*pam.APIServiceAccountCheckoutSettings)

	httpmock.RegisterRegexpResponder("GET",
		regexp.MustCompile(`/v1/teams/httpmock-test-team/resource_groups/.*/projects/.*/okta_universal_directory_checkout_settings`),
		func(req *http.Request) (*http.Response, error) {
			matches := regexp.MustCompile(`/resource_groups/([^/]+)/projects/([^/]+)/`).FindStringSubmatch(req.URL.Path)
			if len(matches) != 3 {

				return httpmock.NewStringResponse(400, "Invalid URL"), nil
			}
			resourceGroupID := matches[1]
			projectID := matches[2]

			entityKey := fmt.Sprintf("%s/%s", resourceGroupID, projectID)

			entitiesLock.RLock()
			settings, exists := entities[entityKey]
			entitiesLock.RUnlock()

			if !exists {
				defaultSettings := &pam.APIServiceAccountCheckoutSettings{
					CheckoutRequired:          false,
					CheckoutDurationInSeconds: int32(900),
					IncludeList:               []pam.ServiceAccountSettingNameObject{},
					ExcludeList:               []pam.ServiceAccountSettingNameObject{},
				}
				return httpmock.NewJsonResponse(200, defaultSettings)
			}

			return httpmock.NewJsonResponse(200, settings)
		},
	)

	httpmock.RegisterRegexpResponder("PUT",
		regexp.MustCompile(`/v1/teams/httpmock-test-team/resource_groups/.*/projects/.*/okta_universal_directory_checkout_settings`),
		func(req *http.Request) (*http.Response, error) {

			matches := regexp.MustCompile(`/resource_groups/([^/]+)/projects/([^/]+)/`).FindStringSubmatch(req.URL.Path)
			if len(matches) != 3 {
				return httpmock.NewStringResponse(400, "Invalid URL"), nil
			}
			resourceGroupID := matches[1]
			projectID := matches[2]

			var settings pam.APIServiceAccountCheckoutSettings
			if err := json.NewDecoder(req.Body).Decode(&settings); err != nil {

				return httpmock.NewStringResponse(400, ""), nil
			}

			entityKey := fmt.Sprintf("%s/%s", resourceGroupID, projectID)

			entitiesLock.Lock()
			entities[entityKey] = &settings
			entitiesLock.Unlock()

			return httpmock.NewJsonResponse(200, settings)
		},
	)

	// Register default responders for user group, resource group, resource group project
	SetupDefaultMockResponders(groupID, resourceGroupID, projectID, delegatedAdminGroupName, resourceGroupName, projectName)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProtoV6ProviderFactories: httpMockTestV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: createOktaUDCheckoutSettingsCreateConfig(delegatedAdminGroupName, resourceGroupName, projectName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPreRefresh: []plancheck.PlanCheck{
						DebugPlan(),
					},
				},

				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "checkout_required", "true"),
					resource.TestCheckResourceAttr(resourceName, "checkout_duration_in_seconds", "900"),
				),
			},
			{
				Config: createOktaUDCheckoutSettingsUpdateWithIncludeListConfig(delegatedAdminGroupName, resourceGroupName, projectName),
				Check: resource.ComposeAggregateTestCheckFunc(
					func(s *terraform.State) error {
						return nil
					},
					resource.TestCheckResourceAttr(resourceName, "checkout_required", "true"),
					resource.TestCheckResourceAttr(resourceName, "checkout_duration_in_seconds", "3600"),
					resource.TestCheckResourceAttr(resourceName, "include_list.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "include_list.0.id", "account1"),
					resource.TestCheckResourceAttr(resourceName, "include_list.0.service_account_user_name", "user1"),
					resource.TestCheckResourceAttr(resourceName, "include_list.0.saas_app_instance_name", "app1"),
				),
			},
			{

				Config: createOktaUDCheckoutSettingsUpdateWithExcludeListConfig(delegatedAdminGroupName, resourceGroupName, projectName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "checkout_required", "true"),
					resource.TestCheckResourceAttr(resourceName, "checkout_duration_in_seconds", "3600"),
					resource.TestCheckResourceAttr(resourceName, "exclude_list.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "exclude_list.0.id", "account3"),
					resource.TestCheckResourceAttr(resourceName, "exclude_list.0.service_account_user_name", "user3"),
					resource.TestCheckResourceAttr(resourceName, "exclude_list.0.saas_app_instance_name", "app3"),
				),
			},
		},
	})
}

func testAccOktaUDCheckoutSettingsCheckExists(resourceName string, expected *pam.APIServiceAccountCheckoutSettings) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Okta UD checkout settings not found: %s", resourceName)
		}

		resourceGroupID := rs.Primary.Attributes[attributes.ResourceGroup]
		projectID := rs.Primary.Attributes[attributes.Project]

		client := mustTestAccAPIClients().SDKClient
		settings, _, err := client.SDKClient.ProjectsAPI.FetchResourceGroupOktaUniversalDirectoryBasedProjectCheckoutSettings(context.Background(), client.Team, resourceGroupID, projectID).Execute()
		if err != nil {
			return fmt.Errorf("Error fetching Okta UD checkout settings: %s", err)
		}

		if settings.CheckoutRequired != expected.CheckoutRequired {
			return fmt.Errorf("Okta UD checkout settings checkout required does not match: %t != %t", settings.CheckoutRequired, expected.CheckoutRequired)
		}

		if settings.CheckoutDurationInSeconds != expected.CheckoutDurationInSeconds {
			return fmt.Errorf("Okta UD checkout settings checkout duration in seconds does not match: %d != %d", settings.CheckoutDurationInSeconds, expected.CheckoutDurationInSeconds)
		}

		return nil
	}
}

func testAccOktaUDCheckoutSettingsCheckDeleted(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[resourceName]
		if ok {
			return fmt.Errorf("Okta Universal Directory checkout settings still exists: %s", resourceName)
		}
		return nil
	}
}

func createOktaUDCheckoutSettingsBaseConfig(delegatedAdminGroupName string, resourceGroupName string, projectName string) string {
	return fmt.Sprintf(testAccOktaUDCheckoutSettingsBaseConfigFormat, delegatedAdminGroupName, resourceGroupName, projectName)
}

func createOktaUDCheckoutSettingsCreateConfig(delegatedAdminGroupName string, resourceGroupName string, projectName string) string {
	return createOktaUDCheckoutSettingsBaseConfig(delegatedAdminGroupName, resourceGroupName, projectName) + testAccOktaUDCheckoutSettingsCreateConfigFormat
}

func createOktaUDCheckoutSettingsUpdateWithIncludeListConfig(delegatedAdminGroupName string, resourceGroupName string, projectName string) string {
	return createOktaUDCheckoutSettingsBaseConfig(delegatedAdminGroupName, resourceGroupName, projectName) + testAccOktaUDCheckoutSettingsUpdateWithIncludeListConfigFormat
}

func createOktaUDCheckoutSettingsUpdateWithExcludeListConfig(delegatedAdminGroupName string, resourceGroupName string, projectName string) string {
	return createOktaUDCheckoutSettingsBaseConfig(delegatedAdminGroupName, resourceGroupName, projectName) + testAccOktaUDCheckoutSettingsUpdateWithExcludeListConfigFormat
}

func createOktaUDCheckoutSettingsUpdateWithBothListsConfig(delegatedAdminGroupName string, resourceGroupName string, projectName string) string {
	return createOktaUDCheckoutSettingsBaseConfig(delegatedAdminGroupName, resourceGroupName, projectName) + testAccOktaUDCheckoutSettingsUpdateWithBothListsConfigFormat
}
