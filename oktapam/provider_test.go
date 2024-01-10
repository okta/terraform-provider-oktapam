package oktapam

import (
	"fmt"
	"os"
	"testing"
	"time"

	"math/rand"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/okta/terraform-provider-oktapam/oktapam/client"
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

const defaultRandSeqLength = 20

var randChars = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func randSeq() string {
	return randSeqWithLength(defaultRandSeqLength)
}

func randSeqWithLength(length uint) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]rune, length) // Character length
	for i := range b {
		b[i] = randChars[r.Intn(len(randChars))]
	}
	return string(b)
}

func isExecutingPAMTest() bool {
	pamAccEnv := os.Getenv("TF_ACC_PAM")
	return pamAccEnv != "" && pamAccEnv != "0"
}

func checkTeamApplicable(t *testing.T, isPAMTest bool) {
	if isExecutingPAMTest() != isPAMTest {
		t.Skip("skipping due to team/test mismatch")
	}
}

// subNamedObjects is used within tests to allow for comparing objects that include named objects.  generally the struct which is created within a test
// will only know either the ids or the names for the named objects.  this method assumes the expectedNamedObjects were returned from the server
// and will have both the id and name.  the method will match up the named objects based on the key the test knows about and return a list with the
// values from the expected list
func subNamedObjects(expectedNamedObjects, actualNamedObjects []client.NamedObject, matchByID bool) ([]client.NamedObject, error) {
	if len(expectedNamedObjects) != len(actualNamedObjects) {
		return nil, fmt.Errorf("number of named objects does not match.  expected %d, got %d", len(expectedNamedObjects), len(actualNamedObjects))
	}

	m := make(map[string]client.NamedObject)

	for _, no := range actualNamedObjects {
		if matchByID {
			m[*no.Id] = no
		} else {
			m[*no.Name] = no
		}
	}

	subs := make([]client.NamedObject, len(expectedNamedObjects))

	for idx, no := range expectedNamedObjects {
		var key string
		if matchByID {
			key = *no.Id
		} else {
			key = *no.Name
		}

		if _, ok := m[key]; ok {
			subs[idx] = no
		} else {
			return nil, fmt.Errorf("could not match named object with key: %s", key)
		}
	}

	return subs, nil
}

func fillNamedObjectValues(expectedNamedObject client.NamedObject, actualNamedObject client.NamedObject) client.NamedObject {
	filled := client.NamedObject{}

	if expectedNamedObject.Id != nil {
		filled.Id = expectedNamedObject.Id
	} else {
		filled.Id = actualNamedObject.Id
	}

	if expectedNamedObject.Name != nil {
		filled.Name = expectedNamedObject.Name
	} else {
		filled.Name = actualNamedObject.Name
	}

	if string(expectedNamedObject.Type) != "" {
		filled.Type = expectedNamedObject.Type
	} else {
		filled.Type = actualNamedObject.Type
	}

	return filled
}

func getTeamName() string {
	teamName := os.Getenv(teamSchemaEnvVar)
	if teamName != "" {
		return teamName
	}
	return DefaultTestTeam
}
