package oktapam

import (
	"os"
	"testing"
	"time"

	"math/rand"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const DefaultTestTeam = "pam-tf-provider-testing"

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

var testAccProviders map[string]func() (*schema.Provider, error)
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]func() (*schema.Provider, error){}
	testAccProviders["oktapam"] = func() (*schema.Provider, error) {
		return testAccProvider, nil
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

func randSeq() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]rune, 20) // Character length
	for i := range b {
		b[i] = randChars[r.Intn(len(randChars))]
	}
	return string(b)
}
