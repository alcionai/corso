package api_test

import (
	"testing"

	discover "github.com/alcionai/corso/src/internal/connector/discovery/api"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/account"
)

func createTestBetaService(t *testing.T, credentials account.M365Config) *discover.BetaService {
	adapter, err := graph.CreateAdapter(
		credentials.AzureTenantID,
		credentials.AzureClientID,
		credentials.AzureClientSecret,
	)
	require.NoError(t, err)

	return discover.NewBetaService(adapter)
}
