package api_test

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/account"
	graphapi "github.com/alcionai/corso/src/pkg/connector/graph"
)

func createTestBetaService(t *testing.T, credentials account.M365Config) *graphapi.BetaService {
	adapter, err := graph.CreateAdapter(
		credentials.AzureTenantID,
		credentials.AzureClientID,
		credentials.AzureClientSecret)
	require.NoError(t, err, clues.ToCore(err))

	return graphapi.NewBetaService(adapter)
}
