package env

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/stretchr/testify/assert"
)

func getStore(p ...string) *SSMStore {
	purpose := "local"
	options := SSMStoreOptions{
		//roleARN:  "arn:aws:iam::721728311103:role/service-role/homer-fn-role",
		profile:  "fake-profile",
		endpoint: aws.String("fake.ssm.us-east-2.amazonaws.com"),
		region:   "us-west-1",
	}

	if len(p) > 0 {
		purpose = p[0]
	}

	switch purpose {
	case "local":
		options.endpoint = aws.String("http://localhost:4583")
		break
	default:
		break
	}

	return NewSSMStore(options)
}

func TestSSMStore(t *testing.T) {

	t.Run("Get a Client", func(t *testing.T) {
		actual := NewSSMStore()
		assert.NotNil(t, actual)
	})

	t.Run("Get a Client with options", func(t *testing.T) {
		actual := getStore()
		assert.NotNil(t, actual)
		assert.NotNil(t, actual.sess)
	})

	t.Run("Implements the Storer interface", func(t *testing.T) {
		assert.Implements(t, (*Storer)(nil), NewSSMStore())
	})

	t.Run("Retrieves env vars", func(t *testing.T) {
		t.Run("With only an unknown service", func(t *testing.T) {
			// Get a SSMStore client to first, populate the Store with values
			store := getStore("local")

			envStore := Storer(store)
			envvars, _ := envStore.QueryVarsForService("serviceA")

			assert.Empty(t, envvars)
		})

		t.Run("With a known service but without prefix path", func(t *testing.T) {
			// Get a SSMStore client to first, populate the Store with values
			store := getStore("local")

			envStore := Storer(store)
			envvars, _ := envStore.QueryVarsForService("proxy")

			assert.Empty(t, envvars)
		})

		t.Run("With a known service with a prefix path", func(t *testing.T) {
			// Get a SSMStore client to first, populate the Store with values
			store := getStore("local")

			fmt.Println(store)

			envStore := Storer(store)
			envvars, err := envStore.QueryVarsForService("proxy", StoreQueryOptions{
				PrefixPath: "/os",
			})

			fmt.Println("ERROR: ", err)

			assert.NotEmpty(t, envvars)
			// We chech that the result contains 2 value
			// In tasks/populate.rb file, we define 2 values
			// for thes key /os/proxy
			assert.Equal(t, len(envvars), 2)

			// Checking that values are correct
			for _, ev := range envvars {
				assert.Contains(t, ev.Name, []string{"PROXY_USER", "PROXY_PASS"})
			}
		})
	})

}
