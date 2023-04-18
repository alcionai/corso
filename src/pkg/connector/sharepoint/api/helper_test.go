package api_test

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/pkg/account"
	discover "github.com/alcionai/corso/src/pkg/connector/discovery/api"
	"github.com/alcionai/corso/src/pkg/connector/graph"
)

func createTestBetaService(t *testing.T, credentials account.M365Config) *discover.BetaService {
	adapter, err := graph.CreateAdapter(
		credentials.AzureTenantID,
		credentials.AzureClientID,
		credentials.AzureClientSecret)
	require.NoError(t, err, clues.ToCore(err))

	return discover.NewBetaService(adapter)
}
