package site

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/account"
)

func createTestService(t *testing.T, credentials account.M365Config) *graph.Service {
	adapter, err := graph.CreateAdapter(
		credentials.AzureTenantID,
		credentials.AzureClientID,
		credentials.AzureClientSecret)
	require.NoError(t, err, "creating microsoft graph service for exchange", clues.ToCore(err))

	return graph.NewService(adapter)
}
