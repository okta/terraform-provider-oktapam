package oktapam

import (
	"fmt"
	"os"
	"testing"
	"time"

	"math/rand"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"oktapam": testAccProvider,
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	requiredEnvVars := []string{apiKeySchemaEnvVar, apiKeySecretSchemaEnvVar, teamSchemaEnvVar}

	for _, envVar := range requiredEnvVars {
		if err := os.Getenv(envVar); err == "" {
			t.Fatalf("%s must be set for acceptance tests", envVar)
		}
	}
}

var randChars = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func randSeq(n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]rune, n)
	for i := range b {
		b[i] = randChars[r.Intn(len(randChars))]
	}
	return string(b)
}

func arrayContainsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func getIndexMappingFromResource(rs *terraform.ResourceState, prefix, identifierAttribute string, expectedLength int) (map[string]string, error) {
	attributes := rs.Primary.Attributes
	mapping := make(map[string]string, expectedLength)

	for i := 0; i < expectedLength; i++ {
		attrName := fmt.Sprintf("%s.%d.%s", prefix, i, identifierAttribute)
		if attr, ok := attributes[attrName]; ok {
			mapping[attr] = fmt.Sprint(i)
		} else {
			return nil, fmt.Errorf("Could not find attribute %s", attrName)
		}
	}

	return mapping, nil
}

func getArrayFromResource(rs *terraform.ResourceState, prefix string, expectedLength int) ([]string, error) {
	attributes := rs.Primary.Attributes
	arr := make([]string, 0, expectedLength)

	for i := 0; i < expectedLength; i++ {
		attrName := fmt.Sprintf("%s.%d", prefix, i)
		if attr, ok := attributes[attrName]; ok {
			arr = append(arr, attr)
		} else {
			return nil, fmt.Errorf("Could not find attribute %s", attrName)
		}
	}

	return arr, nil
}
