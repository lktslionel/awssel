package env

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/stretchr/testify/assert"
)

const (
	LocalstackSSMEndpoint = "http://localhost:4583"
)

func getMockedSSMStore() *SSMStore {

	return NewSSMStore(SSMStoreOptions{
		roleARN:  "arn:aws:iam::721728311103:role/service-role/homer-fn-role",
		Endpoint: LocalstackSSMEndpoint,
		Region:   endpoints.EuWest1RegionID,
	})
}

func TestSSMStore(t *testing.T) {

	t.Run("Get a Client", func(t *testing.T) {
		actual := NewSSMStore()
		assert.NotNil(t, actual)
	})

	t.Run("Get a Client with options", func(t *testing.T) {
		actual := getMockedSSMStore()
		assert.NotNil(t, actual)
		assert.NotNil(t, actual.conn)
	})

	t.Run("Implements the Storer interface", func(t *testing.T) {
		assert.Implements(t, (*Storer)(nil), NewSSMStore())
	})

	t.Run("Retrieves env vars", func(t *testing.T) {
		t.Run("With only an unknown service", func(t *testing.T) {
			// Get a SSMStore client to first, populate the Store with values
			store := getMockedSSMStore()

			envStore := Storer(store)
			envvars, _ := envStore.QueryVarsForService("serviceA")

			assert.Empty(t, envvars)
		})

		t.Run("With a known service but without prefix path", func(t *testing.T) {
			// Get a SSMStore client to first, populate the Store with values
			store := getMockedSSMStore()

			envStore := Storer(store)
			envvars, _ := envStore.QueryVarsForService("proxy")

			assert.Empty(t, envvars)
		})

		t.Run("With a known service with a prefix path", func(t *testing.T) {
			// Get a SSMStore client to first, populate the Store with values
			store := getMockedSSMStore()
			envvars, _ := store.QueryVarsForService("proxy", StoreQueryOptions{
				PrefixPath: "/os",
			})

			assert.NotEmpty(t, envvars)
			// We chech that the result contains 2 value
			// In tasks/populate.rb file, we define 2 values
			// for thes key /os/proxy
			assert.Equal(t, len(envvars), 2)
			fmt.Println(envvars)
			// Checking that values are correct
			for _, ev := range envvars {
				assert.Contains(t, []string{"PROXY_USER", "PROXY_PASS"}, ev.Name)
			}
		})
	})

	t.Run("Filter env vars using patterns", func(t *testing.T) {
		t.Run("With awesome Regex", func(t *testing.T) {

			examples := []struct {
				serviceName string
				path        string
				pttrn       string
				expected    []*Var
			}{
				{
					"proxy",
					"/os",
					"^PR",
					[]*Var{
						&Var{"PROXY_USER", "os-operator"},
						&Var{"PROXY_PASS", "784f43631c05`"},
					},
				},
				{
					"sd-web",
					"/os/prod/support/IT/core",
					"OS.*_HTTP[S]?_.*",
					[]*Var{
						&Var{"OS_SDWEB_HTTP_URL", "http://www.sd-web.com"},
						&Var{"OS_SDWEB_HTTPS_URL", "https://www.sd-web.com"},
					},
				},
				{
					"esb",
					"/os/qa/support/IT/core",
					".*_PORT$",
					[]*Var{
						&Var{"OS_ESB_MULE_PORT", "8080"},
					},
				},
			}
			store := getMockedSSMStore()

			for _, expl := range examples {
				envvars, err := store.QueryVarsForService(expl.serviceName, StoreQueryOptions{
					PrefixPath:    expl.path,
					FilterPattern: expl.pttrn,
				})

				assert.Nil(t, err)
				assert.Equal(t, envvars, expl.expected)

			}

		})

	})

	t.Run("Return more than the default MaxValues (10)", func(t *testing.T) {
		t.Run("When we have more than 10 env values", func(t *testing.T) {

			sample := struct {
				serviceName string
				path        string
			}{
				"common",
				"/os/staging/support/it/core",
			}

			store := getMockedSSMStore()

			envvars, _ := store.QueryVarsForService(sample.serviceName, StoreQueryOptions{
				PrefixPath: sample.path,
			})

			assert.True(t, len(envvars) > 10)

		})
	})

}
