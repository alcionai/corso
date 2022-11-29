package sharepoint

import (
	"bytes"
	"io"

	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/account"
)

type testService struct {
	client      msgraphsdk.GraphServiceClient
	adapter     msgraphsdk.GraphRequestAdapter
	credentials account.M365Config
}

///------------------------------------------------------------
// Functions to comply with graph.Service Interface
//-------------------------------------------------------
func (ts *testService) Client() *msgraphsdk.GraphServiceClient {
	return &ts.client
}

func (ts *testService) Adapter() *msgraphsdk.GraphRequestAdapter {
	return &ts.adapter
}

func (ts *testService) ErrPolicy() bool {
	return false
}

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

func readerWrapper(byteArray []byte) io.ReadCloser {
	return io.NopCloser(bytes.NewReader(byteArray))
}
