package site

import (
	"testing"

	"github.com/alcionai/clues"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/canario/src/internal/m365/support"
	"github.com/alcionai/canario/src/pkg/account"
	"github.com/alcionai/canario/src/pkg/count"
	"github.com/alcionai/canario/src/pkg/services/m365/api/graph"
)

// ---------------------------------------------------------------------------
// SharePoint Test Services
// ---------------------------------------------------------------------------
type MockGraphService struct{}

type MockUpdater struct {
	UpdateState func(*support.ControllerOperationStatus)
}

func (mu *MockUpdater) UpdateStatus(input *support.ControllerOperationStatus) {
	if mu.UpdateState != nil {
		mu.UpdateState(input)
	}
}

//------------------------------------------------------------
// Interface Functions: @See graph.Service
//------------------------------------------------------------

func (ms *MockGraphService) Client() *msgraphsdk.GraphServiceClient {
	return nil
}

func (ms *MockGraphService) Adapter() *msgraphsdk.GraphRequestAdapter {
	return nil
}

func (ms *MockGraphService) UpdateStatus(*support.ControllerOperationStatus) {
}

// ---------------------------------------------------------------------------
// Helper functions
// ---------------------------------------------------------------------------

func createTestService(t *testing.T, credentials account.M365Config) *graph.Service {
	adapter, err := graph.CreateAdapter(
		credentials.AzureTenantID,
		credentials.AzureClientID,
		credentials.AzureClientSecret,
		count.New())
	require.NoError(t, err, "creating microsoft graph service for exchange", clues.ToCore(err))

	return graph.NewService(adapter)
}
