package sharepoint

import (
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"

	"github.com/alcionai/corso/src/internal/m365/support"
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
