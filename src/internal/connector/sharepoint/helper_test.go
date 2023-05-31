package sharepoint

import (
	"testing"

	"github.com/alcionai/clues"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/path"
)

// ---------------------------------------------------------------------------
// SharePoint Test Services
// ---------------------------------------------------------------------------
type MockGraphService struct{}

type MockUpdater struct {
	UpdateState func(*support.ConnectorOperationStatus)
}

func (mu *MockUpdater) UpdateStatus(input *support.ConnectorOperationStatus) {
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

func (ms *MockGraphService) UpdateStatus(*support.ConnectorOperationStatus) {
}

// ---------------------------------------------------------------------------
// Helper Functions
// ---------------------------------------------------------------------------

func createTestService(t *testing.T, credentials account.M365Config) *graph.Service {
	adapter, err := graph.CreateAdapter(
		credentials.AzureTenantID,
		credentials.AzureClientID,
		credentials.AzureClientSecret)
	require.NoError(t, err, "creating microsoft graph service for exchange", clues.ToCore(err))

	return graph.NewService(adapter)
}

func expectedPathAsSlice(
	t *testing.T,
	bh onedrive.BackupHandler,
	tenant, user string,
	rest ...string,
) []string {
	res := make([]string, 0, len(rest))

	for _, r := range rest {
		pb := path.Builder{}.Append(r)
		p, err := bh.CanonicalPath(pb, tenant, user)
		require.NoError(t, err, clues.ToCore(err))

		res = append(res, p.String())
	}

	return res
}
