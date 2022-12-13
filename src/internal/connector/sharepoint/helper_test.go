package sharepoint

import (
	"testing"

	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	"github.com/alcionai/corso/src/pkg/account"
)

// ---------------------------------------------------------------------------
// SharePoint Test Services
// ---------------------------------------------------------------------------
type MockGraphService struct{}

type testService struct {
	client      msgraphsdk.GraphServiceClient
	adapter     msgraphsdk.GraphRequestAdapter
	credentials account.M365Config
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

func (ms *MockGraphService) ErrPolicy() bool {
	return false
}

func (ts *testService) Client() *msgraphsdk.GraphServiceClient {
	return &ts.client
}

func (ts *testService) Adapter() *msgraphsdk.GraphRequestAdapter {
	return &ts.adapter
}

func (ts *testService) ErrPolicy() bool {
	return false
}

// ---------------------------------------------------------------------------
// Helper Functions
// ---------------------------------------------------------------------------

func createTestService(credentials account.M365Config) (*testService, error) {
	{
		adapter, err := graph.CreateAdapter(
			credentials.AzureTenantID,
			credentials.AzureClientID,
			credentials.AzureClientSecret,
		)
		if err != nil {
			return nil, errors.Wrap(err, "creating microsoft graph service for exchange")
		}

		service := testService{
			adapter:     *adapter,
			client:      *msgraphsdk.NewGraphServiceClient(adapter),
			credentials: credentials,
		}

		return &service, nil
	}
}

func expectedPathAsSlice(t *testing.T, tenant, user string, rest ...string) []string {
	res := make([]string, 0, len(rest))

	for _, r := range rest {
		p, err := onedrive.GetCanonicalPath(r, tenant, user, onedrive.SharePointSource)
		require.NoError(t, err)

		res = append(res, p.String())
	}

	return res
}
