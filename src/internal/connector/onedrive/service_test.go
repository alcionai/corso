package onedrive

import (
	"testing"

	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
)

type MockGraphService struct{}

func (ms *MockGraphService) Client() *msgraphsdk.GraphServiceClient {
	return nil
}

func (ms *MockGraphService) Adapter() *msgraphsdk.GraphRequestAdapter {
	return nil
}

func (ms *MockGraphService) ErrPolicy() bool {
	return false
}

type oneDriveService struct {
	client      msgraphsdk.GraphServiceClient
	adapter     msgraphsdk.GraphRequestAdapter
	credentials account.M365Config
	status      support.ConnectorOperationStatus
}

func (ods *oneDriveService) Client() *msgraphsdk.GraphServiceClient {
	return &ods.client
}

func (ods *oneDriveService) Adapter() *msgraphsdk.GraphRequestAdapter {
	return &ods.adapter
}

func (ods *oneDriveService) ErrPolicy() bool {
	return false
}

func NewOneDriveService(credentials account.M365Config) (*oneDriveService, error) {
	adapter, err := graph.CreateAdapter(
		credentials.AzureTenantID,
		credentials.AzureClientID,
		credentials.AzureClientSecret,
	)
	if err != nil {
		return nil, err
	}

	service := oneDriveService{
		adapter:     *adapter,
		client:      *msgraphsdk.NewGraphServiceClient(adapter),
		credentials: credentials,
	}

	return &service, nil
}

func (ods *oneDriveService) updateStatus(status *support.ConnectorOperationStatus) {
	if status == nil {
		return
	}

	ods.status = support.MergeStatus(ods.status, *status)
}

func loadTestService(t *testing.T) *oneDriveService {
	a := tester.NewM365Account(t)
	m365, err := a.M365Config()
	require.NoError(t, err)

	service, err := NewOneDriveService(m365)
	require.NoError(t, err)

	return service
}
