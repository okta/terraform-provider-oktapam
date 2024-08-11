package oktapam

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func Test_Terraform_Version_Check(t *testing.T) {
	t.Parallel()
	groupName := fmt.Sprintf("test_acc_group_tf_version_check_%s", randSeq())
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.RequireAbove(tfversion.Version1_2_0),
		},
		Steps: []resource.TestStep{
			{
				Config: createTestAccGroupCreateConfig(groupName),
			},
		},
	})
}
