package env

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationSSMStore(t *testing.T) {

	t.Run("Return more than the default MaxValues (10)", func(t *testing.T) {
		t.Run("When we have more than 10 env values", func(t *testing.T) {

			store := NewSSMStore(SSMStoreOptions{
				Region: endpoints.EuWest1RegionID,
			})

			serviceName := "configurateur-admin"
			prefixPath := "/os/staging/support/IT/core"

			envvars, _ := store.QueryVarsForService(serviceName, StoreQueryOptions{
				PrefixPath: prefixPath,
			})

			assert.True(t, len(envvars) > 10)

		})
	})

}
