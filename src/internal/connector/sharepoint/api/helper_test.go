package api_test

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

func createTestBetaService(t *testing.T, credentials account.M365Config) *api.BetaService {
	adapter, err := graph.CreateAdapter(
		credentials.AzureTenantID,
		credentials.AzureClientID,
		credentials.AzureClientSecret)
	require.NoError(t, err, clues.ToCore(err))

	return api.NewBetaService(adapter)
}
