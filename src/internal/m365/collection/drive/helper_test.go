package drive

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

const defaultItemSize int64 = 42

// TODO(ashmrtn): Merge with similar structs in graph and exchange packages.
type oneDriveService struct {
	credentials account.M365Config
	status      support.ControllerOperationStatus
	ac          api.Client
}

func NewOneDriveService(credentials account.M365Config) (*oneDriveService, error) {
	ac, err := api.NewClient(
		credentials,
		control.DefaultOptions(),
		count.New())
	if err != nil {
		return nil, err
	}

	service := oneDriveService{
		ac:          ac,
		credentials: credentials,
	}

	return &service, nil
}

func (ods *oneDriveService) updateStatus(status *support.ControllerOperationStatus) {
	if status == nil {
		return
	}

	ods.status = support.MergeStatus(ods.status, *status)
}

func loadTestService(t *testing.T) *oneDriveService {
	a := tconfig.NewM365Account(t)

	creds, err := a.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	service, err := NewOneDriveService(creds)
	require.NoError(t, err, clues.ToCore(err))

	return service
}
