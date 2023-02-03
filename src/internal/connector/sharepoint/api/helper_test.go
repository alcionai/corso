package api

import (
	"testing"

	"github.com/alcionai/corso/src/internal/connector/discovery/api"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/stretchr/testify/require"
)

func createTestBetaService(t *testing.T, credentials account.M365Config) *api.BetaService {
	adapter, err := graph.CreateAdapter(
		credentials.AzureTenantID,
		credentials.AzureClientID,
		credentials.AzureClientSecret,
	)
	require.NoError(t, err)

	return api.NewBetaService(adapter)
}
